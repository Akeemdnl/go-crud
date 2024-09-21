package main

import (
	"database/sql"

	"github.com/Akeemdnl/go-crud/api"
	"github.com/go-sql-driver/mysql"
)

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

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	if err := api.Run(":8080"); err != nil {
		panic(err)
	}
}
