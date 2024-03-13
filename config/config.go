package config

import (
	"fmt"
	"log"
	"tugas__2/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "12345678"
	dbPort   = "5432"
	dbname   = "orders_by"
	db       *gorm.DB
	err      error
)

func StarDB() *gorm.DB {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database :", err)
	}
	defer fmt.Println("Successfully Connected to Database")
	db.Debug().AutoMigrate(model.Order{}, model.Item{})
	return db
}
