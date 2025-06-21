package account

import (
	"database/sql"
	"os"

	"github.com/jacoelho/banking/iban"
	"github.com/joho/godotenv"
)

type AccountRepository struct {
	db *sql.DB
}

func (accountRepository *AccountRepository) GetAllAccounts() ([]Account, error) {
	query := "SELECT * from accounts"
	rows, error := accountRepository.db.Query(query)
	if error != nil {
		return nil, error
	}
	defer rows.Close()

	var accounts []Account
	for rows.Next() {
		var account Account
		if error = rows.Scan(&account.AccountID, &account.UserID, &account.BIC, &account.IBAN); error != nil {
			return nil, error
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

func (accountRepository *AccountRepository) GetAccountById(account_id string) (*Account, error) {
	query := "SELECT * FROM 	accounts WHERE account_id = $1"
	row := accountRepository.db.QueryRow(query, account_id)

	var account Account
	if err := row.Scan(&account.AccountID, &account.UserID, &account.BIC, &account.IBAN); err != nil {
		return nil, err
	}
	return &account, nil
}

func (accountRepository *AccountRepository) CreateNewAccount(user_id string) error {
	error := godotenv.Load()
	if error != nil {
		return error
	}
	BIC := os.Getenv("BIC")
	IBAN, err := iban.Generate("FR")
	if err != nil {
		return err
	}
	query := "INSERT INTO accounts (user_id, bic, iban) VALUES ($1, $2, $3)"
	_, err = accountRepository.db.Exec(query, user_id, BIC, IBAN)
	return err
}

func (accountRepository *AccountRepository) DeleteAccount(account_id string) error {
	query := "DELETE FROM accounts WHERE account_id = $1"
	_, error := accountRepository.db.Exec(query, account_id);
	return error
}
 
func NewUAccountModule(db *sql.DB) *AccountController {
	repo := NewAccountRepository(db)
	svc := NewAccountService(repo)
	return NewAccountController(svc)
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}
