package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() *gorm.DB {
	// dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	// 	os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))

	dsn := "host=localhost port=5432 user=root password=anangs123 sslmode=disable dbname=go_apps"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db.Migrator().AutoMigrate()
	if err != nil {
		log.Printf("error opening database: %v", err)
	}
	return db
}
