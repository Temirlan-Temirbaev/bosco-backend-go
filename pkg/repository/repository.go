package repository

import (
	"bosco-backend/pkg/model"
	"bosco-backend/pkg/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	Create(user model.User) (int, error)
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
