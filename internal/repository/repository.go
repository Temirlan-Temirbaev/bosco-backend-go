package repository

import (
	"bosco-backend/internal/model"
	"bosco-backend/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	Create(user model.User) (int, error)
	GetUser(username string, password string) (model.User, error)
}

type Product interface {
}

type Category interface {
}

type Contact interface {
}

type Repository struct {
	Authorization
	Category
	Product
	Contact
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{Authorization: postgres.NewAuthPostgres(db)}
}