package user

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func (c *UserController) GetUserByID(user_id string) (User, error) {
	user, err := c.Service.GetUserByID(user_id)
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

func (c *UserController) UpdateUser(user_id string, firstName, lastName, dateOfBirth, placeOfBirth, address string) error {
	return c.Service.UpdateUser(user_id, firstName, lastName, dateOfBirth, placeOfBirth, address)
}

func (c *UserController) DeleteUser(user_id string) error {
	return c.Service.DeleteUser(user_id)
}

func (c *UserController) GetAllUsers() ([]User, error) {
	users, err := c.Service.GetAllUsers()
	if err != nil {
		log.Printf("Error getting all users: %v", err)
		return nil, err
	}
	return users, nil
}

func (c *UserController) Route(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case http.MethodPost:
		var newUser User
		err := json.NewDecoder(request.Body).Decode(&newUser)
		fmt.Println(newUser)
		if err != nil {
			http.Error(writer, "Invalid request body", http.StatusBadRequest)
			return
		}
		err = c.CreateUser(
			newUser.FirstName,
			newUser.LastName,
			newUser.DateOfBirth,
			newUser.PlaceOfBirth,
			newUser.Address,
		)
		if err != nil {
			http.Error(writer, "Could not add user", http.StatusInternalServerError)
			log.Fatal(err)
			return
		}
		writer.WriteHeader(http.StatusCreated)
		writer.Write([]byte("User created successfully"))
		return

	case http.MethodGet:
		if request.URL.Query().Has("user_id") {
			userID := request.URL.Query().Get("user_id")
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
		userID := request.FormValue("user_id")
		if userID == "" {
			http.Error(writer, "User ID is required", http.StatusBadRequest)
			return
		}
		err := c.UpdateUser(
			userID,
			request.FormValue("first_name"),
			request.FormValue("last_name"),
			request.FormValue("date_of_birth"),
			request.FormValue("place_of_birth"),
			request.FormValue("address"),
		)
		if err != nil {
			http.Error(writer, "Failed to update user", http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("User updated successfully"))
	case http.MethodDelete:
		// Handle user deletion
		userID := request.FormValue("user_id")
		if userID == "" {
			http.Error(writer, "Invalid user ID", http.StatusInternalServerError)
			return
		}
		err := c.DeleteUser(userID)
		if err != nil {
			http.Error(writer, "Could not delete user!", http.StatusInternalServerError)
			return
		}
		error := c.DeleteUser(userID)
		if error != nil {
			http.Error(writer, "Could not delete user!", http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
		writer.Write([]byte("User deleted successfully"))
		return
	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
		return

	}
}
