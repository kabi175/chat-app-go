package repository

import (
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type Jwt struct {
	key string
}

type customClaims struct {
	UserId string `json:"UserId"`
	jwt.StandardClaims
}

func NewJwt(key string) *Jwt {
	return &Jwt{
		key: key,
	}
}

func (j *Jwt) Key() []byte {
	return []byte(j.key)
}

func (j *Jwt) NewToken(userId string) (AccessToken string, err error) {
	claims := customClaims{
		UserId: userId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(500 * time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	AccessToken, err = token.SignedString([]byte(j.key))
	if err != nil {
		return "", err
	}
	return AccessToken, nil
}

func (j *Jwt) Verify(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&customClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.key), nil
		},
	)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(*customClaims)
	if !ok {
		return "", errors.New("couldn't parse claims")
	}
	return claims.UserId, nil
}
