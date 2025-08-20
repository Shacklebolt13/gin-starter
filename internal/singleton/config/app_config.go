package config

import (
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
)

type AppConfig struct {
	Aws     *aws.Config
	Env     *EnvConfig
	Process *ProcessConfig
}

var config *AppConfig

func NewAppConfig(
	env *EnvConfig,
	awsConfig *aws.Config,
	processConfig *ProcessConfig,
) *AppConfig {
	sync.OnceFunc(func() {
		config = &AppConfig{
			Aws:     awsConfig,
			Env:     env,
			Process: processConfig,
		}
	})()
	return config
}
