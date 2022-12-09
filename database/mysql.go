package database

import (
	"HMS-16-BE/config"
	"database/sql"
	"fmt"
	"os"
)

func InitMySql() *sql.DB {
	cfg := config.Cfg

	addr := os.Getenv("DB_ADDRESS")
	if addr == "" {
		addr = cfg.DB_ADDRESS
	}
	port := os.Getenv("DB_PORT")
	if port == "" {
		port = cfg.DB_PORT
	}
	user := os.Getenv("DB_USERNAME")
	if user == "" {
		user = cfg.DB_USERNAME
	}
	pass := os.Getenv("DB_PASSWORD")
	if pass == "" {
		pass = cfg.DB_PASSWORD
	}
	name := os.Getenv("DB_NAME")
	if name == "" {
		name = cfg.DB_NAME
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, addr, port, name)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}
