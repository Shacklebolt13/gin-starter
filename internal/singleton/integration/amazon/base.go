package amazon

import (
	"sync"

	cidp "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

var onceAmazonIntegration sync.Once
var amazonIntegration *AmazonIntegration

type AmazonIntegration struct {
	DynamoDBClient *dynamodb.Client
	CIDPClient     *cidp.Client
}

func NewAmazonIntegration(dynamoDbClient *dynamodb.Client, cidpClient *cidp.Client) *AmazonIntegration {
	onceAmazonIntegration.Do(func() {
		amazonIntegration = &AmazonIntegration{
			DynamoDBClient: dynamoDbClient,
			CIDPClient:     cidpClient,
		}
	})

	return amazonIntegration
}
