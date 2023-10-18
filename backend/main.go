// main.go
package main

import (
	"github.com/praveenvoonna/samson-kt-elektronik-weather-app/backend/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", handlers.Login)
	r.POST("/register", handlers.Register)
	r.GET("/weather", handlers.GetCurrentWeather)
	r.GET("/history", handlers.GetSearchHistory)
	r.DELETE("/history", handlers.ClearSearchHistory)
	r.Run(":8080")
}
