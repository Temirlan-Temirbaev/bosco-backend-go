package postgres

import (
	"bosco-backend/internal/constants"
	"bosco-backend/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type ContactPostgres struct {
	db *sqlx.DB
}

func NewContactPostgres(db *sqlx.DB) *ContactPostgres {
	return &ContactPostgres{db: db}
}

func (repository ContactPostgres) Create(contact model.Contact) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (coordinates, phone, vip_phone, address) VALUES ($1, $2, $3, $4) RETURNING id", constants.CONTACTS)

	row := repository.db.QueryRow(query, contact.Coordinates, contact.Phone, contact.VipPhone, contact.Address)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (repository ContactPostgres) GetAll() ([]model.Contact, error) {
	var contacts []model.Contact
	query := fmt.Sprintf("SELECT * FROM %s", constants.CONTACTS)
	err := repository.db.Select(&contacts, query)
	return contacts, err
}

func (repository ContactPostgres) Delete(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", constants.CONTACTS)
	_, err := repository.db.Exec(query, id)
	return err
}
