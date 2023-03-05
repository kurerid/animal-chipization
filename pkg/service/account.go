package service

import (
	"animal-chipization/models"
	"animal-chipization/pkg/repository"
)

type AccountService struct {
	repo repository.Account
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}

func (s *AccountService) GetById(id int) (*models.AccountGetByIdOutput, error) {
	return s.repo.GetById(id)
}

func (s *AccountService) Search(input *models.AccountSearchInput) (*models.AccountSearchOutput, error) {
	return s.repo.Search(input)
}

func (s *AccountService) Update(input *models.AccountUpdateInput) (*models.AccountUpdateOutput, error) {
	return s.repo.Update(input)
}

func (s *AccountService) Remove(id int) error {
	return s.repo.Remove(id)
}

func (s *AccountService) Exist(email string) (bool, error) {
	return s.repo.Exist(email)
}
