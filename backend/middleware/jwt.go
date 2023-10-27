package middleware

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/praveenvoonna/weather-app/backend/config"
	"go.uber.org/zap"
)

var jwtConfig = config.GetJwtConfig()

func GenerateToken(username string, logger *zap.Logger) (string, error) {

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

func GetUsernameFromToken(c *gin.Context) string {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		return ""
	}
	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}
	token := splitToken[1]
	claims := jwt.MapClaims{}
	parsedToken, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtConfig.JwtKey, nil
	})
	if err != nil || !parsedToken.Valid {
		return ""
	}
	return claims["username"].(string)
}
