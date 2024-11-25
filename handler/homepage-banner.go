package handler

import (
	"day-36/helper"
	"net/http"

	"go.uber.org/zap"
)

func (ph *ProductHandler) GetAllBanner(w http.ResponseWriter, r *http.Request) {
	banners, err := ph.Service.ProductService.GetAllBanner()
	if err != nil {
		ph.Logger.Error("Error getting banners", zap.Error(err))
		http.Error(w, "Failed to retrieve banners", http.StatusInternalServerError)
		return
	}

	helper.SuccessResponse(w, 200, "success", banners)
}
