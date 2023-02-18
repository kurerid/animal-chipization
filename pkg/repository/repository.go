package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Animal interface {
}

type User interface {
}

type Location interface {
}

type Repository struct {
	Authorization
	Animal
	User
	Location
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: nil,
		Animal:        nil,
		User:          nil,
		Location:      nil,
	}
}
