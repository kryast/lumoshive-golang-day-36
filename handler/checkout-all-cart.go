package handler

import (
	"day-36/helper"
	"net/http"

	"go.uber.org/zap"
)

func (ph *ProductHandler) GetAllCart(w http.ResponseWriter, r *http.Request) {
	carts, err := ph.Service.ProductService.GetAllCart()
	if err != nil {
		ph.Logger.Error("Error retrieving cart items", zap.Error(err))
		return
	}

	helper.SuccessResponse(w, 200, "success", carts)
}
