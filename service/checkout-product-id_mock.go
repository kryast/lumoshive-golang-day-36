package service

import (
	"day-36/model"

	"github.com/stretchr/testify/mock"
)

type ProductServiceMock struct {
	mock.Mock
}

func (ProductServiceMock *ProductServiceMock) GetProductById(productId int) (*model.Product, error) {
	args := ProductServiceMock.Called(productId)
	if ProductResult := args.Get(0); ProductResult != nil {
		return ProductResult.(*model.Product), args.Error(1)
	}
	return nil, args.Error(1)
}
