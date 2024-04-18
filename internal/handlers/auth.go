package handlers

import (
	"gochat/utils"
	"net/http"
)

func CreateAuthenticationToken(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	err := utils.ReadJSON(w, r, &input)

	if err != nil {

	}
}
