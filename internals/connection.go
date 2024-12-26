package internals

import (
	"fmt"
	"os"
	"shopping-site/pkg/loggers"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Pg *gorm.DB

func InitiatePgConnection() *gorm.DB {
	dsn := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=disable", os.Getenv("USER"), os.Getenv("PASSWORD"), os.Getenv("HOST"), os.Getenv("PORT"), os.Getenv("DBNAME"))

	client, err := gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true})
	if err != nil {
		loggers.InfoLog.Fatalf("Failed to open postgres client %v", err)
	}

	loggers.InfoLog.Print("Connected to postgress client")
	fmt.Println("Connected to postgress client")

	Pg = client

	return client
}
