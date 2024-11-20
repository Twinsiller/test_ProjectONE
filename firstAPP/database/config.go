package database

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func StringConnectToBase(filePath string) string {
	// Открытие и чтение файла
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Не удалось прочитать файл: %v", err)
	}
	config := struct {
		User     string `json:"user"`
		Password string `json:"password"`
		Dbname   string `json:"dbname"`
		Sslmode  string `json:"sslmode"`
		Host     string `json:"host"`
		Port     string `json:"port"`
	}{}
	// Парсинг JSON
	if err := json.Unmarshal(data, &config); err != nil {
		log.Fatalf("Ошибка парсинга JSON: %v", err)
	}

	// Формирование строки подключения
	result := fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=%s host=%s port=%s",
		config.User, config.Password, config.Dbname, config.Sslmode, config.Host, config.Port,
	)

	return result
}
