package handlers

import (
	"errors"
	"gochat/internal/models"
	user2 "gochat/internal/repository/user"
	"gochat/utils"
	"net/http"
	"time"
)

func (app *App) CreateAuthenticationToken(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := utils.ReadJSON(w, r, &input)

	if err != nil {
		app.Logger.BadRequestResponse(w, r, err)
		return
	}

	user, err := app.Repo.User.GetByEmail(input.Email)
	if err != nil {
		switch {
		case errors.Is(err, user2.ErrRecordNotFound):
			app.Logger.BadRequestResponse(w, r, err)
		default:
			app.Logger.ServerErrorResponse(w, r, err)
		}
		return
	}

	match, err := user.Password.Matches(input.Password)
	if err != nil {
		app.Logger.ServerErrorResponse(w, r, err)
		return
	}
	if !match {
		app.Logger.BadRequestResponse(w, r, err)
		return
	}
	token, err := app.Repo.Token.New(user.ID, 24*time.Hour, models.ScopeAuthentication)
	if err != nil {
		app.Logger.ServerErrorResponse(w, r, err)
		return
	}
	err = utils.WriteJSON(w, http.StatusCreated, map[string]interface{}{"authentication_token": token}, nil)
	if err != nil {
		app.Logger.ServerErrorResponse(w, r, err)
		return
	}
}
