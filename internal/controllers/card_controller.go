package controllers

import (
	"fmt"
	"net/http"
)

func CardController(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Card")
}