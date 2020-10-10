package router

import (
	"dopi-user/model"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type DopiClaims struct {
	Username string `json:"username"`
	Roles    []string
	jwt.StandardClaims
}

func CreateJwt(user *model.User) (string, *jwt.Token, *DopiClaims, error) {
	var err error

	secret := os.Getenv("JWT_SECRET")
	validityMinutes, _ := strconv.Atoi(os.Getenv("JWT_TOKEN_VALIDITY"))

	var claims = DopiClaims{
		user.Username,
		user.Roles,
		jwt.StandardClaims{
			Issuer:    "dopi",
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(validityMinutes)).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	at := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := at.SignedString([]byte(secret))
	if err != nil {
		return "", nil, nil, err
	}
	return token, at, &claims, nil
}

func VerifyToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.ParseWithClaims(tokenString, &DopiClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetClaims(token *jwt.Token) (*DopiClaims, error) {
	if claims, ok := token.Claims.(*DopiClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("Token invalid")

}
