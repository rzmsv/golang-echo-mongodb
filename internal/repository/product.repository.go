package repository

import (
	"github.com/rzmsv/golang-project/config"
)

type ProductRepository struct {
	configs *config.AppConfigs
}

func NewProductRepository(configs *config.AppConfigs) *ProductRepository {
	return &ProductRepository{
		configs: configs,
	}
}
