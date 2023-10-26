package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/praveenvoonna/weather-app/backend/handlers"

	_ "github.com/lib/pq"
)

var db *sql.DB

func main() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Authorization", "Content-Type"}
	r.Use(cors.New(config))

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

	r.POST("/login", func(c *gin.Context) {
		handlers.Login(c, db)
	})

	r.POST("/register", func(c *gin.Context) {
		handlers.Register(c, db)
	})

	r.GET("/weather", func(c *gin.Context) {
		handlers.GetCurrentWeather(c, db)
	})

	r.GET("/history", func(c *gin.Context) {
		username := c.GetString("username")
		handlers.GetSearchHistory(c, db, username)
	})

	r.DELETE("/history", func(c *gin.Context) {
		username := c.GetString("username")
		handlers.ClearSearchHistory(c, db, username)
	})

	r.Run(":8080")
}
