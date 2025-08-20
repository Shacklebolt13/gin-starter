//go:build wireinject

package di

import (
	"context"
	"gin-starter/internal/controller"
	"gin-starter/internal/controller/health"
	"gin-starter/internal/singleton/config"
	"gin-starter/internal/singleton/intergrations/amazon"

	"github.com/google/wire"
)

var appContext context.Context = nil
var appCancel context.CancelCauseFunc = nil

var configSet = wire.NewSet(
	config.ParseEnvironment,
	config.ConfigureAws,
	config.NewProcessConfig,
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

func ProvideAppConfig() (*config.AppConfig, error) {
	wire.Build(configSet)
	return nil, nil
}
