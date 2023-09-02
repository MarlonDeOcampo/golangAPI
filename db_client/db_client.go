package db_client

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var DBClient *sql.DB

func InitializeDBConnection() {
	connStr := "user=postgres dbname=postgres password=secret host=postgres sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	DBClient = db

	fmt.Printf("\nSuccessfully connected to database!\n")
}
