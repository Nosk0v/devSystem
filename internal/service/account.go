package service

import (
	"devSystem/internal/repository"
	"devSystem/models"
)

type AccountService struct {
	repo repository.Account
}

func (a *AccountService) Get(email string) (*models.Account, error) {
	return a.repo.Get(email)
}

func NewAccountService(repo repository.Account) *AccountService {
	return &AccountService{repo: repo}
}
