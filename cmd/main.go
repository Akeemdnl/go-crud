package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Akeemdnl/go-crud/api"
	_ "github.com/lib/pq"
)

const dbType = "postgres"

func main() {
	dsn := "user=postgres password=mysecretpassword dbname=gocrud host=localhost port=5432 sslmode=disable"
	db, err := sql.Open(dbType, dsn)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connected to db:", dbType)

	if err := api.Run(":8080", db); err != nil {
		log.Fatal(err)
	}
}
