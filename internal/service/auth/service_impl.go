package auth

import (
	"errors"
	"learning-companion/internal/model"
	"learning-companion/pkg/database"
	"learning-companion/pkg/jwt"
	"strings"
	"time"
)

type serviceImpl struct{}

func NewService() Service {
	return &serviceImpl{}
}

func (s *serviceImpl) Login(username, password string) (*model.User, string, string, error) {
	user := model.User{}

	database.DB.Model(&user).Where("username = ? OR email = ?", username, username).First(&user)

	if user.ID == 0 {
		return nil, "", "", errors.New("user not found")
	}

	if !user.CheckPassword(password) {
		return nil, "", "", errors.New("invalid password")
	}

	acessToken := user.CreateToken()
	if acessToken == "" {
		return nil, "", "", errors.New("failed to create access token")
	}

	parseToken, err := user.ParseToken(acessToken)
	if err != nil {
		return nil, "", "", err
	}

	refreshTokenString := user.CreateRefreshToken()

	refreshToken := model.RefreshToken{
		UserID:    user.Uuid,
		Token:     refreshTokenString,
		CreatedAt: time.Now(),
		ExpiresAt: parseToken.ExpiresAt.Time,
	}

	database.DB.Model(&refreshToken).Create(&refreshToken)

	return &user, acessToken, refreshTokenString, nil
}

func (s *serviceImpl) Logout(accessToken string) error {
	accessToken = strings.TrimPrefix(accessToken, "Bearer ")

	if accessToken == "" {
		return errors.New("authorization header is required")
	}

	claim, err := jwt.GetTokenClaims(accessToken)
	if err != nil {
		return err
	}
	Uuid, ok := claim["user_id"].(string)
	if !ok {
		return errors.New("invalid token claims")
	}

	var refreshToken model.RefreshToken
	database.DB.Model(&refreshToken).Where("user_id= ?", Uuid).Last(&refreshToken)
	if refreshToken.ID == 0 {
		return errors.New("no refresh token found for user")
	}

	if refreshToken.Revoked {
		return errors.New("refresh token already revoked")
	}

	refreshToken.Revoked = true
	if err := database.DB.Save(&refreshToken).Error; err != nil {
		return err
	}

	return nil
}
