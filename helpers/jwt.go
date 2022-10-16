package helpers

import (
	"final-project/models"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateToken(id uint, username string, email string) string {
	// create ttl for token 1 day (24 hours)
	ttl := time.Now().Add(time.Hour * 24).Unix()
	claims := jwt.MapClaims{
		"id":       id,
		"username": username,
		"email":    email,
		"exp":      ttl,
	}
	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, _ := parseToken.SignedString(jwtSecret)

	return signedToken
}

func VerifyToken(c *gin.Context) (interface{}, error) {
	headerToken := c.GetHeader("Authorization")
	bearerToken := strings.HasPrefix(headerToken, "Bearer")

	if !bearerToken {
		Response(c, http.StatusUnauthorized, nil, "UNAUTHORIZED", &models.Error{
			Fields:  nil,
			Message: "invalid token",
			Extends: nil,
		})
		c.Abort()
		return nil, nil
	}
	stringToken := strings.Split(headerToken, " ")[1]

	token, _ := jwt.Parse(stringToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, nil
		}
		return jwtSecret, nil
	})

	if _, ok := token.Claims.(jwt.MapClaims); !ok && !token.Valid {
		Response(c, http.StatusUnauthorized, nil, "UNAUTHORIZED", &models.Error{
			Fields:  nil,
			Message: "invalid token",
			Extends: nil,
		})
		c.Abort()
		return nil, nil
	}

	return token.Claims, nil
}
