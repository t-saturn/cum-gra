package middlewares

import (
	"crypto/rsa"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"

	"server/internal/config"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
)

// KeycloakPublicKey estructura para almacenar las claves públicas de Keycloak
type KeycloakPublicKey struct {
	Kid string         `json:"kid"`
	Kty string         `json:"kty"`
	Alg string         `json:"alg"`
	Use string         `json:"use"`
	N   string         `json:"n"`
	E   string         `json:"e"`
	Key *rsa.PublicKey `json:"-"`
}

// KeycloakCerts estructura para las claves públicas de Keycloak
type KeycloakCerts struct {
	Keys []KeycloakPublicKey `json:"keys"`
}

// KeycloakClaims claims personalizados para el token de Keycloak
type KeycloakClaims struct {
	Email             string                            `json:"email"`
	EmailVerified     bool                              `json:"email_verified"`
	Name              string                            `json:"name"`
	PreferredUsername string                            `json:"preferred_username"`
	GivenName         string                            `json:"given_name"`
	FamilyName        string                            `json:"family_name"`
	RealmAccess       map[string]interface{}            `json:"realm_access"`
	ResourceAccess    map[string]map[string]interface{} `json:"resource_access"`
	jwt.RegisteredClaims
}

// KeycloakMiddleware gestiona la validación de tokens
type KeycloakMiddleware struct {
	certsURL      string
	realm         string
	issuer        string
	publicKeys    map[string]*rsa.PublicKey
	mu            sync.RWMutex
	lastFetch     time.Time
	cacheDuration time.Duration
}

var (
	keycloakMiddleware *KeycloakMiddleware
	once               sync.Once
)

// InitKeycloakMiddleware inicializa el middleware de Keycloak
func InitKeycloakMiddleware() error {
	fmt.Println("========== INICIANDO KEYCLOAK MIDDLEWARE ==========")
	
	var initErr error
	once.Do(func() {
		cfg := config.GetConfig()
		
		if cfg.KeycloakSSOURL == "" || cfg.KeycloakRealm == "" {
			fmt.Println("ERROR: Variables de entorno no configuradas")
			initErr = errors.New("KEYCLOAK_SSO_URL y KEYCLOAK_REALM son requeridos")
			return
		}

		issuer := fmt.Sprintf("%s/realms/%s", cfg.KeycloakSSOURL, cfg.KeycloakRealm)
		certsURL := fmt.Sprintf("%s/protocol/openid-connect/certs", issuer)

		keycloakMiddleware = &KeycloakMiddleware{
			certsURL:      certsURL,
			realm:         cfg.KeycloakRealm,
			issuer:        issuer,
			publicKeys:    make(map[string]*rsa.PublicKey),
			cacheDuration: 1 * time.Hour,
		}

		// Obtener las claves públicas al inicializar
		if err := keycloakMiddleware.fetchPublicKeys(); err != nil {
			fmt.Printf("ERROR obteniendo claves: %v\n", err)
			initErr = fmt.Errorf("error al obtener claves públicas: %w", err)
			return
		}
		
	})

	if initErr != nil {
		fmt.Printf("ERROR EN INICIALIZACIÓN: %v\n", initErr)
	}

	return initErr
}

// fetchPublicKeys obtiene las claves públicas del servidor de Keycloak
func (km *KeycloakMiddleware) fetchPublicKeys() error {
	km.mu.RLock()
	if time.Since(km.lastFetch) < km.cacheDuration && len(km.publicKeys) > 0 {
		km.mu.RUnlock()
		return nil
	}
	km.mu.RUnlock()

	resp, err := http.Get(km.certsURL)
	if err != nil {
		return fmt.Errorf("error al obtener certificados: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("error al obtener certificados, status: %d", resp.StatusCode)
	}

	var certs KeycloakCerts
	if err := json.NewDecoder(resp.Body).Decode(&certs); err != nil {
		return fmt.Errorf("error al decodificar certificados: %w", err)
	}

	newKeys := make(map[string]*rsa.PublicKey)
	for _, key := range certs.Keys {
		if key.Kty != "RSA" {
			continue
		}

		pubKey, err := km.buildRSAPublicKey(key.N, key.E)
		if err != nil {
			continue
		}

		newKeys[key.Kid] = pubKey
	}

	km.mu.Lock()
	km.publicKeys = newKeys
	km.lastFetch = time.Now()
	km.mu.Unlock()

	return nil
}

// buildRSAPublicKey construye una clave pública RSA desde N y E
func (km *KeycloakMiddleware) buildRSAPublicKey(nStr, eStr string) (*rsa.PublicKey, error) {
	nBytes, err := base64.RawURLEncoding.DecodeString(nStr)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar N: %w", err)
	}

	eBytes, err := base64.RawURLEncoding.DecodeString(eStr)
	if err != nil {
		return nil, fmt.Errorf("error al decodificar E: %w", err)
	}

	n := new(big.Int).SetBytes(nBytes)
	var e int
	for _, b := range eBytes {
		e = e<<8 + int(b)
	}

	return &rsa.PublicKey{
		N: n,
		E: e,
	}, nil
}

