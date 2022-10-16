package helpers

import (
	"net/mail"
	"net/url"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GetContentType(c *gin.Context) string {
	return c.GetHeader("Content-Type")
}

func GetUserID(c *gin.Context) uint {
	userData := c.MustGet("userData").(jwt.MapClaims)
	return uint(userData["id"].(float64))
}

func ValidateEmail(emailString string) bool {
	_, err := mail.ParseAddress(emailString)
	return err == nil
}

func ValidateURL(urlString string) bool {
	_, err := url.ParseRequestURI(urlString)
	return err == nil
}

func StringToInt(s string) (int, error) {
	data, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return data, nil
}

func StringToUint(s string) (uint, error) {
	data, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}
	return uint(data), nil
}
