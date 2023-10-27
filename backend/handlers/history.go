package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func SaveSearchHistory(db *sql.DB, username, cityName string) error {
	_, err := db.Exec("INSERT INTO search_history (username, city_name, search_time) VALUES ($1, $2, $3)", username, cityName, time.Now())
	return err
}

func GetSearchHistory(c *gin.Context, db *sql.DB) {
	username := getUsernameFromToken(c)
	rows, err := db.Query("SELECT id, city_name, search_time FROM search_history WHERE username = $1", username)
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

func ClearSearchHistory(c *gin.Context, db *sql.DB) {
	username := getUsernameFromToken(c)
	id := c.Query("id")
	_, err := db.Exec("DELETE FROM search_history WHERE username = $1 AND id = $2", username, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear search history"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Search history cleared successfully"})
}
