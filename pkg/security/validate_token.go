package security

import (
	"errors"
	"time"

	"gopkg.in/square/go-jose.v2/jwt"

	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/pkg/logger"
)

// TokenValidationResult contiene el resultado y mensaje opcional
type TokenValidationResult struct {
	Code    int
	Claims  *jwt.Claims
	Message string
}

// ValidateToken descifra y valida un token JWT cifrado como JWE
func ValidateToken(tokenString string) TokenValidationResult {
	secret := []byte(config.GetConfig().JWT_SECRET)

	// Parsear el token JWE
	tok, err := jwt.ParseEncrypted(tokenString)
	if err != nil {
		logger.Log.Warnf("Token inválido (no se pudo parsear): %v", err)
		return TokenValidationResult{Code: 1, Message: "Token inválido"}
	}

	var claims jwt.Claims
	if err := tok.Claims(secret, &claims); err != nil {
		logger.Log.Warnf("Error al obtener claims del token (firma inválida o clave incorrecta): %v", err)
		return TokenValidationResult{Code: 3, Message: "Firma inválida o clave incorrecta"}
	}

	// Validar expiración y otros claims
	err = claims.Validate(jwt.Expected{
		Time: time.Now(),
	})
	if err != nil {
		if errors.Is(err, jwt.ErrExpired) {
			logger.Log.Infof("Token expirado")
			return TokenValidationResult{Code: 2, Message: "Token expirado"}
		}
		logger.Log.Warnf("Token inválido por claims: %v", err)
		return TokenValidationResult{Code: 1, Message: "Claims inválidos"}
	}

	// Token válido
	return TokenValidationResult{Code: 0, Claims: &claims}
}
