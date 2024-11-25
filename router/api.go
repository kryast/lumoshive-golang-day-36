package router

import (
	"day-36/database"
	"day-36/handler"
	"day-36/middleware"
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
	r.Post("/register", Handle.AuthHandler.Register)
	r.Post("/login", Handle.AuthHandler.Login)

	// Homepage - Ecommerce Home Page
	r.Get("/banner", Handle.ProductHandler.GetAllBanner)
	r.Get("/categories", Handle.ProductHandler.GetAllCategory)
	r.Get("/best-seller", Handle.ProductHandler.GetBestSellers)
	r.Get("/products/promo-weekly", Handle.ProductHandler.GetPromoWeekly)
	r.Get("/products/search", Handle.ProductHandler.SearchProducts)
	r.Get("/products/recomment", Handle.ProductHandler.GetAllRecomment)
	r.Post("/products/{id}/wishlist", Handle.ProductHandler.AddToWishlistHandler)
	r.Delete("/products/{id}/wishlist", Handle.ProductHandler.DeleteFromWishlist)
	r.Get("/products/total-cart", Handle.ProductHandler.GetTotalCartQuantity)

	// Homepage - All Product and Category
	r.Post("/products/{id}/add-cart/{quantity}", Handle.ProductHandler.AddToCart)
	r.Post("/products/{id}/add-cart", Handle.ProductHandler.AddToCart)

	// Checkout - Flow
	r.Get("/products/{id}", Handle.ProductHandler.GetProductByIdHandler)
	r.Get("/products/cart", Handle.ProductHandler.GetAllCart)
	r.Put("/products/{id}/cart/{quantity}", Handle.ProductHandler.UpdateCartQuantity)
	r.Delete("/products/{id}/cart", Handle.ProductHandler.DeleteCartItemHandler)
	r.With(middleware.CheckLoginMiddleware(db, logger)).Group(func(r chi.Router) {
		r.Get("/auth/{id}/order", Handle.ProductHandler.CreateOrder)
	})

	// Account Section
	r.Get("/profile/{id}", Handle.AuthHandler.GetProfileByID)
	r.Put("/profile/{id}", Handle.AuthHandler.UpdateProfile)
	r.Post("/profile/{id}", Handle.AuthHandler.AddAddress)
	r.Get("/profile/{id}/address", Handle.AuthHandler.GetAddressesByAuthID)
	r.Put("/profile/{id}/address", Handle.AuthHandler.UpdateAddress)
	r.Delete("/profile/{id}/address", Handle.AuthHandler.DeleteAddress)

	return r, logger, nil
}
