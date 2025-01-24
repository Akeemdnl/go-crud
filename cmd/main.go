package main

import (
	"fmt"
	"log"

	"github.com/Akeemdnl/go-crud/api"
	"github.com/Akeemdnl/go-crud/db"
	_ "github.com/lib/pq"
)

func main() {
	database, err := db.InitDB()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to db")

	if err := api.Run(":8080", database); err != nil {
		log.Fatal(err)
	}
}
