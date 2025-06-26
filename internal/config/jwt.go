package config

import (
	"log"
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

// tokenType: "access" o "refresh"
func GenerateJWT(userID string, tokenType string) (string, string, time.Time, error) {
	jti := uuid.New().String()
	iat := time.Now()

	// Definir duración según tipo
	var exp time.Time
	switch tokenType {
	case "refresh":
		exp = iat.Add(7 * 24 * time.Hour) // 7 días
	default: // "access"
		exp = iat.Add(3 * time.Hour) // 3 horas
	}

	// Construir claims estándar
	claims := jwt.MapClaims{
		"sub": userID,     // subject (ID del usuario)
		"jti": jti,        // ID único del token
		"exp": exp.Unix(), // expiración
		"iat": iat.Unix(), // emisión
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", "", time.Time{}, err
	}

	return signedToken, jti, exp, nil
}
