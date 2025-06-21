package account

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type AccountController struct {
	AccountService *AccountService
}

func NewAccountController(service *AccountService) *AccountController {
	return &AccountController{AccountService: service}
}

func NewAccountModule(db *sql.DB) *AccountController {
	repo := NewAccountRepository(db)
	svc := NewAccountService(repo)
	return NewAccountController(svc)
}

func (accountController *AccountController) Route(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		if !request.URL.Query().Has("user_id") {
			http.Error(writer, "Missing user_id", http.StatusBadRequest)
			return
		}
		user_id := request.URL.Query().Get("user_id")
		if error := accountController.AccountService.CreateNewAccount(user_id); error != nil {
			http.Error(writer, "Missing user_id", http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "text/plain")
		writer.Write([]byte("Account created successfully for: " + user_id))
	case http.MethodGet:
		if request.URL.Query().Has("account_id") {
			account_id := request.URL.Query().Get("account_id")
			account, error := accountController.AccountService.GetAccountById(account_id)
			if error != nil {
				http.Error(writer, "Could not get user : " + error.Error(), http.StatusInternalServerError)
			}
			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(account)
			return
		} else {
			accounts, error := accountController.AccountService.GetAllAccounts()
			if error != nil {
				http.Error(writer, "Could not retrieve accounts", http.StatusNotFound)
				return
			}
			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(accounts)
			return
		}
	case http.MethodDelete:
		if !request.URL.Query().Has("account_id") {
			http.Error(writer, "Missing account_id", http.StatusBadRequest)
			return
		}
		account_id := request.URL.Query().Get("account_id")
		if error := accountController.AccountService.DeleteAccount(account_id); error != nil {
			http.Error(writer, "Could not delete account, account_id : " + account_id, http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "text/plain")
		writer.Write([]byte("Account deleted succesfully!"))
	}
}


