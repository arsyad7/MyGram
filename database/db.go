package database

import (
	"fmt"
	"log"
	"mygram/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	DB_HOST = "localhost"
	DB_PORT = "5432"
	DB_USER = "postgres"
	DB_PASS = "postgres"
	DB_NAME = "mygram"
)

var database *gorm.DB

func StartDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", DB_HOST, DB_PORT, DB_USER, DB_PASS, DB_NAME)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	log.Default().Println("Connection to db success")

	err = migration(db)
	if err != nil {
		panic(err)
	}

	database = db
	return db
}

func migration(db *gorm.DB) error {
	if err := db.AutoMigrate(models.User{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(models.Photo{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(models.Comment{}); err != nil {
		return err
	}
	if err := db.AutoMigrate(models.SocialMedia{}); err != nil {
		return err
	}
	return nil
}

func GetDB() *gorm.DB {
	return database
}
