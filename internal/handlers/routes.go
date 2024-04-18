package handlers

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type Handlers struct {
	Handlers *httprouter.Router
}

func NewHandler() *httprouter.Router {
	return Routes()
}

func Routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/login", Login)
	router.HandlerFunc(http.MethodPost, "/signup", SignUp)
	router.HandlerFunc(http.MethodPost, "/tokens/authentication", CreateAuthenticationToken)
	
	return router

}
