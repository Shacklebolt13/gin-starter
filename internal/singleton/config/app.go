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

var appConfig *AppConfig
var once sync.Once

func NewAppConfig(
	env *EnvConfig,
	awsConfig *aws.Config,
	processConfig *ProcessConfig,
) *AppConfig {
	once.Do(func() {
		appConfig = &AppConfig{
			Aws:     awsConfig,
			Env:     env,
			Process: processConfig,
		}
	})

	return appConfig
}
