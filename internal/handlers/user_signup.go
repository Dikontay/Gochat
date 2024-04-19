package handlers

import (
	"errors"
	"gochat/internal/models"
	user2 "gochat/internal/repository/user"
	"gochat/utils"
	"net/http"
)

func (app *App) SignUp(w http.ResponseWriter, r *http.Request) {
	var user models.User
	var input struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	err := utils.ReadJSON(w, r, &input)
	if err != nil {
		app.Logger.BadRequestResponse(w, r, err)
	}
	user.Username = input.Username
	user.Email = input.Email
	user.Activated = false

	err = user.Password.Set(input.Password)
	if err != nil {
		app.Logger.ServerErrorResponse(w, r, err)
		return
	}

	err = app.Repo.User.CreateUser(&user)
	if err != nil {
		switch {
		case errors.Is(err, user2.ErrDuplicateEmail):
			app.Logger.FailedValidationResponse(w, r, map[string]string{"err": "duplicate email"})
		default:
			app.Logger.ServerErrorResponse(w, r, err)
		}
		return
	}

	err = utils.WriteJSON(w, http.StatusCreated, map[string]interface{}{"user": user}, nil)
	if err != nil {
		app.Logger.ServerErrorResponse(w, r, err)
	}

}

func (app *App) Login(w http.ResponseWriter, r *http.Request) {

}
