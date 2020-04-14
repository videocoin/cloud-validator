package service

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	emitterv1 "github.com/videocoin/cloud-api/emitter/v1"
	pstreamsv1 "github.com/videocoin/cloud-api/streams/private/v1"
	v1 "github.com/videocoin/cloud-api/validator/v1"
	"github.com/videocoin/cloud-pkg/grpcutil"
	"github.com/videocoin/cloud-pkg/retry"
	"github.com/videocoin/cloud-validator/eventbus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthv1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type RPCServerOptions struct {
	Logger        *logrus.Entry
	Addr          string
	Threshold     int
	BaseInputURL  string
	BaseOutputURL string
	EB            *eventbus.EventBus
	Streams       pstreamsv1.StreamsServiceClient
	Emitter       emitterv1.EmitterServiceClient
}

type RPCServer struct {
	logger        *logrus.Entry
	grpc          *grpc.Server
	listen        net.Listener
	addr          string
	threshold     int
	baseInputURL  string
	baseOutputURL string
	eb            *eventbus.EventBus
	streams       pstreamsv1.StreamsServiceClient
	emitter       emitterv1.EmitterServiceClient
}

func NewRPCServer(opts *RPCServerOptions) (*RPCServer, error) {
	grpcOpts := grpcutil.DefaultServerOpts(opts.Logger)
	gRPCServer := grpc.NewServer(grpcOpts...)

	healthService := health.NewServer()
	healthv1.RegisterHealthServer(gRPCServer, healthService)

	listen, err := net.Listen("tcp", opts.Addr)
	if err != nil {
		return nil, err
	}

	RPCServer := &RPCServer{
		logger:        opts.Logger,
		addr:          opts.Addr,
		threshold:     opts.Threshold,
		grpc:          gRPCServer,
		listen:        listen,
		baseInputURL:  opts.BaseInputURL,
		baseOutputURL: opts.BaseOutputURL,
		eb:            opts.EB,
		streams:       opts.Streams,
		emitter:       opts.Emitter,
	}

	v1.RegisterValidatorServiceServer(gRPCServer, RPCServer)
	reflection.Register(gRPCServer)

	return RPCServer, nil
}

func (s *RPCServer) Start() error {
	s.logger.Infof("starting rpc server on %s", s.addr)
	return s.grpc.Serve(s.listen)
}

func (s *RPCServer) ValidateProof(ctx context.Context, req *v1.ValidateProofRequest) (*types.Empty, error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("stream_contract_address", req.StreamContractAddress)

	profileID := new(big.Int)
	profileID.SetBytes(req.ProfileId)

	outputChunkID := new(big.Int)
	outputChunkID.SetBytes(req.OutputChunkId)

	span.SetTag("profile_id", profileID.String())
	span.SetTag("output_chunk_id", outputChunkID.String())

	go func(streamID, streamContractAddress string, profileID, outputChunkID *big.Int) {
		inputChunkURL := fmt.Sprintf("%s/%s/%d.ts", s.baseInputURL, streamID, outputChunkID.Int64()-1)
		outputChunkURL := fmt.Sprintf("%s/%s/%d.ts", s.baseOutputURL, streamID, outputChunkID.Int64())

		logger := s.logger.WithFields(
			logrus.Fields{
				"original":   inputChunkURL,
				"transcoded": outputChunkURL,
			})

		isValid, err := s.validateProof(inputChunkURL, outputChunkURL)
		if err != nil {
			logger.Errorf("failed to validate proof: %s", err)
			return
		}

		vctx, _ := context.WithTimeout(context.Background(), time.Second*60*2)
		if isValid {
			logger.Info("validate proof")

			validateProofReq := &emitterv1.ValidateProofRequest{
				StreamContractAddress: streamContractAddress,
				ProfileId:             profileID.Bytes(),
				ChunkId:               outputChunkID.Bytes(),
			}
			validateProofResp, err := s.emitter.ValidateProof(vctx, validateProofReq)
			if err != nil {
				logger.WithError(err).Error("failed to call contract validate proof")
				return
			}

			logger.Debugf("tx %s", string(validateProofResp.TxId))

			err = s.eb.EmitEvent(context.Background(), &v1.Event{
				Type:                  v1.EventTypeValidatedProof,
				StreamContractAddress: streamContractAddress,
				ChunkNum:              outputChunkID.Uint64(),
			})
			if err != nil {
				logger.WithError(err).Error("failed to emit validated proof event")
				return
			}
		} else {
			logger.Info("scrap proof")

			scrapProofReq := &emitterv1.ScrapProofRequest{
				StreamContractAddress: streamContractAddress,
				ProfileId:             profileID.Bytes(),
				ChunkId:               outputChunkID.Bytes(),
			}
			scrapProofResp, err := s.emitter.ScrapProof(vctx, scrapProofReq)
			if err != nil {
				logger.WithError(err).Error("failed to scrap proof")
				return
			}

			logger.Debugf("tx %s", string(scrapProofResp.TxId))

			err = s.eb.EmitEvent(context.Background(), &v1.Event{
				Type:                  v1.EventTypeScrapedProof,
				StreamContractAddress: streamContractAddress,
				ChunkNum:              outputChunkID.Uint64(),
			})
			if err != nil {
				logger.WithError(err).Error("failed to emit scrap proof event")
				return
			}
		}

		if req.IsLast {
			pdReq := &pstreamsv1.StreamRequest{Id: req.StreamId}
			_, err := s.streams.PublishDone(context.Background(), pdReq)
			if err != nil {
				logger.WithError(err).Error("failed to publish done")
				return
			}
		}

	}(req.StreamId, req.StreamContractAddress, profileID, outputChunkID)

	return new(types.Empty), nil
}

