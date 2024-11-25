package main

import (
	"log"
	"os"

	"test_ProjectONE/internal/config"

	"github.com/gin-gonic/gin"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.LoadConfig(".env")
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

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

	//router.Run(":8080")
	// Запуск сервера
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Printf("Сервер запущен на порту %s", port)
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Ошибка запуска сервера: %v", err)
	}
}
