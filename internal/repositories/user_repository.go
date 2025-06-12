package repositories

import (
	"database/sql"
	"github.com/MehdiBenfredj/banker/internal/model"
)
type UserRepository struct {
    db *sql.DB
}

func (r *UserRepository) GetUserByID(id string) (*model.User, error) {
    var user model.User
    err := r.db.QueryRow("SELECT user_id, first_name FROM users WHERE user_id = $1", id).
        Scan(&user.UserID, &user.FirstName)
    if err != nil {
        return nil, err
    }
    return &user, nil
}

func (r *UserRepository) GetUserByUserName(userName string) (*model.User, error) {
    var user model.User
    err := r.db.QueryRow("SELECT * FROM users WHERE user_name = $1", userName).
        Scan(&user.UserID, &user.FirstName)
    if err != nil {
        return nil, err
    }
    return &user, nil
}