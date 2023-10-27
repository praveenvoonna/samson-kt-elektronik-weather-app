package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/praveenvoonna/weather-app/backend/middleware"
	"github.com/praveenvoonna/weather-app/backend/models"
	"github.com/praveenvoonna/weather-app/backend/utils"
	"go.uber.org/zap"
)

func Register(c *gin.Context, db *sql.DB, logger *zap.Logger) {
	logger.Info("register function called")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		logger.Error("can not bind json of user", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		logger.Error("failed to hash password", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to hash password"})
		return
	}

	var storedUserName string

	err = db.QueryRow("INSERT INTO users(username, password, date_of_birth) VALUES($1, $2, $3) RETURNING username", user.Username, hashedPassword, models.Date(user.DateOfBirth)).Scan(&storedUserName)
	if err != nil {
		logger.Error("can not run inster into users db query", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to insert user"})
		return
	}

	c.Set("username", storedUserName)

	tokenString, err := middleware.GenerateToken(user.Username, logger)
	if err != nil {
		logger.Error("can not generate jwt token for user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "message": "registration successful"})
}

func Login(c *gin.Context, db *sql.DB, logger *zap.Logger) {
	logger.Info("login function called")
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var storedPassword string
	var storedUserName string

	err := db.QueryRow("SELECT username, password FROM users WHERE username=$1", user.Username).Scan(&storedUserName, &storedPassword)
	if err != nil {
		logger.Error("can not fetch username and password data from db", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	if err := utils.ComparePasswords(storedPassword, user.Password); err != nil {
		logger.Error("password validation failed", zap.Error(err))
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username or password"})
		return
	}

	c.Set("username", storedUserName)

	tokenString, err := middleware.GenerateToken(user.Username, logger)
	if err != nil {
		logger.Error("can not generate jwt token for user", zap.Error(err))
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": tokenString, "message": "login successful"})
}
