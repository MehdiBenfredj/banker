package controllers

import (
	"github.com/MehdiBenfredj/banker/internal/services"
)

type AccountController struct {
	Service *services.AccountService
}

func NewAccountController(Service *services.AccountService) *AccountController {
	return &AccountController{Service: Service}
}

func (controller *AccountController) Test() {
	controller.Service.Test()
}
