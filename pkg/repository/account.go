package repository

import (
	"animal-chipization/models"
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

func (r *AccountPostgres) GetById(id int) (*models.AccountGetByIdOutput, error) {
	var output models.AccountGetByIdOutput
	err := r.db.Get(&output, `SELECT "Id", "FirstName", "LastName", "Email" FROM "Account" WHERE "Id" = $1`, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return &output, nil
}

func (r *AccountPostgres) Search(input *models.AccountSearchInput) (*models.AccountSearchOutput, error) {
	var output models.AccountSearchOutput
	query := `SELECT * FROM "Account" WHERE "Email" ILIKE '%' || $1 || '%' OR "FirstName" ILIKE '%' || $2 || '%'
	OR "LastName" ILIKE '%' || $3 || '%' OFFSET $4 LIMIT $5`
	rows, err := r.db.Query(query, input.Email, input.FirstName, input.LastName, input.Offset, input.Limit)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	if !rows.Next() {
		return nil, nil
	}
	err = r.db.Get(&output, query, input.Email, input.FirstName, input.LastName, input.Offset, input.Limit)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return &output, nil
}

func (r *AccountPostgres) Update(input *models.AccountUpdateInput) (*models.AccountUpdateOutput, error) {
	return nil, nil
}

func (r *AccountPostgres) Remove(id int) error {
	return nil
}