func (s *RPCServer) validateProof(inputChunkURL, outputChunkURL string) (bool, error) {
	logger := s.logger.WithFields(
		logrus.Fields{
			"original":   inputChunkURL,
			"transcoded": outputChunkURL,
		})

	err := retry.RetryWithAttempts(10, time.Second*5, func() error {
		logger.Infof("checking output chunk url %s", outputChunkURL)
		return checkSource(outputChunkURL)
	})
	if err != nil {
		return false, err
	}

	inDuration, err := getDuration(inputChunkURL)
	if err != nil || inDuration == 0 {
		return false, fmt.Errorf("failed to get input chunk duration: %s", err)
	}

	logger.Debugf("original duration is %f", inDuration)

	outDuration, err := getDuration(outputChunkURL)
	if err != nil || outDuration == 0 {
		return false, fmt.Errorf("failed to get output chunk duration: %s", err)
	}

	logger.Debugf("transcoded duration is %f", outDuration)

	duration := math.Min(inDuration, outDuration)
	seekTo := rand.Float64() * duration

	logger.Debugf("duration is %f, extracting at time %f", duration, seekTo)

	inFrame, err := extractFrame(inputChunkURL, seekTo)
	if err != nil {
		return false, fmt.Errorf("failed to extract input chunk frame: %s", err)
	}
	defer func() {
		err := os.Remove(inFrame)
		if err != nil {
			logger.WithError(err).Error("failed to remove input chunk frame")
		}
	}()

	inHash, err := getHash(inFrame)
	if err != nil {
		return false, fmt.Errorf("failed to get input chunk hash: %s", err)
	}

	outFrame, err := extractFrame(outputChunkURL, seekTo)
	if err != nil {
		return false, fmt.Errorf("failed to extract output chunk frame: %s", err)
	}
	defer func() {
		err := os.Remove(outFrame)
		if err != nil {
			logger.WithError(err).Error("failed to remove output chunk frame")
		}
	}()

	outHash, err := getHash(outFrame)
	if err != nil {
		return false, fmt.Errorf("failed to get output chunk hash: %s", err)
	}

	distance, err := inHash.Distance(outHash)
	if err != nil {
		return false, fmt.Errorf("failed to calc distance: %s", err)
	}

	logger.Infof("distance is %d", distance)

	// [0,32], 0 - same, 32 - completely different
	if distance <= s.threshold {
		return true, nil
	}

	return false, nil
}
