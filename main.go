package main

import (
	"database/sql"
	_ "encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MehdiBenfredj/banker/internal/controllers"
	"github.com/MehdiBenfredj/banker/internal/repositories"
	"github.com/MehdiBenfredj/banker/internal/services"

	_ "github.com/lib/pq"
)

// disatch depedning on route
func globalHandler(writer http.ResponseWriter, request *http.Request, db *sql.DB) {
	accountRepo := repositories.NewAccountRepository(db)
	accountService := services.NewAccountService(accountRepo)
	accountController := controllers.NewAccountController(accountService)

 
	switch request.URL.Path {
	case "/user":
		controllers.UserController(writer, request)
	case "/account":
		//call the account controller
		accountController.Test()
	case "/card":
		controllers.CardController(writer, request)
	}
}

func main() {
	connStr := "host=localhost port=5432 user=mehdi password=1234 dbname=godb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening DB:", err)
	}
	defer db.Close()

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatal("Cannot connect to DB:", err)
	}

	fmt.Println("Connected to PostgreSQL!")

	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		globalHandler(responseWriter, request, db)
	})
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)

}
