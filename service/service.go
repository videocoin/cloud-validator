package service

import (
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

	rpcConfig := &RpcServerOptions{
		Addr:     cfg.RPCAddr,
		Contract: contract,
		Logger:   cfg.Logger,
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
