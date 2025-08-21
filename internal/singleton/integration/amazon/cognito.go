package amazon

import (
	"gin-starter/internal/singleton/config"
	"sync"

	cidp "github.com/aws/aws-sdk-go-v2/service/cognitoidentityprovider"
)

var cidpOnce sync.Once
var cognitoClient *cidp.Client

func NewCidpClient(appConfig *config.AppConfig) *cidp.Client {
	cidpOnce.Do(func() {
		cognitoClient = cidp.NewFromConfig(*appConfig.Aws)
	})

	return cognitoClient
}
