package main

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB — глобальная переменная для хранения подключения к базе данных
var DbPostgres *gorm.DB

func initDB() {
	// Загружаем переменные окружения из файла .env
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Error loading .env file")
		return
	}
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
	//res, ok := os.LookupEnv("DB_PASSWORD")
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		cfg.Host, cfg.User, cfg.Password, cfg.DBName, cfg.Port, cfg.SSLMode,
	)
	fmt.Println(dsn)
	var err error
	DbPostgres, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return
	}
	fmt.Println("DB is working")
}

type Profile struct {
	Id           int       `json:"id",gorm:"primaryKey;"`
	Nickname     string    `json:"nickname",gorm:""`
	HashPassword string    `json:"hashpassword",gorm:""`
	Status       bool      `json:"status",gorm:""`
	AccessLevel  int       `json:"accesslevel",gorm:""`
	Firstname    string    `json:"firstname",gorm:""`
	Lastname     string    `json:"lastname",gorm:""`
	CreatedAt    time.Time `json:"createdat",gorm:""`
}

func createTables() {
	DbPostgres.AutoMigrate(&Profile{})
}

func main() {
	initDB()
	createTables()
}
