package config

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func connect() {
	// https://github.com/jackc/pgx
	dsn := "host=localhost user=postgres password=suman dbname=gorm port=5433 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
}
