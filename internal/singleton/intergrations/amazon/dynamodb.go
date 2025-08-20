package amazon

import (
	"gin-starter/internal/singleton/config"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewDynamoDBClient(cfg config.AppConfig) *dynamodb.Client {
	return dynamodb.NewFromConfig(*cfg.Aws)
}
