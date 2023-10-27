package config

import (
	"log"
	"os"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

type OpenWeatherConfig struct {
	APIKey string
}

type JwtConfig struct {
	JwtKey []byte
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func GetDatabaseConfig(logger *zap.Logger) *DatabaseConfig {
	return &DatabaseConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     getEnvAsInt("DB_PORT", logger),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
	}
}

func GetCorsConfig() gin.HandlerFunc {
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:3000"}
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Authorization", "Content-Type"}
	return cors.New(corsConfig)
}

func GetJwtConfig() *JwtConfig {
	return &JwtConfig{
		JwtKey: []byte(os.Getenv("JWT_SECRET_KEY")),
	}
}

func getEnvAsInt(key string, logger *zap.Logger) int {
	value, err := strconv.Atoi(os.Getenv(key))
	if err != nil {
		log.Fatalf("error converting %s to int", key)
	}
	return value
}

func GetOpenWeatherConfig() *OpenWeatherConfig {
	return &OpenWeatherConfig{
		APIKey: os.Getenv("OPEN_WEATHER_API_KEY"),
	}
}
