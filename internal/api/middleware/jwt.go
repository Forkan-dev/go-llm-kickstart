package middleware

import (
	"learning-companion/internal/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// JWTPublicMiddleware authenticates requests using JWT tokens.
func JWTPublicMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Error(c, "Authorization header is required", http.StatusUnauthorized)
			c.Abort()
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Error(c, "Authorization header format must be Bearer {token}", http.StatusUnauthorized)
			c.Abort()
			return
		}

		tokenString := parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secretKey), nil
		})

		if err != nil {
			response.Error(c, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// You can access claims here if needed
		// if claims, ok := token.Claims.(jwt.MapClaims); ok {
		// 	c.Set("user_id", claims["user_id"])
		// }

		c.Next()
	}
}

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// This middleware can be used for authenticated routes
		// You can implement JWT validation logic here similar to JWTPublicMiddleware
		c.Next()
	}
}