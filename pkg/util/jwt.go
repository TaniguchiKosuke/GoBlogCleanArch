package util

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecretKey = os.Getenv("JWT_SECRET_KEY")

func CreateJWTToken(userID string, username string) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))

	token.Claims = jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Minute * 30).Unix(),
	}

	jwtSecretKey := jwtSecretKey

	tokenString, err := token.SignedString([]byte(jwtSecretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func VerifyToken(tokenStr string) (*jwt.Token, error) {

	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecretKey), nil
	})
	if err != nil {
		return nil, err
	}

	return token, nil

}
