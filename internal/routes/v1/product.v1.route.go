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
	v1 := p.echo.Group("/api/v1")
	products := v1.Group("/products")
	products.GET("", p.productHandler.ProductsList)
}
