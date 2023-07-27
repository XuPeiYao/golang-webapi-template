package controllers

import (
	"github.com/gofiber/fiber/v2"
	"go-webapi-template/internal/domain/services"
	"go-webapi-template/pkg/common"
)

// base url: /api
type ProductsController struct {
	productsServiceProvider services.IProductsServiceProvider
}

func NewProductsController(productsServiceProvider services.IProductsServiceProvider) *ProductsController {
	return &ProductsController{
		productsServiceProvider: productsServiceProvider,
	}
}

func (this *ProductsController) ConfigureRoutes(app *common.FiberApp) {
	app.Engine.Get("/api/products", this.GetProducts)
}

// @Summary	Get Product
// @Tags	Products
// @version	1.0
// @produce	application/json
// @Success	200 {array}	[]services.Product
// @Router	/api/products	[get]
func (this *ProductsController) GetProducts(c *fiber.Ctx) error {
	result := this.productsServiceProvider.GetProducts()
	return c.JSON(result)
}
