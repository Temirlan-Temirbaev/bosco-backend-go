package service

import (
	"bosco-backend/internal/constants"
	"bosco-backend/internal/model"
	"bosco-backend/internal/repository"
	"bosco-backend/internal/utils"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
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

func (service *AuthService) GetUser(user model.User) (model.User, error) {
	return service.repository.GetUser(user.Username, utils.GeneratePasswordHash(user.Password))
}

type tokenClaims struct {
	jwt.StandardClaims
	Id       int    `json:"id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

func (service *AuthService) GenerateToken(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(constants.TokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.Id, user.Username, user.Role,
	})

	fmt.Println(token)

	return token.SignedString([]byte(constants.SigningKey))
}
