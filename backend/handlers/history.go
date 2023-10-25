package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func GetSearchHistory(c *gin.Context, db *sql.DB, userID int) {
	rows, err := db.Query("SELECT id, city_name, search_time FROM search_history WHERE user_id = $1", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve search history"})
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve search history"})
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

func ClearSearchHistory(c *gin.Context, db *sql.DB, userID int) {
	_, err := db.Exec("DELETE FROM search_history WHERE user_id=$1", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear search history"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Search history cleared successfully"})
}
