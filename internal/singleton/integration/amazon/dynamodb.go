package amazon

import (
	"gin-starter/internal/singleton/config"
	"sync"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var dynamodbClient *dynamodb.Client
var dynamoDbOnce sync.Once

func NewDynamoDBClient(cfg *config.AppConfig) *dynamodb.Client {
	dynamoDbOnce.Do(func() {
		dynamodbClient = dynamodb.NewFromConfig(*cfg.Aws)
	})
	return dynamodbClient
}
