package middleware

import (
	"egardev-gin-socmed/errorHandler"
	"egardev-gin-socmed/helper"
	"github.com/gin-gonic/gin"
)

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			errorHandler.HandleError(c, &errorHandler.UnauthorizedError{Message: "Unauthorized"})
			c.Abort()
			return
		}

		userId, err := helper.ValidateToken(tokenString)
		if err != nil {
			errorHandler.HandleError(c, &errorHandler.UnauthorizedError{Message: err.Error()})
			c.Abort()
			return
		}

		c.Set("userID", *userId)
		c.Next()
	}
}
