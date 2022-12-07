package database

import (
	"HMS-16-BE/config"
	"database/sql"
	"fmt"
)

func InitMySql() *sql.DB {
	cfg := config.Cfg

	addr := cfg.DB_ADDRESS
	port := cfg.DB_PORT
	user := cfg.DB_USERNAME
	pass := cfg.DB_PASSWORD
	name := cfg.DB_NAME

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
