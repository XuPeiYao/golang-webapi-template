package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"go-webapi-template/internal/app/controllers"
	"go-webapi-template/internal/domain/errors"
	"go-webapi-template/pkg/common"
)

func ConfigureFiber() fiber.Config {
	return fiber.Config{
		ServerHeader: "Fiber",
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if domainError, ok := err.(*errors.DomainError); ok {
				return ctx.Status(domainError.StatusCode).JSON(domainError.Serialize())
			} else {
				serverDomainError := errors.ServerError.New(err)
				return ctx.Status(serverDomainError.StatusCode).JSON(serverDomainError.Serialize())
			}
		},
	}
}

func ConfigureFiberRouting(
	requestResponseLoggingMiddleware *common.FiberRequestResponseHookMiddleware,
	helloworldController *controllers.HelloWorldController,
	productsController *controllers.ProductsController,
) common.ConfigureFiberAppFunc {

	return func(fApp *common.FiberApp) {
		app := fApp.Engine
		app.Use(recover.New())
		app.Use(requestid.New())
		app.Use(requestResponseLoggingMiddleware.Handler())

		// custom middleware: add custom header
		app.Use(func(c *fiber.Ctx) error {
			err := c.Next()
			c.Response().Header.Add("EXAMPLE-CUSTOM-HEADER", "HELLOWORLD")
			return err
		})

		// api
		fApp.ConfigureRoutes(helloworldController, productsController)

		// swagger ui
		app.Get("/swagger", func(ctx *fiber.Ctx) error {
			return ctx.Redirect("/swagger/index.html")
		})
		app.Static("/", "./assets", fiber.Static{
			Compress: true,
		})
	}
}
