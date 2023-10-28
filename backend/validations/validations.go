// package validations

package validations

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/praveenvoonna/weather-app/backend/models"
	"go.uber.org/zap"
)

func ValidateUserRegistrationInput(c *gin.Context, user *models.User, logger *zap.Logger) bool {
	if user.Username == "" {
		logger.Error("username input is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
		return false
	}

	if user.Password == "" {
		logger.Error("password input is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is required"})
		return false
	}

	if user.DateOfBirth == "" {
		logger.Error("date of birth input is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "date of birth is required"})
		return false
	}

	dateFormat := "2006-01-02"
	parsedDate, err := time.Parse(dateFormat, user.DateOfBirth)
	if err != nil {
		logger.Error("invalid date of birth format", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date of birth format. Use YYYY-MM-DD"})
		return false
	}

	today := time.Now()
	if parsedDate.After(today) {
		logger.Error("date of birth is in the future")
		c.JSON(http.StatusBadRequest, gin.H{"error": "date of birth cannot be in the future"})
		return false
	}

	return true
}

func ValidateUserLoginInput(c *gin.Context, user *models.User, logger *zap.Logger) bool {
	if user.Username == "" {
		logger.Error("username input is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "username is required"})
		return false
	}

	if user.Password == "" {
		logger.Error("password input is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "password is required"})
		return false
	}

	return true
}

func ValidateWeatherCheckInput(c *gin.Context, city string, logger *zap.Logger) bool {
	if city == "" {
		logger.Error("city input is empty")
		c.JSON(http.StatusBadRequest, gin.H{"error": "city is required"})
		return false
	}

	return true
}
