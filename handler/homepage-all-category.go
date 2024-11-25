package handler

import (
	"day-36/helper"
	"net/http"

	"go.uber.org/zap"
)

func (ph *ProductHandler) GetAllCategory(w http.ResponseWriter, r *http.Request) {
	categories, err := ph.Service.ProductService.GetAllCategory()
	if err != nil {
		ph.Logger.Error("Error getting products", zap.Error(err))
		http.Error(w, "Failed to retrieve products", http.StatusInternalServerError)
		return
	}

	helper.SuccessResponse(w, 200, "success", categories)
}
