package store

import "github.com/b-o-e-v/rest-api-store/app/healpers"

// СТРУКТУРА КОНФИГА
type Config struct {
	DbHost     string
	DbPort     string
	DbName     string
	DbUser     string
	DbPassword string
	SslMode    string
}

// СОЗДАЕМ НОВЫЙ КОНФИГ
func NewConfig() *Config {
	return &Config{
		DbHost:     healpers.GetEnv("DB_HOST", "localost"),
		DbPort:     healpers.GetEnv("DB_PORT", "5432"),
		DbName:     healpers.GetEnv("DB_NAME", "store"),
		DbUser:     healpers.GetEnv("DB_USER", "root"),
		DbPassword: healpers.GetEnv("DB_PASSWORD", "159753"),
		SslMode:    healpers.GetEnv("SSL_MODE", "disable"),
	}
}
