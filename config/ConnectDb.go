package config

import (
	"fmt"
	"log"
	"maca/entities"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDb() *gorm.DB {
	err := godotenv.Load()
	dbUn := os.Getenv("DB_UN")
	dbPw := os.Getenv("DB_PW")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUn, dbPw, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := db.AutoMigrate(
		&entities.Book{},
	); err != nil {
		return nil
	}

	return db
}
