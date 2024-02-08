package utils

import (
	"errors"
	"log"
	"time"

	models "github.com/ashbeelghouri/project1/model"
	"github.com/dgrijalva/jwt-go"
)

func CreateUserToken(claim models.UserToken, secret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"data": claim,
		"exp":  time.Now().Add(time.Minute * 2).Unix(),
	})
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		log.Printf("error while creating user token: %v", err)
		return "", err
	}
	return tokenString, nil
}

func ValidateToken(token string, secret string) (*jwt.Token, error) {
	tokenV, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		log.Printf("error while creating user token: %v", err)
		return nil, errors.New("invalid user token")
	}

	if tokenV.Valid {
		return tokenV, nil
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, errors.New("Not a valid token")
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			// Token is either expired or not active yet
			return nil, errors.New("Token is expired")
		} else {
			log.Printf("error while parsing the token: %v", err)
			return nil, errors.New("Couldn't handle this token, see the logs")
		}
	} else {
		log.Printf("error while parsing the token: %v", err)
		return nil, errors.New("Couldn't handle this token, Validation error, look at logs")
	}
}
