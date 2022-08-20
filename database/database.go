package database

import (
	"os"
	"fmt"
	"github.com/joho/godotenv"
  "gorm.io/driver/postgres"
  "gorm.io/gorm"
)

func Connect() *gorm.DB {

	err := godotenv.Load("database/.env")
	
	if err != nil {
		fmt.Printf("Could not load the .env file: %v", err)
	} 

	db_host := os.Getenv("DB_HOST")
	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	db_port := os.Getenv("DB_PORT")
	db_sslmode := os.Getenv("DB_SSLMODE")
	db_timezone := os.Getenv("DB_TIMEZONE")

  dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", db_host, db_user, db_password, db_name, db_port, db_sslmode, db_timezone)
  db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn, 
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		fmt.Printf("Connection with database failed: %v", err)
	} 
	return db
}
