package handler

import (
	"day-36/helper"
	"net/http"

	"go.uber.org/zap"
)

func (ph *ProductHandler) GetPromoWeekly(w http.ResponseWriter, r *http.Request) {
	promotions, err := ph.Service.ProductService.GetPromoWeekly()
	if err != nil {
		ph.Logger.Error("Error retrieving weekly promotions", zap.Error(err))
		http.Error(w, "Could not retrieve promotions", http.StatusInternalServerError)
		return
	}

	helper.SuccessResponse(w, 200, "success", promotions)
}
