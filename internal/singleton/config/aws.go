package config

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	aws_config "github.com/aws/aws-sdk-go-v2/config"
	"github.com/rs/zerolog/log"
)

var awsCfg aws.Config
var awsOnce sync.Once

func ConfigureAws() *aws.Config {
	awsOnce.Do(func() {
		var err error

		awsCfg, err = aws_config.LoadDefaultConfig(context.TODO())
		if err != nil {
			log.Fatal().Err(err).Send()
		}
	})

	return &awsCfg
}
