package config

import (
	"log"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecret []byte

func InitJWT() {
	secret := GetEnv("JWT_SECRET", "")
	if secret == "" {
		log.Fatal("JWT_SECRET no definido en .env")
	}
	jwtSecret = []byte(secret)
}

func GenerateJWT(userID string) (string, string, time.Time, error) {
	expMinStr := GetEnv("JWT_EXP_MINUTES", "15")
	expMin, err := strconv.Atoi(expMinStr)
	if err != nil {
		expMin = 15
	}

	expirationTime := time.Now().Add(time.Duration(expMin) * time.Minute)
	jti := uuid.New().String()

	claims := jwt.MapClaims{
		"sub": userID,
		//		"user_id": userID,
		"jti": jti,
		"exp": expirationTime.Unix(),
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", "", time.Time{}, err
	}

	return signedToken, jti, expirationTime, nil
}
