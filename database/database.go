package database

import (
	"os"
	"fmt"
  "gorm.io/driver/mysql"
  "gorm.io/gorm"
	"github.com/joho/godotenv"
)

func Connect() *gorm.DB {

	err := godotenv.Load("./database/.env")
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	} 
	

	db_user := os.Getenv("DB_USER")
	db_password := os.Getenv("DB_PASSWORD")
	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_name := os.Getenv("DB_NAME")

  dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", db_user, db_password, db_host, db_port, db_name)
  db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Printf("Connection with database failed: %v", err)
	} 
	return db
}
