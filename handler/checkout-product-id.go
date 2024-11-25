package handler

import (
	"day-36/helper"
	"day-36/service"
	"day-36/util"
	"net/http"

	"go.uber.org/zap"
)

type ProductHandler struct {
	Service service.AllService
	Logger  *zap.Logger
	Config  util.Configuration
}

func NewProductHandler(service service.AllService, logger *zap.Logger, config util.Configuration) ProductHandler {
	return ProductHandler{
		Service: service,
		Logger:  logger,
		Config:  config,
	}
}

func (ph *ProductHandler) GetProductByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := helper.GetID(w, r)

	product, err := ph.Service.ProductService.GetProductById(id)
	if err != nil {
		ph.Logger.Error("Error fetching product by ID", zap.Error(err))
		return
	}

	helper.SuccessResponse(w, 200, "Success", product)
}
