//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"go-webapi-template/internal/app"
	"go-webapi-template/internal/app/controllers"
	serviceProviders2 "go-webapi-template/internal/domain/services"
	"go-webapi-template/internal/infra/repos"
	"go-webapi-template/internal/infra/services"
	"go-webapi-template/pkg/common"
)

// <editor-fold desc="Default configuration file path">
var defaultConfigurationFilePath = common.ConfigurationFilePath("./configs/appsettings.json")

func GetConfigurationFilePath() common.ConfigurationFilePath {
	return defaultConfigurationFilePath
}

// </editor-fold>

var DIContainer = wire.NewSet(
	// Add configuration
	GetConfigurationFilePath,
	common.LoadConfigurationFromJsonFile,

	// Add logging
	app.NewLoggerProvider,

	app.ProcessRequestLoggingFuncProvider,
	app.ProcessResponseLoggingFuncProvider,
	common.NewFiberRequestResponseHookMiddleware,

	// Add http client
	app.NewHttpClient,

	// Add fiber app
	app.ConfigureFiber,
	app.ConfigureFiberRouting,
	common.NewFiberApp,

	// Add repos
	repos.NewHelloWorldRepository,

	// Add services
	serviceProviders2.NewHelloWorldService,

	// Add services
	services.NewApiEmployeesServiceProvider,

	// Add controllers
	controllers.NewHelloWorldController,
	controllers.NewProductsController,
)

func InitializeFiberApp() common.IApp {
	wire.Build(DIContainer)
	return &common.FiberApp{}
}
