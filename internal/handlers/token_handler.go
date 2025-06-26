package handlers

import (
	"errors"
	"strings"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/internal/services"
)

type TokenRequest struct {
	UserID         string            `json:"user_id"`
	ApplicationID  string            `json:"application_id"`
	ApplicationURL string            `json:"application_url"`
	DeviceInfo     models.DeviceInfo `json:"deviceInfo"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
}

func GenerateTokenHandler(c fiber.Ctx) error {
	var body TokenRequest
	if err := c.Bind().Body(&body); err != nil || body.UserID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "user_id is required"})
	}

	tokenStr, err := services.GenerateAndStoreToken(
		body.UserID,
		body.ApplicationID,
		body.ApplicationURL,
		body.DeviceInfo,
	)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "failed to generate token"})
	}

	return c.JSON(fiber.Map{"token": tokenStr})
}

func ValidateTokenHandler(c fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Falta o mal formato en Authorization header"})
	}

	tokenStr := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.NewError(fiber.StatusUnauthorized, "Método de firma no permitido")
		}
		return []byte(config.GetEnv("JWT_SECRET", "")), nil
	})

	/* --- */
	var tokenErrorResponses = []struct {
		match  error
		status int
		msg    string
	}{
		{jwt.ErrTokenExpired, fiber.StatusUnauthorized, "Token expirado"},
		{jwt.ErrTokenSignatureInvalid, fiber.StatusUnauthorized, "Firma inválida del token"},
		{jwt.ErrTokenMalformed, fiber.StatusBadRequest, "Token mal formado"},
	}

	if err != nil {
		for _, te := range tokenErrorResponses {
			if errors.Is(err, te.match) {
				return c.Status(te.status).JSON(fiber.Map{"error": te.msg})
			}
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}
	/* if err != nil {
		switch {
		case errors.Is(err, jwt.ErrTokenExpired):
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token expirado"})
		case errors.Is(err, jwt.ErrTokenSignatureInvalid):
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Firma inválida del token"})
		case errors.Is(err, jwt.ErrTokenMalformed):
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Token mal formado"})
		default:
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}
	} */

	if !token.Valid {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Token inválido"})
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "No se pudo extraer claims del token"})
	}

	formatTime := func(ts interface{}) string {
		if v, ok := ts.(float64); ok {
			return time.Unix(int64(v), 0).Format(time.RFC3339)
		}
		return ""
	}

	return c.JSON(fiber.Map{
		"sub": claims["sub"],
		"jti": claims["jti"],
		"iat": formatTime(claims["iat"]),
		"exp": formatTime(claims["exp"]),
	})
}

func RefreshTokenHandler(c fiber.Ctx) error {
	var body RefreshTokenRequest
	if err := c.Bind().Body(&body); err != nil || body.RefreshToken == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "refresh_token is required"})
	}

	newAccessToken, err := services.RefreshAccessToken(body.RefreshToken)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{
		"access_token": newAccessToken,
	})
}
