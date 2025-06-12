package repositories

import (
	"database/sql"
	"fmt"
)

type AccountRepository struct {
	db *sql.DB
}

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db: db}
}

func (ar *AccountRepository) Test() {
	fmt.Println("account repository ! called")
	// make a test query to the database
	query := "SELECT 1"
	row := ar.db.QueryRow(query)
	if row.Err() != nil {
		fmt.Println("Error executing query:", row.Err())
		return
	}
	var result int
	err := row.Scan(&result)
	if err != nil {
		fmt.Println("Error scanning result:", err)
		return
	}
	fmt.Println("Test query result:", result)
	fmt.Println("Account repository test successful!")
}