package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-webapi-template/internal/app/models"
	"go-webapi-template/internal/domain/services"
	"go-webapi-template/pkg/common"
)

// base url: /api
type HelloWorldController struct {
	helloworldService *services.HelloWorldService
}

func NewHelloWorldController(hwService *services.HelloWorldService) *HelloWorldController {
	return &HelloWorldController{
		helloworldService: hwService,
	}
}

func (this *HelloWorldController) ConfigureRoutes(app *common.FiberApp) {
	app.Engine.Get("/api/helloWorld", this.GetHelloWorld)
	app.Engine.Get("/api/helloWorld/error", this.GetError)
}

// @Summary	Get HelloWorld String
// @Tags	HelloWorld
// @version	1.0
// @produce	application/json
// @Success	200 {object}	models.HelloWorldApiResponse
// @Router	/api/helloWorld	[get]
func (this *HelloWorldController) GetHelloWorld(c *fiber.Ctx) error {
	result, err := this.helloworldService.HelloWorld()

	if err != nil {
		return err
	}

	response := &models.HelloWorldApiResponse{
		Message: result,
	}

	return c.JSON(response)
}

// @Summary	Get Error
// @Tags	HelloWorld
// @version	1.0
// @produce	application/json
// @Success	200 {object}	models.HelloWorldApiResponse
// @Router	/api/helloWorld/error	[get]
func (this *HelloWorldController) GetError(c *fiber.Ctx) error {
	err := this.helloworldService.DoSomething()

	return err
}
