package service

import (
	"animal-chipization/models"
	"animal-chipization/pkg/repository"
)

type LocationService struct {
	repo repository.Location
}

func NewLocationService(repo repository.Location) *LocationService {
	return &LocationService{repo: repo}
}

func (s *LocationService) GetById(id int64) (*models.Location, error) {
	return s.repo.GetById(id)
}

func (s *LocationService) New(latitude float32, longitude float32) error {
	return s.repo.New(latitude, longitude)
}

func (s *LocationService) Modify(id int64, latitude float32, longitude float32) error {
	return s.repo.Modify(id, latitude, longitude)
}

func (s *LocationService) Remove(id int64) error {
	return s.repo.Remove(id)
}
