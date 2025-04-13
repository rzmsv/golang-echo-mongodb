package service

import (
	mongo_model "github.com/rzmsv/golang-project/internal/model/mongo"
	"github.com/rzmsv/golang-project/internal/repository"
)

type ProductService struct {
	productRepo *repository.ProductRepository
}

func NewProductService(productRepo *repository.ProductRepository) *ProductService {
	return &ProductService{
		productRepo: productRepo,
	}
}

/* ----------------------------------- GET ---------------------------------- */
func (S *ProductService) ProductList() ([]mongo_model.Products, error) {
	result, err := S.productRepo.ProductList()
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (S *ProductService) ProductById(id string) (mongo_model.Products, error) {
	result, err := S.productRepo.ProductById(id)
	if err != nil {
		return mongo_model.Products{}, err
	}
	return result, nil
}

/* ---------------------------------- POST ---------------------------------- */
func (S *ProductService) CreateNewProduct(data *mongo_model.Products) (string, error) {
	result, err := S.productRepo.CreateNewProduct(data)
	if err != nil {
		return "", err
	}
	return result, nil
}

/* ---------------------------------- PATCH --------------------------------- */
func (S *ProductService) UpdateProduct(id string, data mongo_model.Products) (string, error) {
	result, err := S.productRepo.UpdateProduct(id, data)
	if err != nil {
		return "", err
	}
	return result, nil
}

/* --------------------------------- DELETE --------------------------------- */
func (S *ProductService) DeleteProduct(id string) (string, error) {
	result, err := S.productRepo.DeleteProduct(id)
	if err != nil {
		return "", err
	}
	return result, nil
}
