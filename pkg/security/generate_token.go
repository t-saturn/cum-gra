package security

import (
	"time"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"

	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/pkg/logger"
)

// GenerateToken crea un JWT y lo cifra como JWE con go-jose.v2
func GenerateToken(userID string) (string, error) {
	// Obtener secreto desde config
	secret := []byte(config.GetConfig().Server.JWTSecret)

	// Crear el token JWT con claims personalizados
	claims := jwt.Claims{
		Issuer:   "auth-service",
		Subject:  userID,
		IssuedAt: jwt.NewNumericDate(time.Now()),
		Expiry:   jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
	}

	// Algoritmo de encriptación: direct + AES256 GCM
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

	// Encriptar el JWT con claims
	raw, err := jwt.Encrypted(encrypter).Claims(claims).CompactSerialize()
	if err != nil {
		logger.Log.Errorf("Error serializando token JWE: %v", err)
		return "", err
	}

	return raw, nil
}

