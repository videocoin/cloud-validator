package service

import (
	"time"

	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpclogrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpctracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpcprometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/opentracing/opentracing-go"
	emitterv1 "github.com/videocoin/cloud-api/emitter/v1"
	"github.com/videocoin/cloud-validator/eventbus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Service struct {
	cfg *Config
	rpc *RPCServer
	eb  *eventbus.EventBus
}

func NewService(cfg *Config) (*Service, error) {
	grpcDialOpts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc.UnaryClientInterceptor(grpcmiddleware.ChainUnaryClient(
				grpctracing.UnaryClientInterceptor(
					grpctracing.WithTracer(opentracing.GlobalTracer()),
				),
				grpcprometheus.UnaryClientInterceptor,
				grpclogrus.UnaryClientInterceptor(cfg.Logger),
			)),
		),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{
			Time:                time.Second * 10,
			Timeout:             time.Second * 10,
			PermitWithoutStream: true,
		}),
	}

	conn, err := grpc.Dial(cfg.EmitterRPCAddr, grpcDialOpts...)
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
