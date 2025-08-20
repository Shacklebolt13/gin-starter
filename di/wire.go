//go:build wireinject

package di

import (
	"cert/internal/controller"
	"cert/internal/controller/health"
	"cert/internal/singleton/config"
	"cert/internal/singleton/intergrations/amazon"

	"github.com/google/wire"
)

var configSet = wire.NewSet(
	config.ParseEnvironment,
	config.ConfigureAws,
	config.NewAppConfig,
)

var intergrationSet = wire.NewSet(
	configSet,
	amazon.NewDynamoDBClient,
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
