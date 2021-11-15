package service

import (
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/kabi175/chat-app-go/messager/domain"
	"github.com/kabi175/chat-app-go/messager/domain/apperrors"
)

type JwtTokenService struct{}

func NewJwtTokenService() domain.TokenService {
	return &JwtTokenService{}
}

func (JwtTokenService) GenerateToken(user *domain.User) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["userID"] = user.ID
	atClaims["exp"] = time.Now().Add(time.Minute * 15).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(os.Getenv("ACCESS_SECRET")))
	if err != nil {
		return "", err
	}
	return token, nil
}

func (JwtTokenService) DecodeToken(tokenString string) (*domain.User, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userID, ok := claims["userID"].(float64)
		if !ok {
			return nil, apperrors.NewInternalServerError("cannot cast claims['userID'] to float64")
		}
		if err != nil {
			return nil, apperrors.NewInternalServerError(err.Error())
		}
		user := &domain.User{
			ID: uint(userID),
		}
		return user, nil
	}
	return nil, apperrors.NewUnauthorizedError("invalid token")
}
