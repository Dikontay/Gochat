package main

import (
	"gochat/internal/handlers"
	"gochat/pkg/db"
	"log"
)

func main() {
	database, err := db.NewDatabase()

	if err != nil {
		log.Fatal(err)
	}
	application := handlers.NewApp(database.GetDB())

	err = application.Run()
	if err != nil {
		log.Fatal(err)
	}

}
