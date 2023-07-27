package common

import (
	"github.com/gofiber/fiber/v2"
)

type ConfigureFiberAppFunc func(*FiberApp)

type FiberApp struct {
	Configuration Configuration
	Engine        *fiber.App

	fiberConfig fiber.Config
	startupFunc ConfigureFiberAppFunc
}

const ConfigurationFiberListen = "fiber:listen"
const DefaultFiberListen = ":3000"

func NewFiberApp(
	configuration Configuration,
	fiberConfig fiber.Config,
	startupFunc ConfigureFiberAppFunc,
) IApp {
	return &FiberApp{
		Configuration: configuration,
		fiberConfig:   fiberConfig,
		startupFunc:   startupFunc,
	}
}

func (this *FiberApp) ConfigureRoutes(controllers ...IController) {
	for _, controller := range controllers {
		controller.ConfigureRoutes(this)
	}
}

func (this *FiberApp) Run() error {
	this.Engine = fiber.New(this.fiberConfig)

	this.startupFunc(this)

	return this.Engine.Listen(this.Configuration.GetValueOrDefault(ConfigurationFiberListen, DefaultFiberListen).(string))
}
