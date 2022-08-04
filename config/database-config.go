package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/luizsouz404/api-rest-go/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//SetupDatabase Connection is creating a new connection to our database
func SetupDatabaseConnection() *gorm.DB {
	err := godotenv.Load()

	if err != nil {
		panic("Failed to load env file!")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbHost := os.Getenv("DB_HOST")
	dbName := os.Getenv("DB_NAME")

	databaseConnection := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)

	db, err := gorm.Open(mysql.Open(databaseConnection), &gorm.Config{})

	if err != nil {
		panic("Failed to create a connection to database")
	}

	db.AutoMigrate(&entity.User{})
	return db
}

// CloseDatabaseConnection is closing a connection between our application to our database
func CloseDatabaseConnection(db *gorm.DB) {
	dbSQL, err := db.DB()

	if err != nil {
		panic("Failed to close connection from database")
	}

	dbSQL.Close()
}
