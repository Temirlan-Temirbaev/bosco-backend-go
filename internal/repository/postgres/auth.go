package postgres

import (
	"bosco-backend/internal/constants"
	"bosco-backend/internal/model"
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
	query := fmt.Sprintf("INSERT INTO %s (username, password) values ($1, $2) RETURNING id", constants.USERS)
	row := repository.db.QueryRow(query, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
func (repository *AuthPostgres) GetUser(username string, password string) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id, username, role FROM %s WHERE username=$1 AND password=$2", constants.USERS)
	row := repository.db.QueryRow(query, username, password)
	if err := row.Scan(&user.Id, &user.Username, &user.Role); err != nil {
		return model.User{}, err
	}
	return user, nil
}

func (repository *AuthPostgres) GetUserById(id int) (model.User, error) {
	var user model.User
	query := fmt.Sprintf("SELECT id, username, role FROM %s WHERE id=$1", constants.USERS)
	row := repository.db.QueryRow(query, id)
	if err := row.Scan(&user.Id, &user.Username, &user.Role); err != nil {
		return model.User{}, err
	}
	return user, nil
}
