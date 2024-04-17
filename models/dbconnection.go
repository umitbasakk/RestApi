package models

import (
	"database/sql"
	_ "github.com/lib/pq"
)

var databases = &sql.DB{}

func init() {
	/*
		godotenv.Load(".env")

		username := os.Getenv("user")
		password := os.Getenv("password")
		port := os.Getenv("dbport")
		dbname := os.Getenv("dbname")
		host := os.Getenv("host")

		dbConnect := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
		db, err := sql.Open("postgres", dbConnect)

		if err != nil {
			fmt.Println("Error")
		}
		databases = db
		err = db.Ping()

		if err != nil {
			log.Panic("Panic...")
		}
		defer db.Close()
	*/
}

func GetDb() *sql.DB {
	return databases
}
