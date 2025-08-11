package security

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"os"
	"strconv"
	"time"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"

	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/pkg/logger"
)

var (
	ErrTokenExpired = errors.New("token_expired")
	ErrTokenInvalid = errors.New("token_invalid")
)

// parseRSAPrivateKeyFromPEM lee y parsea una clave privada RSA
func parseRSAPrivateKeyFromPEM(path string) (*rsa.PrivateKey, error) {
	pemBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("PEM inválido (private key)")
	}
	if key, err := x509.ParsePKCS1PrivateKey(block.Bytes); err == nil {
		return key, nil
	}
	if keyIfc, err := x509.ParsePKCS8PrivateKey(block.Bytes); err == nil {
		if key, ok := keyIfc.(*rsa.PrivateKey); ok {
			return key, nil
		}
	}
	return nil, errors.New("clave privada no es RSA")
}

// GenerateToken crea un JWT firmado con RS256
func GenerateToken(userID string, duration time.Duration) (string, error) {
	privKey, err := parseRSAPrivateKeyFromPEM(config.GetConfig().Server.JWTPrivateKeyPath)
	if err != nil {
		logger.Log.Errorf("Error leyendo clave privada: %v", err)
		return "", err
	}

	signer, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.RS256, Key: privKey},
		(&jose.SignerOptions{}).WithType("JWT"),
	)
	if err != nil {
		logger.Log.Errorf("Error creando signer: %v", err)
		return "", err
	}

	now := time.Now()
	claims := jwt.Claims{
		Issuer:   "auth-service",
		Subject:  userID,
		IssuedAt: jwt.NewNumericDate(now),
		Expiry:   jwt.NewNumericDate(now.Add(duration)),
	}

	raw, err := jwt.Signed(signer).Claims(claims).CompactSerialize()
	if err != nil {
		logger.Log.Errorf("Error firmando token: %v", err)
		return "", err
	}

	return raw, nil
}

// GenerateAccessToken usa duración configurable en minutos
func GenerateAccessToken(userID string) (string, error) {
	expMinutesStr := config.GetConfig().Server.JWTExpMinutes
	expMinutes, err := strconv.Atoi(expMinutesStr)
	if err != nil {
		logger.Log.Errorf("Valor inválido para JWT_EXP_MINUTES (%s): %v; usando 15 minutos por defecto", expMinutesStr, err)
		expMinutes = 15
	}
	return GenerateToken(userID, time.Duration(expMinutes)*time.Minute)
}

// GenerateRefreshToken con duración de 7 días
func GenerateRefreshToken(userID string) (string, error) {
	const refreshDays = 7
	return GenerateToken(userID, time.Duration(refreshDays)*24*time.Hour)
}

func parseRSAPublicKeyFromPEM(path string) (*rsa.PublicKey, error) {
	pemBytes, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("PEM inválido (public key)")
	}
	pubIfc, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	pub, ok := pubIfc.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("clave pública no es RSA")
	}
	return pub, nil
}

func VerifyTokenRS256(tokenStr string) (jwt.Claims, error) {
	pubKey, err := parseRSAPublicKeyFromPEM(config.GetConfig().Server.JWTPublicKeyPath)
	if err != nil {
		return jwt.Claims{}, ErrTokenInvalid
	}

	tok, err := jwt.ParseSigned(tokenStr)
	if err != nil {
		return jwt.Claims{}, ErrTokenInvalid
	}

	var claims jwt.Claims
	if err := tok.Claims(pubKey, &claims); err != nil {
		return jwt.Claims{}, ErrTokenInvalid
	}

	// Validaciones estándar: exp/iat/nbf/iss (ajusta Expected según tu caso)
	if err := claims.Validate(jwt.Expected{
		Time:   time.Now(),
		Issuer: "auth-service",
	}); err != nil {
		// go-jose no trae un tipo “ExpiredError” explícito; si quieres distinguir,
		// puedes parsear el mensaje o validar manualmente claims.Expiry.
		if claims.Expiry != nil && time.Now().After(claims.Expiry.Time()) {
			return jwt.Claims{}, ErrTokenExpired
		}
		return jwt.Claims{}, ErrTokenInvalid
	}

	return claims, nil
}
