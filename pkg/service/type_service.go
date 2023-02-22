package service

import (
	"animal-chipization/models"
	"animal-chipization/pkg/repository"
)

type TypeService struct {
	repo repository.Type
}

func NewTypeService(repo repository.Type) *TypeService {
	return &TypeService{repo: repo}
}

func (s *TypeService) GetById(id int64) (*models.AnimalType, error) {
	return s.repo.GetById(id)
}

func (s *TypeService) New(name string) error {
	return s.repo.New(name)
}

func (s *TypeService) Modify(id int64, name string) error {
	return s.repo.Modify(id, name)
}

func (s *TypeService) Remove(id int) error {
	return s.repo.Remove(id)
}
