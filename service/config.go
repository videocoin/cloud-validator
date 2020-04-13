package service

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Name    string        `envconfig:"-"`
	Version string        `envconfig:"-"`
	Logger  *logrus.Entry `envconfig:"-"`

	RPCAddr        string `envconfig:"RPC_ADDR" default:"127.0.0.1:5020"`
	Threshold      int    `envconfig:"THRESHOLD" default:"10"`
	BaseInputURL   string `envconfig:"BASE_INPUT_URL" required:"true"`
	BaseOutputURL  string `envconfig:"BASE_OUTPUT_URL" required:"true"`
	MQURI          string `envconfig:"MQURI" default:"amqp://guest:guest@127.0.0.1:5672"`
	StreamsRPCAddr string `envconfig:"STREAMS_RPC_ADDR" default:"0.0.0.0:5102"`
	EmitterRPCAddr string `envconfig:"EMITTER_RPC_ADDR" default:"0.0.0.0:5003"`
}

var cfg Config
var once sync.Once

func LoadConfig(serviceName string) *Config {
	once.Do(func() {
		err := envconfig.Process(serviceName, &cfg)
		if err != nil {
			logrus.Fatalf("failed to load config: %s", err.Error())
		}
	})
	return &cfg
}
