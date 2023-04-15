package main

import (
	"finalproject_mygram/database"
	router "finalproject_mygram/routers"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	database.StartDB()
	r := router.StartApp()
	r.Run(":" + os.Getenv("PORT"))
}
