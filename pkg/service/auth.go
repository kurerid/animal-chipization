package service

import (
	"animal-chipization/models"
	"animal-chipization/pkg/repository"
)

type AuthorizationService struct {
	repo repository.Authorization
}

func NewAuthorizationService(repo repository.Authorization) *AuthorizationService {
	return &AuthorizationService{repo: repo}
}

func (s *AuthorizationService) SignUp(input *models.SignUpInput) (*models.SignUpOutput, error) {
	return s.repo.SignUp(input)
}

func (s *AuthorizationService) SignIn(email string, password string) error {
	return s.repo.SignIn(email, password)
}
