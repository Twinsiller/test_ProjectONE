package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB — глобальная переменная для хранения подключения к базе данных
var DbPostgres *gorm.DB

func initDB() {
	// Config — структура для хранения конфигурации подключения
	cfg := struct {
		User     string
		Password string
		Host     string
		Port     string
		DBName   string
		SSLMode  string
	}{
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}
	fmt.Println("sadas ", cfg.User)
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, os.Getenv("DB_USER"), cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode,
	)
	fmt.Println(dsn)
	var err error
	DbPostgres, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	fmt.Println("DB is working")
}

func main() {
	initDB()
}
