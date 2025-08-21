//go:build wireinject

package di

import (
	"gin-starter/internal/controller"
	"gin-starter/internal/controller/health"
	"gin-starter/internal/singleton/config"
	"gin-starter/internal/singleton/integration"
	"gin-starter/internal/singleton/integration/amazon"

	"github.com/google/wire"
)

var configSet = wire.NewSet(
	config.ParseEnvironment,
	config.ConfigureAws,
	config.NewProcessConfig,
	config.NewAppConfig,
)

var integrationSet = wire.NewSet(
	configSet,
	amazon.NewDynamoDBClient,
	amazon.NewCidpClient,
	amazon.NewAmazonIntegration,
	integration.NewIntegration,
)

var healthControllerSet = wire.NewSet(
	health.NewHealthController,
)

var controllerSet = wire.NewSet(
	healthControllerSet,
	controller.NewUrlMapping,
)

func ProvideHealthController() (health.HealthController, error) {
	wire.Build(healthControllerSet)
	return nil, nil
}

func ProvideUrlMappings() controller.UrlMapping {
	wire.Build(controllerSet)
	return nil
}

func ProvideAppConfig() (*config.AppConfig, error) {
	wire.Build(configSet)
	return nil, nil
}

func ProvideIntegration() (*integration.Integration, error) {
	wire.Build(integrationSet)
	return nil, nil
}
