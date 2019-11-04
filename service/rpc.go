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
	v1 "github.com/videocoin/cloud-api/validator/v1"
	"github.com/videocoin/cloud-pkg/grpcutil"
	"github.com/videocoin/cloud-pkg/retry"
	"github.com/videocoin/cloud-validator/contract"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

type RpcServerOptions struct {
	Addr          string
	Contract      *contract.ContractClient
	Threshold     int
	Logger        *logrus.Entry
	BaseInputURL  string
	BaseOutputURL string
}

type RpcServer struct {
	grpc   *grpc.Server
	listen net.Listener
	addr   string

	contract  *contract.ContractClient
	threshold int

	logger *logrus.Entry

	baseInputURL  string
	baseOutputURL string
}

func NewRpcServer(opts *RpcServerOptions) (*RpcServer, error) {
	grpcOpts := grpcutil.DefaultServerOpts(opts.Logger)
	grpcServer := grpc.NewServer(grpcOpts...)
	healthService := health.NewServer()
	grpc_health_v1.RegisterHealthServer(grpcServer, healthService)
	listen, err := net.Listen("tcp", opts.Addr)
	if err != nil {
		return nil, err
	}

	rpcServer := &RpcServer{
		addr:          opts.Addr,
		contract:      opts.Contract,
		threshold:     opts.Threshold,
		grpc:          grpcServer,
		listen:        listen,
		logger:        opts.Logger,
		baseInputURL:  opts.BaseInputURL,
		baseOutputURL: opts.BaseOutputURL,
	}

	v1.RegisterValidatorServiceServer(grpcServer, rpcServer)
	reflection.Register(grpcServer)

	return rpcServer, nil
}

func (s *RpcServer) Start() error {
	s.logger.Infof("starting rpc server on %s", s.addr)
	return s.grpc.Serve(s.listen)
}

func (s *RpcServer) ValidateProof(ctx context.Context, req *v1.ValidateProofRequest) (*types.Empty, error) {
	span := opentracing.SpanFromContext(ctx)
	span.SetTag("stream_contract_address", req.StreamContractAddress)

	profileID := new(big.Int)
	profileID.SetBytes(req.ProfileId)

	outputChunkID := new(big.Int)
	outputChunkID.SetBytes(req.OutputChunkId)

	span.SetTag("profile_id", profileID.String())
	span.SetTag("output_chunk_id", outputChunkID.String())

	inputChunkURL := fmt.Sprintf("%s/%s/%d.ts", s.baseInputURL, req.StreamId, outputChunkID.Int64()-1)
	outputChunkURL := fmt.Sprintf("%s/%s/%d.ts", s.baseOutputURL, req.StreamId, outputChunkID.Int64())

	go func() {

		logger := s.logger.WithFields(
			logrus.Fields{
				"original":   inputChunkURL,
				"transcoded": outputChunkURL,
			})

		isValid, err := s.validateProof(inputChunkURL, outputChunkURL)
		if err != nil {
			logger.Error(err)

			_, err := s.contract.ValidateProof(ctx, req.StreamContractAddress, profileID, outputChunkID)
			if err != nil {
				logger.WithError(err).Error("failed to validate proof")
				return
			}

			return
		}

		if isValid {
			tx, err := s.contract.ValidateProof(ctx, req.StreamContractAddress, profileID, outputChunkID)
			if err != nil {
				if tx != nil {
					logger.Debugf("tx %s\n", tx.Hash().String())
				}
				logger.WithError(err).Error("failed to validate proof")
				return
			}

			logger.Debugf("tx %s\n", tx.Hash().String())
		} else {
			tx, err := s.contract.ScrapProof(ctx, req.StreamContractAddress, profileID, outputChunkID)
			if err != nil {
				if tx != nil {
					logger.Debugf("tx %s\n", tx.Hash().String())
				}
				logger.WithError(err).Error("failed to scrap proof")
				return
			}

			logger.Debugf("tx %s\n", tx.Hash().String())
		}

	}()

	return new(types.Empty), nil
}

func (s *RpcServer) validateProof(inputChunkURL, outputChunkURL string) (bool, error) {
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

	logger.Debugf("original duration is %f\n", inDuration)

	outDuration, err := getDuration(outputChunkURL)
	if err != nil || outDuration == 0 {
		return false, fmt.Errorf("failed to get output chunk duration: %s", err)
	}

	logger.Debugf("transcoded duration is %f\n", outDuration)

	duration := math.Min(inDuration, outDuration)
	seekTo := rand.Float64() * duration

	logger.Debugf("duration is %f, extracting at time %f\n", duration, seekTo)

	inFrame, err := extractFrame(inputChunkURL, seekTo)
	if err != nil {
		return false, fmt.Errorf("failed to extract input chunk frame: %s", err)
	}
	defer func() {
		err := os.Remove(inFrame)
		if err == nil {
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
		if err == nil {
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

	logger.Infof("distance is %d\n", distance)

	// [0,32], 0 - same, 32 - completely different
	if distance <= s.threshold {
		return true, nil
	}

	return false, nil
}
