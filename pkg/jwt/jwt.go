package jwt

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenClaims struct {
	UserID string `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(userID string) (string, error) {
	// Implement JWT token generation logic here
	claims := TokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "learning-companion",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Second * 20)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Replace
	secretKey := []byte("your_secret_key") // Use a secure key in production
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateRefreshToken(userID string) (string, error) {
	// Implement JWT refresh token generation logic here
	claims := TokenClaims{
		UserID: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "learning-companion",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)), // 30 days
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Replace
	secretKey := []byte("your_secret_key") // Use a secure key in production
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ParseToken(tokenString string) (*TokenClaims, error) {
	// Implement JWT token parsing logic here
	secretKey := []byte("your_secret_key") // Use the same key used for signing
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
