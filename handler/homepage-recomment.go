package handler

import (
	"day-36/helper"
	"net/http"

	"go.uber.org/zap"
)

func (ph *ProductHandler) GetAllRecomment(w http.ResponseWriter, r *http.Request) {
	recomments, err := ph.Service.ProductService.GetAllRecomment()
	if err != nil {
		ph.Logger.Error("Error retrieving recommendations", zap.Error(err))
		http.Error(w, "Could not retrieve recommendations", http.StatusInternalServerError)
		return
	}

	helper.SuccessResponse(w, 200, "success", recomments)
}
