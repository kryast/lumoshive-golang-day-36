package handler

import (
	"day-36/service"
	"day-36/util"

	"go.uber.org/zap"
)

type AllHandler struct {
	AuthHandler    AuthHandler
	ProductHandler ProductHandler
}

func NewAllHandler(service service.AllService, log *zap.Logger, config util.Configuration) AllHandler {
	return AllHandler{
		AuthHandler:    NewAuthHandler(service, log, config),
		ProductHandler: NewProductHandler(service, log, config),
	}

}