// getPublicKey obtiene la clave pública por kid
func (km *KeycloakMiddleware) getPublicKey(kid string) (*rsa.PublicKey, error) {
	km.mu.RLock()
	key, exists := km.publicKeys[kid]
	km.mu.RUnlock()

	if !exists {
		// Intentar refrescar las claves
		if err := km.fetchPublicKeys(); err != nil {
			return nil, err
		}

		km.mu.RLock()
		key, exists = km.publicKeys[kid]
		km.mu.RUnlock()

		if !exists {
			return nil, fmt.Errorf("clave pública no encontrada para kid: %s", kid)
		}
	}

	return key, nil
}

// ValidateToken valida el token JWT
func (km *KeycloakMiddleware) ValidateToken(tokenString string) (*KeycloakClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &KeycloakClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Verificar el algoritmo
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("método de firma inesperado: %v", token.Header["alg"])
		}

		// Obtener kid del header
		kid, ok := token.Header["kid"].(string)
		if !ok {
			return nil, errors.New("kid no encontrado en el header del token")
		}

		// Obtener la clave pública correspondiente
		return km.getPublicKey(kid)
	})

	if err != nil {
		return nil, fmt.Errorf("error al parsear token: %w", err)
	}

	claims, ok := token.Claims.(*KeycloakClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token inválido")
	}

	// Validar issuer
	if claims.Issuer != km.issuer {
		return nil, fmt.Errorf("issuer inválido: esperado %s, obtenido %s", km.issuer, claims.Issuer)
	}

	return claims, nil
}

func KeycloakAuth() fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token de autorización requerido",
			})
		}

		parts := strings.Split(authHeader, " ")
		
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Formato de token inválido",
			})
		}

		tokenString := parts[1]

		claims, err := keycloakMiddleware.ValidateToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": fmt.Sprintf("Token inválido: %v", err),
			})
		}


		c.Locals("user", claims)
		c.Locals("user_id", claims.Subject)
		c.Locals("email", claims.Email)
		c.Locals("username", claims.PreferredUsername)

		return c.Next()
	}
}

// RequireRealmRole middleware para validar roles del realm
func RequireRealmRole(requiredRole string) fiber.Handler {
	return func(c fiber.Ctx) error {
		claims, ok := c.Locals("user").(*KeycloakClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Usuario no autenticado",
			})
		}

		if claims.RealmAccess == nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Sin acceso al realm",
			})
		}

		roles, ok := claims.RealmAccess["roles"].([]interface{})
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Roles no encontrados",
			})
		}

		for _, role := range roles {
			if roleStr, ok := role.(string); ok && roleStr == requiredRole {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": fmt.Sprintf("Rol requerido: %s", requiredRole),
		})
	}
}

// RequireResourceRole middleware para validar roles de recursos específicos
func RequireResourceRole(resource string, requiredRole string) fiber.Handler {
	return func(c fiber.Ctx) error {
		claims, ok := c.Locals("user").(*KeycloakClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Usuario no autenticado",
			})
		}

		if claims.ResourceAccess == nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Sin acceso a recursos",
			})
		}

		resourceAccess, ok := claims.ResourceAccess[resource]
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": fmt.Sprintf("Sin acceso al recurso: %s", resource),
			})
		}

		roles, ok := resourceAccess["roles"].([]interface{})
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Roles no encontrados en el recurso",
			})
		}

		for _, role := range roles {
			if roleStr, ok := role.(string); ok && roleStr == requiredRole {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": fmt.Sprintf("Rol requerido en %s: %s", resource, requiredRole),
		})
	}
}

// HasAnyResourceRole middleware para validar si tiene algún rol en un recurso
func HasAnyResourceRole(resource string, requiredRoles []string) fiber.Handler {
	return func(c fiber.Ctx) error {
		claims, ok := c.Locals("user").(*KeycloakClaims)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Usuario no autenticado",
			})
		}

		if claims.ResourceAccess == nil {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Sin acceso a recursos",
			})
		}

		resourceAccess, ok := claims.ResourceAccess[resource]
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": fmt.Sprintf("Sin acceso al recurso: %s", resource),
			})
		}

		roles, ok := resourceAccess["roles"].([]interface{})
		if !ok {
			return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
				"error": "Roles no encontrados en el recurso",
			})
		}

		userRoles := make(map[string]bool)
		for _, role := range roles {
			if roleStr, ok := role.(string); ok {
				userRoles[roleStr] = true
			}
		}

		for _, requiredRole := range requiredRoles {
			if userRoles[requiredRole] {
				return c.Next()
			}
		}

		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"error": fmt.Sprintf("Se requiere alguno de estos roles en %s: %v", resource, requiredRoles),
		})
	}
}