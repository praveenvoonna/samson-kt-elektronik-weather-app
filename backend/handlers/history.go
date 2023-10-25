// handlers/history.go
package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ...

func GetSearchHistory(c *gin.Context, db *sql.DB, userID int) {
	// Fetch search history from the database
	rows, err := db.Query("SELECT id, city_name, search_time FROM search_history WHERE user_id=$1", userID) // provide the user ID here
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch search history"})
		return
	}
	defer rows.Close()

	var searchHistory []gin.H
	for rows.Next() {
		var id int
		var cityName string
		var searchTime time.Time
		err := rows.Scan(&id, &cityName, &searchTime)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan search history"})
			return
		}
		searchHistory = append(searchHistory, gin.H{"id": id, "city_name": cityName, "search_time": searchTime})
	}

	c.JSON(http.StatusOK, searchHistory)
}

func ClearSearchHistory(c *gin.Context, db *sql.DB, userID int) {
	// Implement logic to clear search history from the database
	_, err := db.Exec("DELETE FROM search_history WHERE user_id=$1", userID) // provide the user ID here
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to clear search history"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Search history cleared successfully"})
}

// ...
