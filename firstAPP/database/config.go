package database

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	user     string `json:"user"`
	password string `json:"password"`
	dbname   string `json:"dbname"`
	sslmode  string `json:"sslmode"`
	host     string `json:"host"`
	port     string `json:"port"`
}

func LoadConfig(filePath string) (Config, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return _, fmt.Errorf("не удалось открыть файл конфигурации: %w", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&Config); err != nil {
		return _, fmt.Errorf("ошибка декодирования JSON: %w", err)
	}
	return _, nil
}
