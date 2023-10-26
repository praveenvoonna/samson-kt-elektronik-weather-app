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

	// router.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"https://foo.com"},
	// 	AllowMethods:     []string{"PUT", "PATCH"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	AllowOriginFunc: func(origin string) bool {
	// 	  return origin == "https://github.com"
	// 	},
	// 	MaxAge: 12 * time.Hour,
	//   }))

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"} // Replace with your frontend's URL
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
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

	r.GET("/weather", handlers.GetCurrentWeather)

	r.GET("/history", func(c *gin.Context) {
		userID := c.GetInt("userID")
		handlers.GetSearchHistory(c, db, userID)
	})

	r.DELETE("/history", func(c *gin.Context) {
		userID := c.GetInt("userID")
		handlers.ClearSearchHistory(c, db, userID)
	})

	r.Run(":8080")
}
