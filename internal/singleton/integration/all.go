package integration

import (
	"gin-starter/internal/singleton/integration/amazon"
	"sync"
)

var onceIntegration sync.Once
var integration *Integration

type Integration struct {
	Amz *amazon.AmazonIntegration
}

func NewIntegration(amz *amazon.AmazonIntegration) (*Integration, error) {
	onceIntegration.Do(func() {
		integration = &Integration{
			Amz: amz,
		}
	})
	return integration, nil
}
