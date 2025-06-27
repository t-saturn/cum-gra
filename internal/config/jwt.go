package config

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var jwtSecret []byte

// Inicializa la clave secreta desde .env
func InitJWT() {
	secret := GetEnv("JWT_SECRET", "")
	if secret == "" {
		Logger.Fatal("JWT_SECRET no definido en .env")
	}
	jwtSecret = []byte(secret)
	Logger.Info("JWT inicializado correctamente")
}

// Tipos válidos de token
var validTokenTypes = map[string]time.Duration{
	"access":  3 * time.Hour,
	"refresh": 7 * 24 * time.Hour,
}

// Genera un token JWT
func GenerateJWT(userID string, tokenType string) (tokenStr string, jti string, exp time.Time, err error) {
	duration, ok := validTokenTypes[tokenType]
	if !ok {
		Logger.WithField("token_type", tokenType).Error("Tipo de token inválido")
		return "", "", time.Time{}, fmt.Errorf("tipo de token inválido: %s", tokenType)
	}

	jti = uuid.New().String()
	iat := time.Now()
	exp = iat.Add(duration)

	// Construir claims estándar + personalizados
	claims := jwt.MapClaims{
		"sub":  userID,         // ID del usuario
		"jti":  jti,            // ID único del token
		"exp":  exp.Unix(),     // expiración
		"iat":  iat.Unix(),     // emisión
		"nbf":  iat.Unix(),     // not before
		"iss":  "auth-service", // emisor
		"aud":  "auth-client",  // audiencia
		"type": tokenType,      // tipo personalizado
	}

	// Crear token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token.Header["typ"] = tokenType // tipado personalizado en header

	tokenStr, err = token.SignedString(jwtSecret)
	if err != nil {
		Logger.WithError(err).Error("Error al firmar token JWT")
		return "", "", time.Time{}, err
	}

	Logger.WithFields(map[string]interface{}{
		"user_id":    userID,
		"token_type": tokenType,
		"jti":        jti,
		"exp":        exp,
	}).Info("Token JWT generado exitosamente")

	return tokenStr, jti, exp, nil
}

// Valida un token JWT y verifica su firma
func ValidateJWT(tokenStr string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Verificar método de firma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			Logger.WithField("alg", token.Header["alg"]).Error("Método de firma inválido")
			return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		Logger.WithError(err).Warn("Token JWT inválido")
		return nil, err
	}

	Logger.WithField("token_valid", token.Valid).Debug("Token JWT validado")
	return token, nil
}

// Extrae los claims de un JWT válido
func ExtractClaims(tokenStr string) (jwt.MapClaims, error) {
	token, err := ValidateJWT(tokenStr)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	Logger.Warn("Claims inválidos o token no válido")
	return nil, fmt.Errorf("claims inválidos o token inválido")
}
