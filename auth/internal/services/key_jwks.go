package services

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"math/big"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/dto"
)

type JWKSService struct {
	kid     string
	alg     string
	pubPath string
	maxAge  int // segundos

	mu        sync.RWMutex
	cached    *dto.JWKSResponseDTO
	expiresAt time.Time
}

func NewJWKSService() *JWKSService {
	cfg := config.GetConfig()
	maxAge := 60
	if v, err := strconv.Atoi(cfg.Server.JWKSMaxAge); err == nil && v > 0 {
		maxAge = v
	}
	alg := cfg.Server.JWTAlg
	if alg == "" {
		alg = "RS256"
	}
	return &JWKSService{
		kid:     cfg.Server.JWTKid,
		alg:     alg,
		pubPath: cfg.Server.JWTPublicKeyPath,
		maxAge:  maxAge,
	}
}

func (s *JWKSService) parseRSAPublicKeyFromPEMBytes(pemBytes []byte) (*rsa.PublicKey, error) {
	block, _ := pem.Decode(pemBytes)
	if block == nil {
		return nil, errors.New("PEM inválido (public key)")
	}
	// Soporta PUBLIC KEY (PKIX) y RSA PUBLIC KEY (PKCS#1)
	if block.Type == "RSA PUBLIC KEY" {
		return x509.ParsePKCS1PublicKey(block.Bytes)
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

func (s *JWKSService) toRSAJWK(pub *rsa.PublicKey, kid, alg string) dto.RSAJWK {
	// n y e en Base64URL sin padding
	n := base64.RawURLEncoding.EncodeToString(pub.N.Bytes())

	// exponent (e) en big-endian
	eBytes := big.NewInt(int64(pub.E)).Bytes()
	e := base64.RawURLEncoding.EncodeToString(eBytes)

	return dto.RSAJWK{
		Kty: "RSA",
		Kid: kid,
		Alg: alg,   // "RS256"
		Use: "sig", // firma
		N:   n,
		E:   e, // 65537 -> "AQAB"
	}
}

func (s *JWKSService) buildJWKS() (*dto.JWKSResponseDTO, error) {
	pemBytes, err := os.ReadFile(s.pubPath)
	if err != nil {
		return nil, err
	}
	rsaPub, err := s.parseRSAPublicKeyFromPEMBytes(pemBytes)
	if err != nil {
		return nil, err
	}

	jwk := s.toRSAJWK(rsaPub, s.kid, s.alg)
	set := &dto.JWKSResponseDTO{Keys: []dto.RSAJWK{jwk}}
	return set, nil
}

// Get devuelve el JWKS desde caché (si no expiró) o lo reconstruye.
// Retorna (jwks, maxAgeSeconds, error)
func (s *JWKSService) Get() (*dto.JWKSResponseDTO, int, error) {
	now := time.Now()
	s.mu.RLock()
	if s.cached != nil && now.Before(s.expiresAt) {
		defer s.mu.RUnlock()
		return s.cached, s.maxAge, nil
	}
	s.mu.RUnlock()

	s.mu.Lock()
	defer s.mu.Unlock()

	// doble chequeo
	if s.cached != nil && now.Before(s.expiresAt) {
		return s.cached, s.maxAge, nil
	}

	set, err := s.buildJWKS()
	if err != nil {
		return nil, 0, err
	}
	s.cached = set
	s.expiresAt = now.Add(time.Duration(s.maxAge) * time.Second)
	return s.cached, s.maxAge, nil
}
