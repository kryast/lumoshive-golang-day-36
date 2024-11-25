package service

import (
	"day-36/model"
	"day-36/repository"
	"errors"

	"go.uber.org/zap"
)

type ProductServiceInterface interface {
	GetProductById(productId int) (*model.Product, error)
}

type ProductService struct {
	Repo   repository.AllRepository
	Logger *zap.Logger
}

func NewProductService(repo repository.AllRepository, logger *zap.Logger) ProductServiceInterface {
	return &ProductService{
		Repo:   repo,
		Logger: logger,
	}
}

func (ps *ProductService) GetProductById(productId int) (*model.Product, error) {
	product, err := ps.Repo.ProductRepo.GetProductById(productId)
	if err != nil {
		ps.Logger.Error("Failed to get product by ID", zap.Int("productId", productId), zap.Error(err))
		return nil, errors.New("failed to get product by id")
	}
	return product, nil
}
