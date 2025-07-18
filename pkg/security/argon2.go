// Package security proporciona utilidades para encriptación y verificación de contraseñas usando Argon2.
package security

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

// Argon2Service ofrece funciones para hashear y verificar contraseñas con el algoritmo Argon2id.
type Argon2Service struct{}

// NewArgon2Service crea una nueva instancia de Argon2Service.
func NewArgon2Service() *Argon2Service {
	return &Argon2Service{}
}

// HashPassword genera un hash seguro a partir de una contraseña usando Argon2id.
func (a *Argon2Service) HashPassword(password string) (string, error) {
	salt := make([]byte, 16)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	encoded := fmt.Sprintf("$argon2id$v=19$m=65536,t=1,p=4$%s$%s",
		base64.RawStdEncoding.EncodeToString(salt),
		base64.RawStdEncoding.EncodeToString(hash),
	)

	return encoded, nil
}

// CheckPasswordHash compara una contraseña con su hash codificado y devuelve si coinciden.
func (a *Argon2Service) CheckPasswordHash(password, encodedHash string) bool {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false
	}
	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false
	}

	hash := argon2.IDKey([]byte(password), salt, 1, 64*1024, 4, 32)

	return subtleCompare(hash, expectedHash)
}

func subtleCompare(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	var result byte
	for i := range a {
		result |= a[i] ^ b[i]
	}
	return result == 0
}
