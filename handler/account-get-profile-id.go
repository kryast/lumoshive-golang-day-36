package handler

import (
	"day-36/helper"
	"net/http"
)

func (ah *AuthHandler) GetProfileByID(w http.ResponseWriter, r *http.Request) {
	id := helper.GetID(w, r)

	profile, err := ah.Service.AuthService.GetProfileByID(id)
	if err != nil {
		helper.SuccessResponse(w, 404, "ID not found", nil)
		return
	}

	helper.SuccessResponse(w, 200, "success", profile)
}
