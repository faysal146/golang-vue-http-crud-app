package helpers

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(os.Getenv("JWT_SECRETKEY"))

type JWTClaim struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateAuthToken(id, email string) (jwt_token string, refresh_token string) {
	tokenClaims := &JWTClaim{
		Id:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(1 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	refreshTokenClaim := &JWTClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}
	tokenString, err := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims).SignedString(jwtKey)
	if err != nil {
		log.Fatal("could not generate token", err)
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshTokenClaim).SignedString(jwtKey)
	if err != nil {
		log.Fatal("could not generate token", err)
	}
	return tokenString, refreshToken
}
