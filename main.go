package main

import (
	"finalproject_mygram/database"
	router "finalproject_mygram/routers"
	"os"
)

func main() {
	database.StartDB()
	r := router.StartApp()
	r.Run(":" + os.Getenv("PORT"))
}
