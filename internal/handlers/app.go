package handlers

import (
	"database/sql"
	"fmt"
	"gochat/internal/logger"
	"gochat/internal/repository"
	"log"
	"net/http"
	"time"
)

type App struct {
	Logger *logger.Logger
	Repo   repository.Repository
}

func NewApp(db *sql.DB) *App {
	logg := logger.Logger{Logger: &log.Logger{}}
	repo := repository.NewRepo(db)

	return &App{
		Logger: &logg,
		Repo:   *repo,
	}
}

func (a *App) Run() error {
	server := &http.Server{
		Addr:         ":8080",
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      a.Routes(),
	}
	fmt.Println("listen and serve error")
	return server.ListenAndServe()
}
