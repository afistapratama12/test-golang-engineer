package config

import (
	"fmt"
	"log"
	"oauthGoogleGo/user"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	err    = godotenv.Load()
	dbUser = os.Getenv("DB_USER")
	dbpass = os.Getenv("DB_PASS")
	dbHost = os.Getenv("DB_HOST")
	dbName = os.Getenv("DB_NAME")
)

func Connect() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbpass, dbHost, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&user.User{})

	return db
}
