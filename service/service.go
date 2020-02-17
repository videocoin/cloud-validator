package service

import (
	pstreamsv1 "github.com/videocoin/cloud-api/streams/private/v1"
	"github.com/videocoin/cloud-pkg/grpcutil"
	"github.com/videocoin/cloud-validator/contract"
)

type Service struct {
	cfg *Config
	rpc *RpcServer
}

func NewService(cfg *Config) (*Service, error) {
	contractOpts := &contract.ContractClientOpts{
		RPCNodeHTTPAddr: cfg.RPCNodeHTTPAddr,
		ContractAddr:    cfg.StreamManagerContractAddr,
		Key:             cfg.Key,
		Secret:          cfg.Secret,
		Logger:          cfg.Logger.WithField("system", "contract"),
	}

	contract, err := contract.NewContractClient(contractOpts)
	if err != nil {
		return nil, err
	}

	conn, err := grpcutil.Connect(cfg.StreamsRPCAddr, cfg.Logger.WithField("system", "streamscli"))
	if err != nil {
		return nil, err
	}
	streams := pstreamsv1.NewStreamsServiceClient(conn)

	rpcConfig := &RpcServerOptions{
		Addr:          cfg.RPCAddr,
		Contract:      contract,
		Threshold:     cfg.Threshold,
		Logger:        cfg.Logger,
		BaseInputURL:  cfg.BaseInputURL,
		BaseOutputURL: cfg.BaseOutputURL,
		Streams:       streams,
	}

	rpc, err := NewRpcServer(rpcConfig)
	if err != nil {
		return nil, err
	}

	svc := &Service{
		cfg: cfg,
		rpc: rpc,
	}

	return svc, nil
}

func (s *Service) Start() error {
	go s.rpc.Start()
	return nil
}

func (s *Service) Stop() error {
	return nil
}
