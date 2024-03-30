package postgres

import (
	"bosco-backend/internal/constants"
	"bosco-backend/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type CategoryPostgres struct {
	db *sqlx.DB
}

func NewCategoryPostgres(db *sqlx.DB) *CategoryPostgres {
	return &CategoryPostgres{db: db}
}

func (repository *CategoryPostgres) Create(category model.Category) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (name) VALUES ($1) RETURNING id", constants.CATEGORIES)
	row := repository.db.QueryRow(query, category.Name)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (repository *CategoryPostgres) Update(id int, category model.Category) error {
	query := fmt.Sprintf("UPDATE %s SET name=$1 WHERE id=$2", constants.CATEGORIES)
	_, err := repository.db.Exec(query, category.Name, id)
	return err
}

func (repository *CategoryPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", constants.CATEGORIES)
	_, err := repository.db.Exec(query, id)
	return err
}

func (repository *CategoryPostgres) GetAll() ([]model.Category, error) {
	var categories []model.Category
	query := fmt.Sprintf("SELECT * FROM %s", constants.CATEGORIES)
	if err := repository.db.Select(&categories, query); err != nil {
		return []model.Category{}, err
	}

	return categories, nil
}
func (repository *CategoryPostgres) GetById(id int) (model.Category, error) {
	var category model.Category
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", constants.CATEGORIES)
	if err := repository.db.Get(&category, query, id); err != nil {
		return model.Category{}, err
	}

	return category, nil
}
