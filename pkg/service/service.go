package service

import "animal-chipization/pkg/repository"

type Authorization interface {
}

type Animal interface {
}

type User interface {
}

type Location interface {
}

type Service struct {
	Authorization
	Animal
	User
	Location
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: nil,
		Animal:        nil,
		User:          nil,
		Location:      nil,
	}
}
