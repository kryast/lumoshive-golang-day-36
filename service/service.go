package service

import (
	"day-36/repository"

	"go.uber.org/zap"
)

type AllService struct {
	ProductService ProductServiceInterface
}

func NewAllService(repo repository.AllRepository, log *zap.Logger) AllService {
	return AllService{
		ProductService: NewProductService(repo, log),
	}
}
