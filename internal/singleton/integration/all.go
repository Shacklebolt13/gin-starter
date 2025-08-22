package integration

import (
	"gin-starter/internal/singleton/integration/amazon"
	"sync"
)

var onceIntegration sync.Once
var integration *ExternalClients

type ExternalClients struct {
	Amz *amazon.AmazonIntegration
}

func NewIntegration(amz *amazon.AmazonIntegration) *ExternalClients {
	onceIntegration.Do(func() {
		integration = &ExternalClients{
			Amz: amz,
		}
	})

	return integration
}
