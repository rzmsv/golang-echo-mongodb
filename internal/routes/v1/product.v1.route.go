package v1

import (
	"github.com/labstack/echo/v4"
	"github.com/rzmsv/golang-project/internal/handler"
)

type ProductRoute struct {
	echo           *echo.Echo
	productHandler *handler.ProductHandelr
}

func NewProductRoutes(productHandler *handler.ProductHandelr, e *echo.Echo) *ProductRoute {
	return &ProductRoute{
		echo:           e,
		productHandler: productHandler,
	}
}

func (p *ProductRoute) ProductV1ApiRoutes() {
	/* ------------------------ prifix: /api/v1/products ------------------------ */
	v1 := p.echo.Group("/api/v1")
	products := v1.Group("/products")
	/* ----------------------------------- GET ---------------------------------- */
	products.GET("", p.productHandler.ProductList)
	products.GET("/:id", p.productHandler.ProductById)
	/* ---------------------------------- POST ---------------------------------- */
	products.POST("/new-products", p.productHandler.CreateNewProduct)
	/* ---------------------------------- PATCH --------------------------------- */
	products.PATCH("/:id", p.productHandler.UpdateProduct)
	/* --------------------------------- DELETE --------------------------------- */
	products.DELETE("/:id", p.productHandler.DeleteProduct)
}
