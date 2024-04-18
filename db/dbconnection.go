package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"log"
	"os"
)

var databases = &gorm.DB{}

func init() {

	godotenv.Load(".env")

	username := os.Getenv("user")
	password := os.Getenv("password")
	port := os.Getenv("dbport")
	dbname := os.Getenv("dbname")
	host := os.Getenv("host")

	dbConnect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	db, err := gorm.Open("postgres", dbConnect)

	if err != nil {
		log.Panic("Panic...")
	}
	databases = db

}

func GetDb() *gorm.DB {
	return databases
}
