package config

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	aws_config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/rs/zerolog/log"
)

func ConfigureAws() *aws.Config {
	cfg, err := aws_config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal().Err(err).Send()
	}
	return &cfg
}
