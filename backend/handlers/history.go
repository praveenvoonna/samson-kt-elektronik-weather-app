package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/praveenvoonna/weather-app/backend/middleware"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SaveSearchHistory(db *sql.DB, username, cityName string) error {
	_, err := db.Exec("INSERT INTO search_history (username, city_name, search_time) VALUES ($1, $2, $3)", username, cityName, time.Now())
	return err
}

func GetSearchHistory(c *gin.Context, db *sql.DB, logger *zap.Logger) {
	authHeader := c.GetHeader("Authorization")
	username, errMsg, err := middleware.AuthenticateJwtToken(authHeader)
	if errMsg != "" || err != nil {
		logger.Error("can not authenticate jwt token", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": errMsg})
	}

	logger.Info("get search history called for " + username)
	rows, err := db.Query("SELECT id, city_name, search_time FROM search_history WHERE username = $1", username)
	if err != nil {
		logger.Error("can not fetch weather history data from db", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve search history"})
		return
	}
	defer rows.Close()

	var searchHistory []map[string]interface{}

	for rows.Next() {
		var id int
		var cityName string
		var searchTime time.Time
		err := rows.Scan(&id, &cityName, &searchTime)
		if err != nil {
			logger.Error("can not scan rows of weather data from db", zap.Error(err))
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve search history"})
			return
		}
		searchHistory = append(searchHistory, map[string]interface{}{
			"id":          id,
			"city_name":   cityName,
			"search_time": searchTime,
		})
	}

	c.JSON(http.StatusOK, searchHistory)
}

func ClearSearchHistory(c *gin.Context, db *sql.DB, logger *zap.Logger) {
	authHeader := c.GetHeader("Authorization")
	username, errMsg, err := middleware.AuthenticateJwtToken(authHeader)
	if errMsg != "" || err != nil {
		logger.Error("can not authenticate jwt token", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": errMsg})
	}

	id := c.Query("id")
	logger.Info("clear search history called for " + username + " id " + id)
	result, err := db.Exec("DELETE FROM search_history WHERE username = $1 AND id = $2", username, id)
	if err != nil {
		logger.Error("can not delete weather history data from db", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to clear search history"})
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		logger.Error("error while checking affected rows", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "error while checking affected rows"})
		return
	}

	if rowsAffected > 0 {
		c.JSON(http.StatusOK, gin.H{"message": "search history cleared successfully"})
	} else {
		logger.Error("can not find history with given id", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no rows were affected"})
	}

}
