package security

import (
	"crypto/sha256"
	"crypto/subtle"
	"encoding/base64"
	"encoding/hex"
)

// HashTokenHex devuelve el SHA-256 del token en HEX (minúsculas)
func HashTokenHex(raw string) string {
	sum := sha256.Sum256([]byte(raw))
	return hex.EncodeToString(sum[:])
}

// (Opcional) HashTokenB64URL devuelve el SHA-256 en Base64URL sin padding
func HashTokenB64URL(raw string) string {
	sum := sha256.Sum256([]byte(raw))
	return base64.RawURLEncoding.EncodeToString(sum[:])
}

// (Opcional) CompareHash realiza comparación en tiempo constante
func CompareHash(a, b string) bool {
	// Ambas entradas deben tener el mismo encoding (HEX con HEX, o B64URL con B64URL)
	if len(a) != len(b) {
		return false
	}
	return subtle.ConstantTimeCompare([]byte(a), []byte(b)) == 1
}
