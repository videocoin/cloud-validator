package service

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Name    string `envconfig:"-"`
	Version string `envconfig:"-"`

	RPCAddr string `required:"true" envconfig:"RPC_ADDR" default:"validator:50055"`

	RPCNodeHTTPAddr           string `default:"" envconfig:"RPC_NODE_HTTP_ADDR"`
	StreamManagerContractAddr string `default:"" envconfig:"STREAM_MANAGER_CONTRACT_ADDR"`
	Key                       string `required:"true" envconfig:"KEY"`
	Secret                    string `required:"true" envconfig:"SECRET"`

	Logger *logrus.Entry `envconfig:"-"`
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
