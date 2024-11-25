package service

import (
	"day-36/model"

	"go.uber.org/zap"
)

func (as *AuthService) UpdateProfile(id int, profile *model.Profile) (*model.Profile, error) {
	updatedProfile, err := as.Repo.AuthRepo.UpdateProfile(id, profile)
	if err != nil {
		as.Logger.Error("Failed to update profile", zap.Int("id", id), zap.Error(err))
		return nil, err
	}

	return updatedProfile, nil
}
