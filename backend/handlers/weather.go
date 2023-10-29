package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/praveenvoonna/weather-app/backend/config"
	"github.com/praveenvoonna/weather-app/backend/middleware"
	"github.com/praveenvoonna/weather-app/backend/models"
	"github.com/praveenvoonna/weather-app/backend/validations"
	"go.uber.org/zap"
)

func GetCurrentWeather(c *gin.Context, db *sql.DB, logger *zap.Logger) {
	authHeader := c.GetHeader("Authorization")
	username, errMsg, err := middleware.AuthenticateJwtToken(authHeader)
	if errMsg != "" || err != nil {
		logger.Error("can not authenticate jwt token", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": errMsg})
	}

	city := c.Query("city")
	logger.Info("city " + city)

	if !validations.ValidateWeatherCheckInput(c, city, logger) {
		return
	}

	apiKey := config.GetOpenWeatherConfig().APIKey

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		logger.Error("can not fetch weather data from openweathermap api", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch weather data"})
		return
	}
	defer resp.Body.Close()

	var weatherData models.WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		logger.Error("can not decode weather data from openweathermap api", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode weather data"})
		return
	}

	if weatherData.StausCode == http.StatusNotFound {
		logger.Warn("can not find weather data for given city from openweathermap api")
		c.JSON(http.StatusNotFound, gin.H{"error": weatherData.Message})
		return
	} else if weatherData.StausCode != http.StatusOK {
		c.JSON(http.StatusInternalServerError, gin.H{"error": weatherData.Message})
		return
	}

	err = SaveSearchHistory(db, username, city)
	if err != nil {
		logger.Error("can not save weather history data", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save search history"})
		return
	}

	c.JSON(http.StatusOK, weatherData)
}
