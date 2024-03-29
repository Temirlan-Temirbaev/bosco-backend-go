package postgres

import (
	"bosco-backend/internal/constants"
	"bosco-backend/internal/model"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type ContactPostgres struct {
	db *sqlx.DB
}

func NewContactPostgres(db *sqlx.DB) *ContactPostgres {
	return &ContactPostgres{db: db}
}

func (repository ContactPostgres) Create(contact model.Contact) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (coordinates, phone, vip_phone, address) VALUES ($1, $2, $3, $4) RETURNING id", constants.CONTACTS)

	coordinatesLiteral := "{" + strings.Join(contact.Coordinates, ",") + "}"
	addressLiteral := "{" + strings.Join(contact.Address, ",") + "}"

	var id int
	err := repository.db.QueryRow(query, coordinatesLiteral, contact.Phone, contact.VipPhone, addressLiteral).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
