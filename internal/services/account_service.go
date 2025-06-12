package services

import (
	"github.com/MehdiBenfredj/banker/internal/repositories"
)

type AccountService struct {
	Repository *repositories.AccountRepository
}

func NewAccountService(repository *repositories.AccountRepository) *AccountService {
	return &AccountService{Repository: repository}
}

func (service *AccountService) Test() {
	service.Repository.Test()
}
