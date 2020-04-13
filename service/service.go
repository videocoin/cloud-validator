package service

import (
	emitterv1 "github.com/videocoin/cloud-api/emitter/v1"
	pstreamsv1 "github.com/videocoin/cloud-api/streams/private/v1"
	"github.com/videocoin/cloud-pkg/grpcutil"
	"github.com/videocoin/cloud-validator/eventbus"
)

type Service struct {
	cfg *Config
	rpc *RPCServer
	eb  *eventbus.EventBus
}

func NewService(cfg *Config) (*Service, error) {
	conn, err := grpcutil.Connect(cfg.StreamsRPCAddr, cfg.Logger.WithField("system", "streamscli"))
	if err != nil {
		return nil, err
	}
	streams := pstreamsv1.NewStreamsServiceClient(conn)

	conn, err = grpcutil.Connect(cfg.EmitterRPCAddr, cfg.Logger.WithField("system", "emittercli"))
	if err != nil {
		return nil, err
	}
	emitter := emitterv1.NewEmitterServiceClient(conn)

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
		Threshold:     cfg.Threshold,
		Logger:        cfg.Logger,
		BaseInputURL:  cfg.BaseInputURL,
		BaseOutputURL: cfg.BaseOutputURL,
		EB:            eb,
		Streams:       streams,
		Emitter:       emitter,
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
