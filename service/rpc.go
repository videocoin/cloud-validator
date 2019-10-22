package service

import (
	"context"
	"math"
	"math/big"
	"math/rand"
	"net"
	"os"

	"github.com/gogo/protobuf/types"
	protoempty "github.com/gogo/protobuf/types"
	"github.com/opentracing/opentracing-go"
	"github.com/sirupsen/logrus"
	rpc "github.com/videocoin/cloud-api/rpc"
	v1 "github.com/videocoin/cloud-api/validator/v1"
	"github.com/videocoin/cloud-pkg/grpcutil"
	"github.com/videocoin/cloud-validator/contract"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type RpcServerOptions struct {
	Addr      string
	Contract  *contract.ContractClient
	Threshold int
	Logger    *logrus.Entry
}

type RpcServer struct {
	grpc   *grpc.Server
	listen net.Listener
	addr   string

	contract  *contract.ContractClient
	threshold int

	logger *logrus.Entry
}

func NewRpcServer(opts *RpcServerOptions) (*RpcServer, error) {
	grpcOpts := grpcutil.DefaultServerOpts(opts.Logger)
	grpcServer := grpc.NewServer(grpcOpts...)

	listen, err := net.Listen("tcp", opts.Addr)
	if err != nil {
		return nil, err
	}

	rpcServer := &RpcServer{
		addr:      opts.Addr,
		contract:  opts.Contract,
		threshold: opts.Threshold,
		grpc:      grpcServer,
		listen:    listen,
		logger:    opts.Logger,
	}

	v1.RegisterValidatorServiceServer(grpcServer, rpcServer)
	reflection.Register(grpcServer)

	return rpcServer, nil
}

func (s *RpcServer) Start() error {
	s.logger.Infof("starting rpc server on %s", s.addr)
	return s.grpc.Serve(s.listen)
}

func (s *RpcServer) Health(ctx context.Context, req *protoempty.Empty) (*rpc.HealthStatus, error) {
	return &rpc.HealthStatus{Status: "OK"}, nil
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

	span.SetTag("input_chunk_url", string(req.InputChunkUrl))
	span.SetTag("output_chunk_url", string(req.OutputChunkUrl))

	go func() {
		logger := s.logger.WithFields(
			logrus.Fields{"original": req.InputChunkUrl, "transcoded": req.OutputChunkUrl})

		inDuration, err := getDuration(req.InputChunkUrl)
		if err != nil || inDuration == 0 {
			logger.Error("failed to get duration")
			return
		}

		logger.Debugf("original duration is %f\n", inDuration)

		outDuration, err := getDuration(req.OutputChunkUrl)
		if err != nil || outDuration == 0 {
			logger.WithError(err).Error("failed to get duration")
			return
		}

		logger.Debugf("transcoded duration is %f\n", outDuration)

		duration := math.Min(inDuration, outDuration)
		seekTo := rand.Float64() * duration

		logger.Debugf("duration is %f, extracting at time %f\n", duration, seekTo)

		inFrame, err := extractFrame(req.InputChunkUrl, seekTo)
		if err != nil {
			logger.WithError(err).Error("failed to extract frame")
			return
		}
		defer func() {
			err := os.Remove(inFrame)
			if err == nil {
				logger.WithError(err).Error("failed to remove frame")
			}
		}()

		inHash, err := getHash(inFrame)
		if err != nil {
			logger.WithError(err).Error("failed to get hash")
			return
		}

		outFrame, err := extractFrame(req.OutputChunkUrl, seekTo)
		if err != nil {
			logger.WithError(err).Error("failed to extract frame")
			return
		}
		defer func() {
			err := os.Remove(outFrame)
			if err == nil {
				logger.WithError(err).Error("failed to remove frame")
			}
		}()

		outHash, err := getHash(outFrame)
		if err != nil {
			logger.WithError(err).Error("failed to get hash")
			return
		}

		distance, err := inHash.Distance(outHash)
		if err != nil {
			logger.WithError(err).Error("failed to get distance")
			return
		}

		logger.Infof("distance is %d\n", distance)

		// [0,32], 0 - same, 32 - completely different
		if distance <= s.threshold {
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
