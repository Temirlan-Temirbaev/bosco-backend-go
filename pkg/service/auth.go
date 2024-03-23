package service

import (
	"bosco-backend/pkg/model"
	"bosco-backend/pkg/repository"
	"bosco-backend/pkg/utils"
)

type AuthService struct {
	repository repository.Authorization
}

func NewAuthService(repository repository.Authorization) *AuthService {
	return &AuthService{repository: repository}
}

func (service *AuthService) CreateUser(user model.User) (int, error) {
	user.Password = utils.GeneratePasswordHash(user.Password)
	return service.repository.Create(user)
}
