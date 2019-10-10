package service

import (
	"context"
	"math/big"
	"net"

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
	Addr     string
	Contract *contract.ContractClient
	Logger   *logrus.Entry
}

type RpcServer struct {
	grpc   *grpc.Server
	listen net.Listener
	addr   string

	contract *contract.ContractClient

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
		addr:     opts.Addr,
		contract: opts.Contract,
		grpc:     grpcServer,
		listen:   listen,
		logger:   opts.Logger,
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
	span.SetTag("stream_contract_id", req.StreamContractId)

	profileID := new(big.Int)
	profileID.SetBytes(req.ProfileId)

	inputChunkID := new(big.Int)
	inputChunkID.SetBytes(req.InputChunkId)

	outputChunkID := new(big.Int)
	outputChunkID.SetBytes(req.OutputChunkId)

	span.SetTag("profile_id", profileID.String())
	span.SetTag("input_chunk_id", inputChunkID.String())
	span.SetTag("output_chunk_id", outputChunkID.String())

	tx, err := s.contract.ValidateProof(ctx, req.StreamContractAddress, profileID, outputChunkID)
	if err != nil {
		if tx != nil {
			s.logger.Debugf("tx %s\n", tx.Hash().String())
		}
		s.logger.Errorf("failed to validate proof: %+v", err.Error())
		return nil, rpc.ErrRpcInternal
	}

	s.logger.Debugf("tx %s\n", tx.Hash().String())

	return new(types.Empty), err

}
