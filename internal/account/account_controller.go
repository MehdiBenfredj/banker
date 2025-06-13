package account

import (
	"database/sql"
	"fmt"
)

type AccountController struct {
	Service *AccountService
}



func NewAccountController(service *AccountService) *AccountController {
	return &AccountController{Service: service}
}

func NewAccountModule(db *sql.DB) *AccountController {
	repo := NewAccountRepository(db)
	svc := NewAccountService(repo)
	return NewAccountController(svc)
}


func (a *AccountController) TestController() {
	fmt.Println("account controller ! called")
	a.Service.Repository.Test()
}