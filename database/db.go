package database

import (
	"finalproject_mygram/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db  *gorm.DB
	err error
)

func StartDB() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	var (
		host     = os.Getenv("PGHOST")
		port     = os.Getenv("PGPORT")
		user     = os.Getenv("PGUSER")
		password = os.Getenv("PGPASSWORD")
		dbname   = os.Getenv("PGDATABASE")
	)
	config := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database:", err)
	}

	db.Debug().AutoMigrate(models.User{}, models.Comment{}, models.Photo{}, models.SocialMedia{})
}

func GetDB() *gorm.DB {
	return db
}
