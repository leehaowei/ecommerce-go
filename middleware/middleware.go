package middleware

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leehaowei/ecommerce-go/token"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		ClientToken := c.Request.Header.Get("token")
		fmt.Printf("ClientToken: %s\n", ClientToken)

		if ClientToken == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No Authorization Header Provided"})
			c.Abort()
			return
		}
		claims, err := token.ValidateToken(ClientToken)
		if err != nil {
			log.Fatal(err.Error())
			c.JSON(http.StatusInternalServerError, gin.H{"error": err})
			c.Abort()
			return
		}

		userEmail := claims["email"].(string)
		userID := claims["id"].(string)
		c.Set("email", userEmail)
		c.Set("uid", userID)
		c.Next()
	}
}
