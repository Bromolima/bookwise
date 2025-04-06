package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Username  string    `json:"username,omitempty"`
	Email     string    `json:"email,omitempty"`
	Passsword string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
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

func (u *User) trim() {
	u.Name = strings.TrimSpace(u.Name)
	u.Username = strings.TrimSpace(u.Username)
	u.Email = strings.TrimSpace(u.Email)
	u.Passsword = strings.TrimSpace(u.Passsword)
}

func (u *User) Validate() error {
	if err := u.validate(); err != nil {
		return err
	}

	u.trim()
	return nil
}
