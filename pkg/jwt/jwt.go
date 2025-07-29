package jwt

import (
	"learning-companion/internal/config"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID string) (string, error) {
	appConfig := config.Get()
	tokenConfig := appConfig.Token
	secretKey := []byte(appConfig.Server.JWTSecret)
	// Implement JWT token generation logic here
	claims := TokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    tokenConfig.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(tokenConfig.AccessTokenExpiration))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Replace
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateRefreshToken(userID string) (string, error) {
	appConfig := config.Get()
	tokenConfig := appConfig.Token
	secretKey := []byte(appConfig.Server.JWTSecret)
	// Implement JWT refresh token generation logic here
	claims := TokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    tokenConfig.Issuer,
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Minute * time.Duration(tokenConfig.RefreshTokenExpiration))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Replace
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseToken(tokenString string) (*TokenClaims, error) {
	secretKey := []byte(config.Get().Server.JWTSecret)
	// Implement JWT token parsing logic here
	token, err := jwt.ParseWithClaims(tokenString, &TokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		log.Writer().Write([]byte("Error parsing token: " + err.Error()))
		return nil, err
	}

	if claims, ok := token.Claims.(*TokenClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, jwt.ErrSignatureInvalid
}

func ValidateToken(tokenString string) (bool, error) {
	// Implement JWT token validation logic here
	claims, err := ParseToken(tokenString)
	if err != nil {
		return false, err
	}

	// Check if the token is expired
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return false, jwt.ErrTokenExpired
	}

	return true, nil
}

func claimToMap(claims *TokenClaims) map[string]interface{} {
	// Convert claims to a map for easier access
	return map[string]interface{}{
		"user_id": claims.UserID,
		"iss":     claims.Issuer,
		"exp":     claims.ExpiresAt.Unix(),
		"iat":     claims.IssuedAt.Unix(),
	}
}

func GetTokenClaims(tokenString string) (map[string]interface{}, error) {
	claims, err := ParseToken(tokenString)
	if err != nil {
		return nil, err
	}
	return claimToMap(claims), nil
}
