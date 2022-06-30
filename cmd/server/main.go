package main

import (
	"log"

	"github.com/b-o-e-v/rest-api-store/app/server"
	"github.com/joho/godotenv"
)

// ПРОВЕРКА НА НАЛИЧИЕ ФАЙЛА .env
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal(".env файл не найден")
	}
}

func main() {
	config := server.NewConfig()
	s := server.New(config)

	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
