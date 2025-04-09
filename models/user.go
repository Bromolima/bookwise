package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/book-wise/secutiry"
)

var (
	MethodCreate = "create"
	MethodGet    = "get"
	MethodUpdate = "update"
	MethodLogin  = "login"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Passsword string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func (u *User) validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}

	if u.Username == "" {
		return errors.New("username is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	if err := checkmail.ValidateFormat(u.Email); err != nil {
		return err
	}

	if u.Passsword == "" {
		return errors.New("password is required")
	}

	return nil
}

func (u *User) trim(method string) error {
	u.Name = strings.TrimSpace(u.Name)
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)

	if method == MethodCreate {
		hashedPassword, err := secutiry.Hash(u.Passsword)
		if err != nil {
			return err
		}

		u.Passsword = string(hashedPassword)
	}

	return nil
}

func (u *User) Validate(method string) error {
	if err := u.validate(); err != nil {
		return err
	}

	if err := u.trim(method); err != nil {
		return err
	}

	return nil
}
