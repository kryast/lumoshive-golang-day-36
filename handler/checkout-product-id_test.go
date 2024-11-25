package handler

import (
	"day-36/model"
	"day-36/service"
	"day-36/util"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

func TestGetProductByIdHandler_Success(t *testing.T) {
	mockService := service.ProductServiceMock{}
	allService := service.AllService{
		ProductService: &mockService,
	}
	config := util.Configuration{}
	logger := zap.NewNop()
	productHandler := NewProductHandler(allService, logger, config)

	expectedProduct := &model.Product{
		ID:       1,
		Name:     "Test Product",
		Detail:   "This is a test product",
		Price:    100.0,
		Category: "Electronics",
		PhotoURL: "http://example.com/product.jpg",
	}

	mockService.On("GetProductById", 1).Once().Return(expectedProduct, nil)

	req := httptest.NewRequest(http.MethodGet, "/products/1", nil)
	w := httptest.NewRecorder()

	productHandler.GetProductByIdHandler(w, req)

	resp := w.Result()

	assert.Equal(t, http.StatusOK, resp.StatusCode)

}
