package server

import (
	"fmt"

	"github.com/b-o-e-v/rest-api-store/app/store"
	"github.com/b-o-e-v/rest-api-store/healpers"
)

// СТРУКТУРА КОНФИГА
type Config struct {
	Port     string
	LogLevel string
	Store    *store.Config
}

// СОЗДАЕМ НОВЫЙ КОНФИГ
func NewConfig() *Config {
	return &Config{
		Port:     fmt.Sprintf(":%s", healpers.GetEnv("PORT", "3000")),
		LogLevel: healpers.GetEnv("LOG_LEVEL", "debug"),
		Store:    store.NewConfig(),
	}
}
