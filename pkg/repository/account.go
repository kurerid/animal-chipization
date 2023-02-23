package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AccountPostgres struct {
	db *sqlx.DB
}

func NewAccountPostgres(db *sqlx.DB) *AccountPostgres {
	return &AccountPostgres{db: db}
}

func (r *AccountPostgres) Exist(email string) (bool, error) {
	rows, err := r.db.Query(`SELECT "Email" FROM "Account" WHERE "Email" = $1`, email)
	if err != nil {
		logrus.Error(err)
		return false, err
	}
	if !rows.Next() {
		return false, nil
	}
	return true, nil
}
