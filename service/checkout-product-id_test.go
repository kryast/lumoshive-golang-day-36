package service

import (
	"day-36/model"
	"day-36/repository"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetProductById(t *testing.T) {

	logger := zap.NewNop()
	productRepoMock := new(repository.ProductRepositoryMock)
	repo := repository.AllRepository{
		ProductRepo: productRepoMock,
	}

	service := NewProductService(repo, logger)

	t.Run("Success", func(t *testing.T) {
		product := &model.Product{
			ID:   1,
			Name: "Test Product",
		}

		productRepoMock.On("GetProductById", 1).Once().Return(product, nil)

		result, err := service.GetProductById(1)

		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, product.Name, result.Name)
		productRepoMock.AssertExpectations(t)
	})

	t.Run("ProductNotFound", func(t *testing.T) {
		productRepoMock.On("GetProductById", 999).Once().Return(nil, errors.New("ID not found"))

		result, err := service.GetProductById(999)

		assert.Error(t, err)
		assert.Nil(t, result)
		assert.Equal(t, "failed to get product by id", err.Error())
		productRepoMock.AssertExpectations(t)
	})

}
