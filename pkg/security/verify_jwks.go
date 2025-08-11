// verify_jwks.go
package security

import (
	"encoding/json"
	"errors"
	"net/http"
	"sync"
	"time"

	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"

	"github.com/t-saturn/auth-service-server/internal/config"
)

var (
	jwksMu  sync.RWMutex
	jwksSet *jose.JSONWebKeySet
	jwksExp time.Time
)

func fetchJWKS(url string) (*jose.JSONWebKeySet, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var set jose.JSONWebKeySet
	if err := json.NewDecoder(resp.Body).Decode(&set); err != nil {
		return nil, err
	}
	// Lee max-age para cache simple
	if cc := resp.Header.Get("Cache-Control"); cc != "" {
		// simplificado: fija 60s si no parseas exacto
		jwksExp = time.Now().Add(60 * time.Second)
	} else {
		jwksExp = time.Now().Add(60 * time.Second)
	}
	return &set, nil
}

func getJWKS() (*jose.JSONWebKeySet, error) {
	cfg := config.GetConfig()
	jwksMu.RLock()
	if jwksSet != nil && time.Now().Before(jwksExp) {
		defer jwksMu.RUnlock()
		return jwksSet, nil
	}
	jwksMu.RUnlock()

	jwksMu.Lock()
	defer jwksMu.Unlock()
	// doble check
	if jwksSet != nil && time.Now().Before(jwksExp) {
		return jwksSet, nil
	}
	set, err := fetchJWKS(cfg.Server.JWKSURL) // p.ej. http://localhost:9190/.well-known/jwks.json
	if err != nil {
		return nil, err
	}
	jwksSet = set
	return jwksSet, nil
}

func VerifyTokenWithJWKS(tokenStr string) (jwt.Claims, error) {
	// 1 Parse para obtener header y KID
	parsed, err := jwt.ParseSigned(tokenStr)
	if err != nil {
		return jwt.Claims{}, ErrTokenInvalid
	}
	var rawHeader map[string]any
	headerBytes, err := json.Marshal(parsed.Headers[0])
	if err != nil {
		return jwt.Claims{}, ErrTokenInvalid
	}
	if err := json.Unmarshal(headerBytes, &rawHeader); err != nil {
		return jwt.Claims{}, ErrTokenInvalid
	}
	kid, _ := rawHeader["kid"].(string)
	if kid == "" {
		return jwt.Claims{}, ErrTokenInvalid
	}

	// 2 Resolver clave por KID desde JWKS (con cache)
	set, err := getJWKS()
	if err != nil {
		return jwt.Claims{}, ErrTokenInvalid
	}
	keys := set.Key(kid)
	if len(keys) == 0 {
		// fuerza refetch si expiro/cambio
		jwksMu.Lock()
		jwksSet = nil
		jwksMu.Unlock()
		// reintenta una vez
		set, _ = getJWKS()
		keys = set.Key(kid)
		if len(keys) == 0 {
			return jwt.Claims{}, errors.New("kid_not_found")
		}
	}

	// 3 Verificar firma y claims estándar
	var claims jwt.Claims
	if err := parsed.Claims(keys[0].Key, &claims); err != nil {
		return jwt.Claims{}, ErrTokenInvalid
	}
	cfg := config.GetConfig()
	if err := claims.Validate(jwt.Expected{
		Time:   time.Now(),
		Issuer: cfg.Server.JWTIss,
		// opcional: Audience: jwt.Audience{expectedAud},
	}); err != nil {
		if claims.Expiry != nil && time.Now().After(claims.Expiry.Time()) {
			return jwt.Claims{}, ErrTokenExpired
		}
		return jwt.Claims{}, ErrTokenInvalid
	}
	return claims, nil
}
