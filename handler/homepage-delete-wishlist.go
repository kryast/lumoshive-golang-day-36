package handler

import (
	"day-36/helper"
	"net/http"

	"go.uber.org/zap"
)

func (ph *ProductHandler) DeleteFromWishlist(w http.ResponseWriter, r *http.Request) {
	id := helper.GetID(w, r)

	err := ph.Service.ProductService.DeleteFromWishlist(id)
	if err != nil {
		ph.Logger.Error("error delete from wishlist", zap.Error(err))
		return
	}

	helper.SuccessResponse(w, 200, "Success", nil)
}