package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *App) Routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/login", app.Login)
	router.HandlerFunc(http.MethodPost, "/signup", app.SignUp)
	//router.HandlerFunc(http.MethodPost, "/tokens/authentication", CreateAuthenticationToken)

	return router

}
