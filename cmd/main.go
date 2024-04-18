package main

import (
	"github.com/joho/godotenv"
	"gochat/internal/app"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	handler := app.Routes()
	server := app.NewServer(port, handler)

	err = server.ListenAndServe()

	log.Fatal(err)

}
