package user

import (
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

// CDUD Create, Delete, Update, and Display
func (r *UserRepository) CreateUser(firstName, lastName, dateOfBirth, placeOfBirth, address string) error {
	query := `INSERT INTO users (first_name, last_name, date_of_birth, place_of_birth, address) 
			  VALUES ($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(query, firstName, lastName, dateOfBirth, placeOfBirth, address)
	return err
}

func (r *UserRepository) GetUserByID(user_id int) (User, error) {
	query := `SELECT *
			  FROM users WHERE user_id = $1`
	row := r.db.QueryRow(query, user_id)

	var user User
	err := row.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.DateOfBirth, &user.PlaceOfBirth, &user.Address)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func (r *UserRepository) GetUserByLastName(lastName string) ([]User, error) {
	query := `SELECT *
			  FROM users WHERE last_name = $1`
	rows, err := r.db.Query(query, lastName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.DateOfBirth, &user.PlaceOfBirth, &user.Address); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) UpdateUser(user_id int, firstName, lastName, dateOfBirth, placeOfBirth, address string) error {
	query := `UPDATE users 
			  SET first_name = $1, last_name = $2, date_of_birth = $3, place_of_birth = $4, address = $5 
			  WHERE user_id = $6`
	_, err := r.db.Exec(query, firstName, lastName, dateOfBirth, placeOfBirth, address, user_id)
	return err
}

func (r *UserRepository) DeleteUser(user_id int) error {
	query := `DELETE FROM users WHERE user_id = $1`
	_, err := r.db.Exec(query, user_id)
	return err
}

func (r *UserRepository) GetAllUsers() ([]User, error) {
	query := `SELECT * FROM users`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.UserID, &user.FirstName, &user.LastName, &user.DateOfBirth, &user.PlaceOfBirth, &user.Address); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
