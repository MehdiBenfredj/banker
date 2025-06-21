package main

import (
	"database/sql"
	_ "encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/MehdiBenfredj/banker/internal/account"
	"github.com/MehdiBenfredj/banker/internal/card"
	"github.com/MehdiBenfredj/banker/internal/user"

	_ "github.com/lib/pq"
)

// disatch depedning on route
func globalHandler(
	writer http.ResponseWriter,
	request *http.Request,
	accountController *account.AccountController,
	userController *user.UserController,
	cardController *card.CardController,
) {

	switch request.URL.Path {
	case "/user":
		userController.Route(writer, request)
	case "/account":
		accountController.Route(writer, request)
	case "/card":
		//	controllers.CardController(writer, request)
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

	accountController := account.NewAccountModule(db)
	cardController := card.NewCardModule(db)
	userController := user.NewUserModule(db)

	http.HandleFunc("/", func(responseWriter http.ResponseWriter, request *http.Request) {
		globalHandler(responseWriter, request, accountController, userController, cardController)
	})
	fmt.Println("Listening on port 8080")
	http.ListenAndServe(":8080", nil)

}
