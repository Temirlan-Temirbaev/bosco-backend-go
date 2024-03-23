package postgres

import (
	"bosco-backend/pkg/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (repository *AuthPostgres) Create(user model.User) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (username, password) values ($1, $2) RETURNING id", "users")
	row := repository.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
