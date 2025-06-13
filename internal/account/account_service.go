package account

type AccountService struct {
	Repository *AccountRepository
}

func NewAccountService(repository *AccountRepository) *AccountService {
	return &AccountService{Repository: repository}
}


