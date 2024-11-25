package service

import (
	"day-36/model"

	"go.uber.org/zap"
)

func (ps *ProductService) GetAllBanner() ([]*model.Banner, error) {
	banners, err := ps.Repo.ProductRepo.GetAllBanner()
	if err != nil {
		ps.Logger.Error("Error fetching banners", zap.Error(err))
		return nil, err
	}

	return banners, nil
}
