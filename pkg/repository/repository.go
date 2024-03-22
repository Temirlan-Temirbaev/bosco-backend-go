package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Repository{}
}
