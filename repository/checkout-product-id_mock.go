package repository

import (
	"day-36/model"

	"github.com/stretchr/testify/mock"
)

type ProductRepositoryMock struct {
	mock.Mock
}

func (productRepositoryMock *ProductRepositoryMock) GetProductById(productId int) (*model.Product, error) {
	args := productRepositoryMock.Called(productId)
	if productResult := args.Get(0); productResult != nil {
		return productResult.(*model.Product), args.Error(1)
	}
	return nil, args.Error(1)
}
