package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type PGConfig struct {
	Username string
	Password string
	Host     string
	DBName   string
	Port     string
	Mode     string
}

type HTTPServer struct {
	AppPort        string
	AppBindAddress string
	*PGConfig
}

func getEnv(envValue string) string {
	val := os.Getenv(envValue)
	if envValue != "APP_BIND_ADDRESS" && val == "" {
		log.Fatalf("Environment variable '%s' not set or empty.", envValue)
	}
	return os.Getenv(envValue)
}

func MustLoad() *HTTPServer {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env file")
	}

	return &HTTPServer{
		AppPort:        getEnv("APP_PORT"),
		AppBindAddress: getEnv("APP_BIND_ADDRESS"),
		PGConfig: &PGConfig{
			Username: getEnv("USERNAME"),
			Password: getEnv("PASSWORD"),
			Host:     getEnv("HOST"),
			DBName:   getEnv("NAME"),
			Port:     getEnv("PORT"),
			Mode:     getEnv("MODE"),
		},
	}
}
