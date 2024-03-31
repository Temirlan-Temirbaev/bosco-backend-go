package repository

import (
	"bosco-backend/internal/model"
	"bosco-backend/internal/repository/postgres"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	Create(user model.User) (int, error)
	GetUser(username string, password string) (model.User, error)
	GetUserById(id int) (model.User, error)
}

type Product interface {
}

type Category interface {
	Create(category model.Category) (int, error)
	Update(id int, category model.Category) error
	Delete(id int) error
	GetAll() ([]model.Category, error)
	GetById(id int) (model.Category, error)
}

type Contact interface {
	Create(contact model.Contact) (int, error)
	GetAll() ([]model.Contact, error)
	Delete(id int) error
	Update(id int, contact model.Contact) error
}

type Repository struct {
	Authorization
	Category
	Product
	Contact
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
		Contact:       postgres.NewContactPostgres(db),
		Category:      postgres.NewCategoryPostgres(db),
	}
}
