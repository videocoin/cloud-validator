package service

import (
	"context"
	"fmt"
	"math"
	"math/big"
	"net"
	"os"
	"time"

	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	emitterv1 "github.com/videocoin/cloud-api/emitter/v1"
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

func (s *RPCServer) ValidateProof(ctx context.Context, req *v1.ValidateProofRequest) (*v1.ValidateProofResponse, error) {
	profileID := new(big.Int).SetBytes(req.ProfileId)
	chunkID := new(big.Int).SetBytes(req.ChunkId)

	span := opentracing.SpanFromContext(ctx)
	span.SetTag("stream_id", req.StreamId)
	span.SetTag("stream_contract_address", req.StreamContractAddress)
	span.SetTag("profile_id", profileID.String())
	span.SetTag("output_chunk_id", chunkID.String())

	streamID := req.StreamId
	streamContractAddress := req.StreamContractAddress

	inputChunkURL := fmt.Sprintf("%s/%s/%d.ts", s.baseInputURL, streamID, chunkID.Int64()-1)
	outputChunkURL := fmt.Sprintf("%s/%s/%d.ts", s.baseOutputURL, streamID, chunkID.Int64())

	logger := s.logger.WithFields(
		logrus.Fields{
			"original":   inputChunkURL,
			"transcoded": outputChunkURL,
		})

	resp := &v1.ValidateProofResponse{}

	otCtx := opentracing.ContextWithSpan(context.Background(), span)

	isValid, err := s.validateProof(inputChunkURL, outputChunkURL)
	if !isValid || err != nil {
		if err != nil {
			logger.Errorf("failed to validate proof: %s", err)
		}

		logger.Info("scrap proof")

		scrapProofReq := &emitterv1.ScrapProofRequest{
			StreamContractAddress: streamContractAddress,
			ProfileId:             profileID.Bytes(),
			ChunkId:               chunkID.Bytes(),
		}
		scrapProofResp, err := s.emitter.ScrapProof(otCtx, scrapProofReq)
		if err != nil {
			logger.WithError(err).Error("failed to emitter.ScrapProof")
			if scrapProofResp != nil {
				resp.ScrapProofTx = scrapProofResp.Tx
				resp.ScrapProofTxStatus = scrapProofResp.Status
			}
			return resp, err
		}

		resp.ScrapProofTx = scrapProofResp.Tx
		resp.ScrapProofTxStatus = scrapProofResp.Status

		logger.Debugf("scrap proof tx %s", scrapProofResp.Tx)

		err = s.eb.EmitEvent(otCtx, &v1.Event{
			Type:                  v1.EventTypeScrapedProof,
			StreamContractAddress: streamContractAddress,
			ChunkNum:              chunkID.Uint64(),
		})
		if err != nil {
			logger.WithError(err).Error("failed to emit scrap proof event")
			return resp, nil
		}

		return resp, nil
	}

	logger.Info("validate proof")

	validateProofReq := &emitterv1.ValidateProofRequest{
		StreamContractAddress: streamContractAddress,
		ProfileId:             profileID.Bytes(),
		ChunkId:               chunkID.Bytes(),
	}
	validateProofResp, err := s.emitter.ValidateProof(otCtx, validateProofReq)
	if err != nil {
		logger.WithError(err).Error("failed to emitter.ValidateProof")
		if validateProofResp != nil {
			resp.ValidateProofTx = validateProofResp.Tx
			resp.ValidateProofTxStatus = validateProofResp.Status
		}
		return resp, err
	}

	logger.Debugf("validate proof tx %s", validateProofResp.Tx)

	resp.ValidateProofTx = validateProofResp.Tx
	resp.ValidateProofTxStatus = validateProofResp.Status

	err = s.eb.EmitEvent(otCtx, &v1.Event{
		Type:                  v1.EventTypeValidatedProof,
		StreamContractAddress: streamContractAddress,
		ChunkNum:              chunkID.Uint64(),
	})
	if err != nil {
		logger.WithError(err).Error("failed to emit validated proof event")
		return resp, nil
	}

	return resp, nil
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

	inFrames, err := getFrames(inputChunkURL)
	if err != nil || inFrames == 0 {
		return false, fmt.Errorf("failed to get input chunk frames: %s", err)
	}

	logger.Debugf("original frames is %d", inFrames)

	outFrames, err := getFrames(outputChunkURL)
	if err != nil || outFrames == 0 {
		return false, fmt.Errorf("failed to get output chunk frames: %s", err)
	}

	logger.Debugf("transcoded duration is %d", outFrames)

	mf := int(math.Min(float64(inFrames), float64(outFrames)))
	seekTo := mf / 2

	logger.Debugf("frames is %d, extracting at frame %d", mf, seekTo)

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
