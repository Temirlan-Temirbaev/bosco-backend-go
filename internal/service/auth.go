package service

import (
	"bosco-backend/internal/constants"
	"bosco-backend/internal/model"
	"bosco-backend/internal/repository"
	"bosco-backend/internal/utils"
	"errors"
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

func (service *AuthService) GetUserById(id int) (model.User, error) {
	return service.repository.GetUserById(id)
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

func (service *AuthService) GetIdFromToken(accessToken string) (int, error) {
	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("Invalid signing method")
		}

		return []byte(constants.SigningKey), nil
	})
	if err != nil {
		return 0, nil
	}
	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("Not correct type of JWT token")
	}

	return claims.Id, nil
}
