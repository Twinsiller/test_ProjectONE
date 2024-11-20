package main

import (
	"database/sql"
	"gin-notes-api/database"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var connStr string

var DbGlobal **sql.DB

func main() {
	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Ошибка при открытии файла: %v", err)
	}
	defer file.Close()
	log.SetOutput(file)
	log.Print("Логи запущены")
	connStr := database.StringConnectToBase("database/config_db.json")
	var db, errDb = sql.Open("postgres", connStr)
	if errDb != nil {
		log.Fatalf("Ошибка при открытии файла: %v", errDb)
	}
	defer db.Close()
	DbGlobal = &db
	log.Println("Data base is working!")

	log.Println("Запуск сервера")
	router := gin.Default()

	// Получение всех профилей
	router.GET("/profiles", GetProfiles)

	// Получение профиля по ID
	router.GET("/profiles/:id", GetProfileById)

	// Создание профиля
	router.POST("/profiles", CreateProfile)

	// Обновление существующего профиля
	router.PUT("/profiles/:id", UpdateProfile)

	router.Run(":8080")
}
