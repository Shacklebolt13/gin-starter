package intergrations

import "github.com/aws/aws-sdk-go-v2/service/dynamodb"

type Integration struct {
	DynamoDb *dynamodb.Client
}
