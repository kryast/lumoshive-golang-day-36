package service

import (
	"day-36/model"

	"go.uber.org/zap"
)

func (as *AuthService) GetAddressesByAuthID(authID int) ([]model.Address, error) {
	addresses, err := as.Repo.AuthRepo.GetAddressesByAuthID(authID)
	if err != nil {
		as.Logger.Error("Failed to get addresses", zap.Int("authID", authID), zap.Error(err))
		return nil, err
	}
	return addresses, nil
}
