package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type ServerConfig struct {
	Host string
	Port string
}

type RedisConfig struct {
	URL string
}

type Config struct {
	Server ServerConfig
	Redis  RedisConfig
}

var AppConfig *Config

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Printf("Error loading .env file: %v", err)
	}

	return &Config{
		Server: ServerConfig{
			Host: os.Getenv("Server__Host"),
			Port: os.Getenv("Server__Port"),
		},
		Redis: RedisConfig{
			URL: os.Getenv("Redis__URL"),
		},
	}
}
