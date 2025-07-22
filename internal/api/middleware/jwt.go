package middleware

import (
	"fmt"
	"learning-companion/internal/model"
	"learning-companion/internal/response"
	"learning-companion/pkg/database"
	auth "learning-companion/pkg/jwt"
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

		accessToken := c.GetHeader("Authorization")
		accessToken = strings.TrimPrefix(accessToken, "Bearer ")
		if accessToken == "" {
			response.Error(c, "Authorization header is required", http.StatusUnauthorized)
			c.Abort()
			return
		}

		valid, err := auth.ValidateToken(accessToken)
		if err != nil || !valid {
			response.Error(c, "Invalid or expired token", http.StatusUnauthorized)
			c.Abort()
			return
		}

		claim, _ := auth.GetTokenClaims(accessToken)
		Uuid, ok := claim["user_id"].(string)
		if !ok {
			response.Error(c, "Invalid token claims", http.StatusUnauthorized)
			c.Abort()
			return
		}

		user := model.User{}
		database.DB.Model(&user).Where("uuid = ?", Uuid).First(&user)
		if user.ID == 0 {
			response.Error(c, "User not found", http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set("user", user)
		fmt.Println("claim:", claim)
		c.Next()
	}
}
