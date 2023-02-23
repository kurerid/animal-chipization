package repository

import (
	"animal-chipization/models"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) SignUp(input *models.SignUpInput) (*models.SignUpOutput, error) {
	var output models.SignUpOutput
	err := r.db.Get(&output, `INSERT INTO "Account"("FirstName", "LastName", "Email", "Password") VALUES ($1,$2,$3,$4) RETURNING "Id","FirstName","LastName","Email"`,
		input.FirstName, input.LastName, input.Email, input.Password)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return &output, nil
}

func (r *AuthPostgres) SignIn(email string, password string) error {
	return nil
}
