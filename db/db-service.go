package db

import (
	"fmt"
	"go-api-poc/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// TODO: move database config
var (
	dbUser              = "go"
	dbPassword          = "password"
	dbHost              = "localhost"
	dbPort              = "5000"
	dbName              = "go_db"
	dsn                 = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName) //"go:password@tcp(127.0.0.1:5000)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	dbConnection        *gorm.DB
	shouldMigrateTables = true
)

func ConnectToDb() error {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
		return err
	}
	dbConnection = db
	log.Println("Connected to " + dbName + " database on " + dbHost + ":" + dbPort + " as user " + dbUser)

	if shouldMigrateTables {
		autoMigrateTables()
	}
	return nil
}

func autoMigrateTables() {
	log.Println("Auto migrating tables...")
	dbConnection.AutoMigrate(models.GetModels()...)
	log.Println("Auto migration complete")
}
