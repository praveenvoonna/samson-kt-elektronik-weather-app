package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/praveenvoonna/weather-app/backend/config"
	"github.com/praveenvoonna/weather-app/backend/middleware"
	"github.com/praveenvoonna/weather-app/backend/validations"
	"go.uber.org/zap"
)

type WeatherResponse struct {
	Coord struct {
		Lon float64 `json:"lon"`
		Lat float64 `json:"lat"`
	} `json:"coord"`
	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Base       string `json:"base"`
	Main       Main   `json:"main"`
	Visibility int    `json:"visibility"`
	Wind       struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
		Gust  float64 `json:"gust"`
	} `json:"wind"`
	Rain struct {
		Hourly float64 `json:"1h"`
	} `json:"rain"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	DT       int    `json:"dt"`
	Sys      Sys    `json:"sys"`
	Timezone int    `json:"timezone"`
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Cod      int    `json:"cod"`
}

type Main struct {
	Temp      float64 `json:"temp"`
	FeelsLike float64 `json:"feels_like"`
	TempMin   float64 `json:"temp_min"`
	TempMax   float64 `json:"temp_max"`
	Pressure  int     `json:"pressure"`
	Humidity  int     `json:"humidity"`
	SeaLevel  int     `json:"sea_level"`
	GrndLevel int     `json:"grnd_level"`
}

type Sys struct {
	Type    int    `json:"type"`
	ID      int    `json:"id"`
	Country string `json:"country"`
	Sunrise int    `json:"sunrise"`
	Sunset  int    `json:"sunset"`
}

func GetCurrentWeather(c *gin.Context, db *sql.DB, logger *zap.Logger) {
	authHeader := c.GetHeader("Authorization")
	username, errMsg, err := middleware.AuthenticateJwtToken(authHeader)
	if errMsg != "" || err != nil {
		logger.Error("can not authenticate jwt token", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": errMsg})
	}

	city := c.Query("city")

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

	var weatherData WeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&weatherData); err != nil {
		logger.Error("can not decode weather data from openweathermap api", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to decode weather data"})
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
