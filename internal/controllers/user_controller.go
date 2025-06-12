package controllers

import (
	"fmt"
	"net/http"
)

func UserController(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("User")
}
