package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rzmsv/golang-project/core"
	"github.com/rzmsv/golang-project/internal/service"
)

type ProductHandelr struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandelr {
	return &ProductHandelr{
		productService: productService,
	}
}

func (H ProductHandelr) ProductsList(c echo.Context) error {
	return core.BadRequestHandler("dsadsadsa", c)
}
