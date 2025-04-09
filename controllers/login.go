package controllers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/book-wise/auth"
	"github.com/book-wise/database"
	"github.com/book-wise/models"
	"github.com/book-wise/repositories"
	"github.com/book-wise/responses"
	"github.com/book-wise/secutiry"
)

func Login(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, err)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewUserRepository(db)
	savedUser, err := repository.GetByEmail(user.Email)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}

	if err = secutiry.VerifyPassword(savedUser.Passsword, user.Passsword); err != nil {
		responses.Err(w, http.StatusUnauthorized, err)
		return
	}

	token, err := auth.GenerateToken(savedUser.ID)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, err)
		return
	}
	w.Write([]byte(token))

}
