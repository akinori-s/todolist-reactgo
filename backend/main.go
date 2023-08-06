package main

import (
	"log"
	"todolist/app"
	"todolist/database"
)

func main() {
	db, err := database.InitDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	application := app.NewApplication(db)
	err = application.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to run application: %v", err)
	}
}
