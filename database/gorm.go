package database

import (
	"HMS-16-BE/config"
	"HMS-16-BE/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func InitGorm() {
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

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		AllowGlobalUpdate: true,
	})

	DB = db

	DB.AutoMigrate(
		&model.Admins{},
		&model.Users{},
		&model.Patients{},
		&model.Guardians{},
	)
}
