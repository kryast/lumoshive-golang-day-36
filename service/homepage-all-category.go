package service

import (
	"day-36/model"

	"go.uber.org/zap"
)

func (ps *ProductService) GetAllCategory() ([]*model.Category, error) {
	categories, err := ps.Repo.ProductRepo.GetAllCategory()
	if err != nil {
		ps.Logger.Error("Error fetching products", zap.Error(err))
		return nil, err
	}

	return categories, nil
}
