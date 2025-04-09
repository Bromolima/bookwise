package repositories

import (
	"database/sql"
	"fmt"

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
		"insert into users (user_name, username, email, user_password) values(?, ?, ?, ?)",
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

func (r *userRepository) Get(nameOrUsername string) ([]models.User, error) {
	nameOrUsername = fmt.Sprintf("%%%s%%", nameOrUsername)

	rows, err := r.db.Query("select id, user_name, username, email, createdAt from users where user_name LIKE ? or username LIKE ?", nameOrUsername, nameOrUsername)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User

		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Email,
			&user.UpdatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) GetByID(ID uint64) (models.User, error) {
	rows, err := r.db.Query(
		"select id, user_name, username, email, createdAt from users where id = ?",
		ID,
	)

	if err != nil {
		return models.User{}, err
	}

	var user models.User

	if rows.Next() {
		if err = rows.Scan(
			&user.ID,
			&user.Name,
			&user.Username,
			&user.Username,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

func (r *userRepository) Update(ID uint64, user models.User) error {
	statement, err := r.db.Prepare(
		"update users set user_name = ?, username = ?, email = ?, where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Username, user.Email, ID); err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Delete(ID uint64) error {
	statement, err := r.db.Prepare("delete from users where id = ?")
	if err != nil {
		return err
	}

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

func (r *userRepository) GetByEmail(email string) (models.User, error) {
	row, err := r.db.Query("select id, user_password from users where email = ?", email)
	if err != nil {
		return models.User{}, err
	}
	defer row.Close()

	var user models.User

	if row.Next() {
		if err = row.Scan(&user.ID, &user.Passsword); err != nil {
			return models.User{}, err
		}
	}

	return user, err
}
