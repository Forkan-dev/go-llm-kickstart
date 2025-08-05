package middleware

import (
	"fmt"
	"learning-companion/internal/model/auth"
	"learning-companion/internal/model/user"
	"learning-companion/internal/response"
	"learning-companion/pkg/database"
	jwtPkg "learning-companion/pkg/jwt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

// JWTPublicMiddleware authenticates requests using JWT tokens.
func JWTPublicMiddleware(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		log.Default().Println("authHeader:", authHeader)

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
		tokenString = strings.TrimSpace(tokenString)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return []byte(secretKey), nil
		})
		// log.Default().Println("token:", token, "err:", err, err != nil)

		if err != nil {
			response.Error(c, "Invalid token: "+err.Error(), http.StatusUnauthorized)
			c.Abort()
			return
		}
		
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			log.Default().Println("is valid:", token.Valid)
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

		valid, err := jwtPkg.ValidateToken(accessToken)
		if err != nil || !valid {
			response.Error(c, "Invalid or expired token", http.StatusUnauthorized)
			c.Abort()
			return
		}

		claim, _ := jwtPkg.GetTokenClaims(accessToken)
		Uuid, ok := claim["user_id"].(string)
		if !ok {
			response.Error(c, "Invalid token claims", http.StatusUnauthorized)
			c.Abort()
			return
		}

		user := user.User{}
		database.DB.Model(&user).Where("uuid = ?", Uuid).First(&user)
		if user.ID == 0 {
			response.Error(c, "User not found", http.StatusUnauthorized)
			c.Abort()
			return
		}

		refreshToken := auth.RefreshToken{}
		database.DB.Model(&refreshToken).Where("user_id = ?", Uuid).Last(&refreshToken)
		if refreshToken.ID == 0 {
			response.Error(c, "No refresh token found for user", http.StatusUnauthorized)
			c.Abort()
			return
		}

		if refreshToken.Revoked {
			response.Error(c, "Refresh token has been revoked", http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Set("user", user)
		fmt.Println("claim:", claim)
		c.Next()
	}
}
