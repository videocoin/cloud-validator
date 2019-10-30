package service

import (
	"sync"

	"github.com/kelseyhightower/envconfig"
	"github.com/sirupsen/logrus"
)

type Config struct {
	Name    string `envconfig:"-"`
	Version string `envconfig:"-"`

	RPCAddr                   string `required:"true" envconfig:"RPC_ADDR" default:"127.0.0.1:5020"`
	RPCNodeHTTPAddr           string `required:"true" envconfig:"RPC_NODE_HTTP_ADDR"`
	StreamManagerContractAddr string `required:"true" envconfig:"STREAM_MANAGER_CONTRACT_ADDR"`
	Key                       string `required:"true" envconfig:"VALIDATOR_KEY"`
	Secret                    string `required:"true" envconfig:"VALIDATOR_SECRET"`
	Threshold                 int    `required:"true" envconfig:"THRESHOLD" default:"10"`
	BaseInputURL              string `required:"true" envconfig:"BASE_INPUT_URL"`
	BaseOutputURL             string `required:"true" envconfig:"BASE_OUTPUT_URL"`

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
