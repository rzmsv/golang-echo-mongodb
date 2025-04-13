package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/rzmsv/golang-project/core"
	mongo_model "github.com/rzmsv/golang-project/internal/model/mongo"
	"github.com/rzmsv/golang-project/internal/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductHandelr struct {
	productService *service.ProductService
}

func NewProductHandler(productService *service.ProductService) *ProductHandelr {
	return &ProductHandelr{
		productService: productService,
	}
}

/* ----------------------------------- GET ---------------------------------- */
func (H *ProductHandelr) ProductList(c echo.Context) error {
	result, err := H.productService.ProductList()
	if err != nil {
		return core.InternalServerError(err.Error(), c)
	}
	if len(result) == 0 {
		return core.ResponseOk([]mongo_model.Products{}, c)
	}
	return core.ResponseOk(result, c)
}

func (H *ProductHandelr) ProductById(c echo.Context) error {
	id := c.Param("id")
	result, err := H.productService.ProductById(id)
	if err != nil {
		return core.InternalServerError(err.Error(), c)
	}
	if result.ID.IsZero() {
		return core.NotFound(c)
	}
	return core.ResponseOk(result, c)
}

/* ---------------------------------- POST ---------------------------------- */
func (H *ProductHandelr) CreateNewProduct(c echo.Context) error {
	var reqBody mongo_model.Products
	/* -------------------------- Generate new objectID ------------------------- */
	reqBody.ID = primitive.NewObjectID()
	/* ----------------------------------- --- ---------------------------------- */
	if err := c.Bind(&reqBody); err != nil {
		return core.BadRequestHandler(err.Error(), c)
	}
	result, err := H.productService.CreateNewProduct(&reqBody)
	if err != nil {
		return core.InternalServerError(err.Error(), c)
	}
	return core.Created(result, c)
}

/* ---------------------------------- PATCH --------------------------------- */
func (H *ProductHandelr) UpdateProduct(c echo.Context) error {
	var reqBody mongo_model.Products
	id := c.Param("id")
	if err := c.Bind(&reqBody); err != nil {
		return core.BadRequestHandler(err.Error(), c)
	}
	if err := c.Validate(&reqBody); err != nil {
		return core.BadRequestHandler(err.Error(), c)
	}
	result, err := H.productService.UpdateProduct(id, reqBody)
	if err != nil {
		return core.InternalServerError(err.Error(), c)
	}
	if result == "not_found!" {
		return core.NotFound(c)
	}
	return core.ResponseOk(result, c)
}

/* --------------------------------- DELETE --------------------------------- */
func (H *ProductHandelr) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	result, err := H.productService.DeleteProduct(id)
	if err != nil {
		return core.InternalServerError(err.Error(), c)
	}
	if result == "not_found!" {
		return core.NotFound(c)
	}
	return core.ResponseOk(result, c)
}
