package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	"github.com/videocoin/cloud-pkg/logger"
	"github.com/videocoin/cloud-pkg/tracer"
	"github.com/videocoin/cloud-validator/service"
)

var (
	ServiceName string = "validator"
	Version     string = "dev"
)

func main() {
	logger.Init(ServiceName, Version)  //nolint

	log := logrus.NewEntry(logrus.New())

	log.Logger.SetReportCaller(true)

	log = logrus.WithFields(logrus.Fields{
		"service": ServiceName,
		"version": Version,
	})

	closer, err := tracer.NewTracer(ServiceName)
	if err != nil {
		log.Info(err.Error())
	} else {
		defer closer.Close()
	}

	cfg := service.LoadConfig(ServiceName)

	cfg.Name = ServiceName
	cfg.Version = Version
	cfg.Logger = log

	svc, err := service.NewService(cfg)
	if err != nil {
		log.Fatal(err.Error())
	}

	signals := make(chan os.Signal, 1)
	exit := make(chan bool, 1)

	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signals

		log.Infof("recieved signal %s", sig)
		exit <- true
	}()

	log.Info("starting")
	go log.Error(svc.Start())

	<-exit

	log.Info("stopping")
	err = svc.Stop()
	if err != nil {
		log.Error(err)
		return
	}

	log.Info("stopped")
}
