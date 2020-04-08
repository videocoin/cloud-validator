package service

import (
	pstreamsv1 "github.com/videocoin/cloud-api/streams/private/v1"
	"github.com/videocoin/cloud-pkg/grpcutil"
	"github.com/videocoin/cloud-validator/contract"
	"github.com/videocoin/cloud-validator/eventbus"
)

type Service struct {
	cfg *Config
	rpc *RPCServer
	eb  *eventbus.EventBus
}

func NewService(cfg *Config) (*Service, error) {
	contractOpts := &contract.ClientOpts{
		RPCNodeHTTPAddr: cfg.RPCNodeHTTPAddr,
		ContractAddr:    cfg.StreamManagerContractAddr,
		Key:             cfg.Key,
		Secret:          cfg.Secret,
		Logger:          cfg.Logger.WithField("system", "contract"),
	}

	contract, err := contract.NewClient(contractOpts)
	if err != nil {
		return nil, err
	}

	conn, err := grpcutil.Connect(cfg.StreamsRPCAddr, cfg.Logger.WithField("system", "streamscli"))
	if err != nil {
		return nil, err
	}
	streams := pstreamsv1.NewStreamsServiceClient(conn)

	eb, err := eventbus.NewEventBus(
		cfg.MQURI,
		eventbus.WithLogger(cfg.Logger.WithField("system", "eventbus")),
		eventbus.WithName("validator"),
	)
	if err != nil {
		return nil, err
	}

	rpcConfig := &RPCServerOptions{
		Addr:          cfg.RPCAddr,
		Contract:      contract,
		Threshold:     cfg.Threshold,
		Logger:        cfg.Logger,
		BaseInputURL:  cfg.BaseInputURL,
		BaseOutputURL: cfg.BaseOutputURL,
		Streams:       streams,
		EB:            eb,
	}

	rpc, err := NewRPCServer(rpcConfig)
	if err != nil {
		return nil, err
	}

	svc := &Service{
		cfg: cfg,
		rpc: rpc,
		eb:  eb,
	}

	return svc, nil
}

func (s *Service) Start(errCh chan error) {
	go func() {
		errCh <- s.rpc.Start()
	}()

	go func() {
		errCh <- s.eb.Start()
	}()
}

func (s *Service) Stop() error {
	return s.eb.Stop()
}
