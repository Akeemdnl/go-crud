package db

import (
	"database/sql"
	"fmt"

	"github.com/Akeemdnl/go-crud/configs"
)

const DbType = "postgres"

func InitDB() (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s",
		configs.Env.User,
		configs.Env.Password,
		configs.Env.DBName,
		configs.Env.Host,
		configs.Env.Port,
		configs.Env.Sslmode,
	)

	db, err := sql.Open(DbType, dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
