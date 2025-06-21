package account

type AccountService struct {
	Repository *AccountRepository
}

func (accountService *AccountService) GetAllAccounts() ([]Account, error) {
	accounts, error := accountService.Repository.GetAllAccounts()
	if error != nil {
		return nil, error
	} else {
		return accounts, nil
	}
}

func (accountService *AccountService) GetAccountById(account_id string) (*Account, error) {
	return accountService.Repository.GetAccountById(account_id)
}
func (accountService *AccountService) CreateNewAccount(user_id string) error {
	return accountService.Repository.CreateNewAccount(user_id)
}
func (accountService *AccountService) DeleteAccount(account_id string) error {
	return accountService.Repository.DeleteAccount(account_id)
}

func NewAccountService(repository *AccountRepository) *AccountService {
	return &AccountService{Repository: repository}
}
