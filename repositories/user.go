package repositories

import (
	"database/sql"

	"github.com/book-wise/models"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user models.User) (uint64, error) {
	statement, err := r.db.Prepare(
		"insert into user (user_name, username, email, password) values(?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}

	defer statement.Close()

	result, err := statement.Exec(user.Name, user.Username, user.Email, user.Passsword)
	if err != nil {
		return 0, err
	}

	LastID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(LastID), nil
}
