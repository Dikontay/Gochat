package logger

import (
	"fmt"
	"gochat/utils"
	"log"
	"net/http"
)

type Logger struct {
	Logger *log.Logger
}

func (log Logger) LogError(r *http.Request, err error) {
	log.Logger.Print(err)
}

type envelope map[string]any

func (log Logger) ErrorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	env := envelope{"error": message}
	// Write the response using the writeJSON() helper. If this happens to return an // error then log it, and fall back to sending the client an empty response with a // 500 Internal Server Error status code.
	err := utils.WriteJSON(w, status, env, nil)
	if err != nil {
		log.LogError(r, err)
		w.WriteHeader(500)
	}
}

func (log Logger) ServerErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.LogError(r, err)
	message := "the server encountered a problem and could not process your request"
	log.ErrorResponse(w, r, http.StatusInternalServerError, message)
}

func (log Logger) NotFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	log.ErrorResponse(w, r, http.StatusNotFound, message)
}

func (log Logger) MethodNotAllowedResponse(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("the %s method is not supported for this resource", r.Method)
	log.ErrorResponse(w, r, http.StatusMethodNotAllowed, message)
}

func (log Logger) BadRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	log.ErrorResponse(w, r, http.StatusBadRequest, err.Error())
}

func (log Logger) FailedValidationResponse(w http.ResponseWriter, r *http.Request, errors map[string]string) {
	log.ErrorResponse(w, r, http.StatusUnprocessableEntity, errors)
}
