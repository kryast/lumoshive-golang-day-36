package service

import (
	"day-36/model"

	"go.uber.org/zap"
)

func (ps *ProductService) GetAllCart() ([]*model.Cart, error) {
	carts, err := ps.Repo.ProductRepo.GetAllCart()
	if err != nil {
		ps.Logger.Error("Error retrieving cart items", zap.Error(err))
		return nil, err
	}
	return carts, nil
}
