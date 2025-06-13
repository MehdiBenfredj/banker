package user

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

type UserController struct {
	Service *UserService
}

func NewUserModule(db *sql.DB) *UserController {
	repo := NewUserRepository(db)
	svc := NewUserService(repo)
	return NewUserController(svc)
}

func NewUserController(service *UserService) *UserController {
	return &UserController{Service: service}
}

// CRUD operations dispatch to the service layer
func (c *UserController) CreateUser(firstName, lastName, dateOfBirth, placeOfBirth, address string) error {
	return c.Service.CreateUser(firstName, lastName, dateOfBirth, placeOfBirth, address)
}

func (c *UserController) GetUserByID(id int) (User, error) {
	user, err := c.Service.GetUserByID(id)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (c *UserController) GetUserByLastName(lastName string) ([]User, error) {
	users, err := c.Service.GetUserByLastName(lastName)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (c *UserController) UpdateUser(id int, firstName, lastName, dateOfBirth, placeOfBirth, address string) error {
	return c.Service.UpdateUser(id, firstName, lastName, dateOfBirth, placeOfBirth, address)
}

func (c *UserController) DeleteUser(id int) error {
	return c.Service.DeleteUser(id)
}

func (c *UserController) GetAllUsers() ([]User, error) {
	users, err := c.Service.GetAllUsers()
	if err != nil {
		// Log the error or handle it as needed
		// For example, you can log it or return a specific error message
		log.Printf("Error getting all users: %v", err)
		return nil, err
	}
	return users, nil
}

func (c *UserController) Route(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		c.CreateUser(request.FormValue("first_name"),
			request.FormValue("last_name"),
			request.FormValue("date_of_birth"),
			request.FormValue("place_of_birth"),
			request.FormValue("address"))
		writer.WriteHeader(http.StatusCreated)
		writer.Write([]byte("User created successfully"))
	case http.MethodGet:
		if request.URL.Query().Has("user_id") {
			id := request.URL.Query().Get("user_id")
			userID, err := strconv.Atoi(id)
			if err != nil {
				http.Error(writer, "Invalid user ID", http.StatusBadRequest)
				return
			}
			user, err := c.GetUserByID(userID)
			if err != nil {
				http.Error(writer, "User not found", http.StatusNotFound)
				return
			}
			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(user)
			return
		}
		if request.URL.Query().Has("last_name") {
			lastName := request.URL.Query().Get("last_name")
			users, err := c.GetUserByLastName(lastName)
			if err != nil {
				log.Printf("Error getting users by last name: %v", err)
				http.Error(writer, "Failed to get users", http.StatusInternalServerError)
				return
			}
			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(users)
			return
		}
		users, err := c.GetAllUsers()
		if err != nil {
			http.Error(writer, "Failed to get users", http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writer).Encode(users)
		return
	case http.MethodPut:
		// Handle user update
	case http.MethodDelete:
		// Handle user deletion
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return

	}
}
