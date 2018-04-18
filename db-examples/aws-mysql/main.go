package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	sqlInfo := os.Getenv("SQL_USER") + ":" +
		os.Getenv("SQL_PASSWORD") + "@tcp(" +
		os.Getenv("SQL_STRING") + ")/"
	db, err := sql.Open("mysql", sqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		fmt.Println(err)
	}

	// db name is in tfvars file
	if _, err := db.Exec("use demodb"); err != nil {
		fmt.Println("couldn't use db....")
	}

	if _, err := db.Exec("create table demotable (id integer, name varchar(20) )"); err != nil {
		fmt.Println(err)
	}

}
