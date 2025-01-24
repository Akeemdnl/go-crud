package configs

import "os"

var Env = initEnv()

type Configs struct {
	User     string
	Password string
	DBName   string
	Host     string
	Port     string
	Sslmode  string
}

func initEnv() *Configs {
	return &Configs{
		User:     getEnv("DB_USER", "postgres"),
		Password: getEnv("DB_PASSWORD", "mysecretpassword"),
		DBName:   getEnv("DB_NAME", "gocrud"),
		Host:     getEnv("DB_HOST", "go_crud_db"),
		Port:     getEnv("DB_PORT", "5432"),
		Sslmode:  getEnv("DB_SSLMODE", "disable"),
	}
}

func getEnv(key string, fallback string) string {
	if key, ok := os.LookupEnv(key); ok {
		return key
	}
	return fallback
}
