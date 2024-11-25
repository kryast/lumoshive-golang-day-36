package router

import (
	"day-36/database"
	"day-36/handler"
	"day-36/repository"
	"day-36/service"
	"day-36/util"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

func InitRouter() (*chi.Mux, *zap.Logger, error) {
	r := chi.NewRouter()

	logger := util.InitLog()

	config, err := util.ReadConfiguration()
	if err != nil {
		return nil, logger, err
	}

	db, err := database.InitDB(config)
	if err != nil {
		return nil, logger, err
	}

	repo := repository.NewAllRepository(db, logger)
	service := service.NewAllService(repo, logger)
	Handle := handler.NewAllHandler(service, logger, config)

	// Auth Section

	// Homepage - Ecommerce Home Page

	// Homepage - All Product and Category

	// Checkout - Flow
	r.Get("/products/{id}", Handle.ProductHandler.GetProductByIdHandler)

	// Account Section

	return r, logger, nil
}
