package handler

import (
	"day-36/service"
	"day-36/util"

	"go.uber.org/zap"
)

type AllHandler struct {
	ProductHandler ProductHandler
}

func NewAllHandler(service service.AllService, log *zap.Logger, config util.Configuration) AllHandler {
	return AllHandler{
		ProductHandler: NewProductHandler(service, log, config),
	}

}
