package config

import "github.com/aws/aws-sdk-go-v2/aws"

type AppConfig struct {
	Aws *aws.Config
	Env *envConfig
}

func NewAppConfig(env *envConfig, awsConfig *aws.Config) *AppConfig {
	return &AppConfig{
		Aws: awsConfig,
		Env: env,
	}
}
