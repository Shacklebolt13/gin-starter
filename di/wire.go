//go:build wireinject

package di

import (
	"gin-starter/internal/controller"
	"gin-starter/internal/controller/health"
	"gin-starter/internal/service/integration/amazon/cognito"
	"gin-starter/internal/service/integration/amazon/cognito/usecase/user/create"
	"gin-starter/internal/service/integration/amazon/cognito/usecase/user/login"
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

var integrationConfigSet = wire.NewSet(
	amazon.NewDynamoDBClient,
	amazon.NewCidpClient,
	amazon.NewAmazonIntegration,
	integration.NewIntegration,
)

var cognitoSet = wire.NewSet(
	login.NewLoginService,
	create.NewCreateUserService,
	cognito.NewCognitoService,
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

func ProvideIntegrationConfig() (*integration.ExternalClients, error) {
	wire.Build(configSet, integrationConfigSet)
	return nil, nil
}

func ProvideCognitoService() (*cognito.CognitoService, error) {
	wire.Build(configSet, integrationConfigSet, cognitoSet)
	return nil, nil
}
