package middleware

import (
	helpers "final-project/helpers"

	"github.com/gin-gonic/gin"
)

func Authenticated() gin.HandlerFunc {
	return func(c *gin.Context) {
		verifyToken, err := helpers.VerifyToken(c)
		_ = verifyToken

		if err != nil {
			helpers.Response(c, 401, nil, "UNAUTHORIZED", nil)
			c.Abort()
			return
		}
		c.Set("userData", verifyToken)
		c.Next()
	}
}
