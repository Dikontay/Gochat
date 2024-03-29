package main

import (
	"github.com/joho/godotenv"
	"gochat/internal/handlers"
	"log"
	"net/http"
	"os"
)

func NewServer(port string, handler http.Handler) *http.Server {
	return &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	handler := handlers.Routes()
	server := NewServer(port, handler)

	err = server.ListenAndServe()

	log.Fatal(err)

}
