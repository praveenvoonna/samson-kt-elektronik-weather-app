package middleware

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/praveenvoonna/weather-app/backend/config"
)

var jwtConfig = config.GetJwtConfig()

func GenerateToken(username string) (string, error) {

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 2).Unix() // Token expires in 2 hours

	tokenString, err := token.SignedString(jwtConfig.JwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AuthenticateJwtToken(authHeader string) (string, string, error) {
	if authHeader == "" {
		// logger.Error("no auth hearder found")
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "authorization header is required"})
		return "", "authorization header is required", errors.New("no auth hearder found")
	}

	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		// logger.Error("invalid token format")
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
		return "", "invalid token format", errors.New("invalid token format")
	}

	tokenString := splitToken[1]

	token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	if err != nil {
		// logger.Error("invalid token", zap.Error(err))
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token " + err.Error()})
		return "", "invalid token", err
	}
	var username string
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		username = fmt.Sprint(claims["username"])
		// logger.Info("name " + username)
	} else {
		// logger.Error("invalid token no user name found")
		// c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return "", "invalid token", errors.New("invalid token no user name found")
	}
	return username, "", nil
}
