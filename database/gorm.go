package database

import (
	"HMS-16-BE/config"
	"HMS-16-BE/model"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitGorm() {
	cfg := config.Cfg

	addr := cfg.DB_ADDRESS
	port := cfg.DB_PORT
	user := cfg.DB_USERNAME
	pass := cfg.DB_PASSWORD
	name := cfg.DB_NAME

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, addr, port, name)

	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		AllowGlobalUpdate: true,
	})

	DB = db

	DB.AutoMigrate(
		&model.Admins{},
	)
}
