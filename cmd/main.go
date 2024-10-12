package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Akeemdnl/go-crud/api"
	"github.com/go-sql-driver/mysql"
)

const dbType = "mysql"

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "mypassword",
		DBName:               "gocrud",
		Addr:                 "127.0.0.1:3306",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	db, err := sql.Open(dbType, cfg.FormatDSN())
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
