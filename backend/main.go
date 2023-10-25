package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/praveenvoonna/weather-app/backend/handlers"

	_ "github.com/lib/pq"
)

func main() {
	r := gin.Default()

	host := "localhost"
	port := 5432
	user := "postgres"
	password := "112233"
	dbname := "weather_app_db"

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
	}

	fmt.Println("Successfully connected to the PostgreSQL database!")

	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)
	r.GET("/weather", handlers.GetCurrentWeather)
	r.GET("/history", handlers.GetSearchHistory)
	r.DELETE("/history", handlers.ClearSearchHistory)

	r.Run(":8080")
}
