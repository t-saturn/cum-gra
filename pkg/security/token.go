package security

import (
	"strconv"
	"time"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"

	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/pkg/logger"
)

// GenerateToken crea un JWT cifrado como JWE con duración personalizada
func GenerateToken(userID string, duration time.Duration) (string, error) {
	secret := []byte(config.GetConfig().Server.JWTSecret)

	now := time.Now()

	claims := jwt.Claims{
		Issuer:   "auth-service",
		Subject:  userID,
		IssuedAt: jwt.NewNumericDate(now),
		Expiry:   jwt.NewNumericDate(now.Add(duration)),
	}

	encrypter, err := jose.NewEncrypter(
		jose.A256GCM,
		jose.Recipient{
			Algorithm: jose.DIRECT,
			Key:       secret,
		},
		(&jose.EncrypterOptions{}).WithContentType("JWT"),
	)
	if err != nil {
		logger.Log.Errorf("Error creando encrypter: %v", err)
		return "", err
	}

	raw, err := jwt.Encrypted(encrypter).Claims(claims).CompactSerialize()
	if err != nil {
		logger.Log.Errorf("Error serializando token JWE: %v", err)
		return "", err
	}

	return raw, nil
}

// GenerateAccessToken crea un JWT cifrado como JWE con la duración
// definida en la variable de entorno JWT_EXP_MINUTES (en minutos).
func GenerateAccessToken(userID string) (string, error) {
	expMinutesStr := config.GetConfig().Server.JWTExpMinutes
	expMinutes, err := strconv.Atoi(expMinutesStr)
	if err != nil {
		logger.Log.Errorf("Valor inválido para JWT_EXP_MINUTES (%s): %v; usando 15 minutos por defecto", expMinutesStr, err)
		expMinutes = 15
	}

	// Llamamos a la función genérica pasándole la duración en minutos
	return GenerateToken(userID, time.Duration(expMinutes)*time.Minute)
}

// GenerateRefreshToken crea un JWT cifrado como JWE para refresh,
// con una duración fija de 7 días (puedes cambiar este valor o hacerlo configurable).
func GenerateRefreshToken(userID string) (string, error) {
	const refreshDays = 7
	return GenerateToken(userID, time.Duration(refreshDays)*24*time.Hour)
}
