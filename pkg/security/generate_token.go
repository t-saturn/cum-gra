package security

import (
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
