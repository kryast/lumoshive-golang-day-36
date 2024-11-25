package service

import (
	"day-36/model"

	"go.uber.org/zap"
)

func (ps *ProductService) GetPromoWeekly() ([]*model.PromoWeekly, error) {
	promotions, err := ps.Repo.ProductRepo.GetPromoWeekly()
	if err != nil {
		ps.Logger.Error("Failed to get weekly promotions", zap.Error(err))
		return nil, err
	}
	return promotions, nil
}
