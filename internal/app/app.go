package app

import (
	"gochat/internal/handlers"
	"gochat/internal/repository"
	"log"
)

type App struct {
	Logger   *log.Logger
	Repo     repository.Repository
	Handlers *handlers.Handlers
}
