package helpers

import (
	"errors"
	"log"
	"os"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte(os.Getenv("JWT_SECRETKEY"))

type JWTClaim struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

func IsTokenValid(token string) bool {
	err := validator.New().Var(token, "required,jwt")
	return err == nil
}

func GenerateAuthToken(id, email string) (jwt_token string, refresh_token string) {
	tokenClaims := &JWTClaim{
		Id:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(1 * time.Hour).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
		},
	}
	refreshTokenClaim := &JWTClaim{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Local().Unix(),
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

func VerifyToken(token string) (payload *JWTClaim, err error) {
	tkn, err := jwt.ParseWithClaims(token, &JWTClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(jwtKey), nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return &JWTClaim{}, errors.New("invalid signature")
		}
		return &JWTClaim{}, errors.New("could not parse token")
	}
	if !tkn.Valid {
		return &JWTClaim{}, errors.New("invalid token")
	}

	claims, ok := tkn.Claims.(*JWTClaim)
	if !ok {
		return &JWTClaim{}, errors.New("couldn't parse claims")
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return &JWTClaim{}, errors.New("token expired")
	}
	return claims, nil
}
