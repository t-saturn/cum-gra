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
func GenerateToken(userID, audience string, duration time.Duration) (string, error) {
	cfg := config.GetConfig()

	privKey, err := parseRSAPrivateKeyFromPEM(cfg.Server.JWTPrivateKeyPath)
	if err != nil {
		logger.Log.Errorf("Error leyendo clave privada: %v", err)
		return "", err
	}

	// Header: typ=JWT + kid=ENV (go-jose v2)
	opts := (&jose.SignerOptions{}).
		WithType("JWT").
		WithHeader("kid", cfg.Server.JWTKid)

	signer, err := jose.NewSigner(
		jose.SigningKey{Algorithm: jose.RS256, Key: privKey},
		opts,
	)
	if err != nil {
		logger.Log.Errorf("Error creando signer: %v", err)
		return "", err
	}

	now := time.Now()
	claims := jwt.Claims{
		Issuer:   cfg.Server.JWTIss, // <-- ISS del env
		Subject:  userID,
		Audience: jwt.Audience{audience}, // <-- AUD por app
		IssuedAt: jwt.NewNumericDate(now),
		Expiry:   jwt.NewNumericDate(now.Add(duration)),
		// (opcional) NotBefore: jwt.NewNumericDate(now),
		// (opcional extras propios)
	}

	// (opcional) claims personalizados
	custom := map[string]any{
		"sid": audience,
		"jti": audience,
	}

	raw, err := jwt.Signed(signer).Claims(claims).Claims(custom).CompactSerialize()
	if err != nil {
		logger.Log.Errorf("Error firmando token: %v", err)
		return "", err
	}
	return raw, nil
}

// GenerateAccessToken usa duración configurable en minutos
func GenerateAccessToken(userID, audience string) (string, error) {
	expMinutesStr := config.GetConfig().Server.JWTExpMinutes
	expMinutes, err := strconv.Atoi(expMinutesStr)
	if err != nil {
		logger.Log.Errorf("Valor inválido para JWT_EXP_MINUTES (%s): %v; usando 15", expMinutesStr, err)
		expMinutes = 15
	}
	return GenerateToken(userID, audience, time.Duration(expMinutes)*time.Minute)
}

// GenerateRefreshToken con duración de 7 días
func GenerateRefreshToken(userID, audience string) (string, error) {
	return GenerateToken(userID, audience, 7*24*time.Hour)
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
	cfg := config.GetConfig()

	pubKey, err := parseRSAPublicKeyFromPEM(cfg.Server.JWTPublicKeyPath)
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

	if err := claims.Validate(jwt.Expected{
		Time:   time.Now(),
		Issuer: cfg.Server.JWTIss, // <-- usar el mismo ISS con el que firmas
		// opcional: Audience: jwt.Audience{expectedAud},
	}); err != nil {
		if claims.Expiry != nil && time.Now().After(claims.Expiry.Time()) {
			return jwt.Claims{}, ErrTokenExpired
		}
		return jwt.Claims{}, ErrTokenInvalid
	}
	return claims, nil
}
