package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/security"
)

type AuthRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AuthVerifyHandler(c fiber.Ctx) error {
	var req AuthRequest
	if err := c.Bind().Body(&req); err != nil {
		logger.Log.Errorf("Error parseando body: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Formato incorrecto"})
	}

	payload, _ := json.Marshal(req)
	resp, err := http.Post("http://localhost:8080/auth/verify", "application/json", bytes.NewBuffer(payload))
	if err != nil {
		logger.Log.Errorf("Error al llamar al verificador: %v", err)
		return c.Status(fiber.StatusBadGateway).JSON(fiber.Map{"error": "Error verificando credenciales"})
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		logger.Log.Warnf("Verificación fallida: %s", string(body))
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Credenciales inválidas"})
	}

	var result map[string]string
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		logger.Log.Errorf("Error decodificando respuesta: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error interno"})
	}

	userID := result["user_id"]
	token, err := security.GenerateToken(userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error generando token"})
	}

	logger.Log.Infof("Token generado para userID: %s", userID)
	return c.JSON(fiber.Map{"token": token})
}

// TokenValidationRequest representa el JSON recibido con el token
type TokenValidationRequest struct {
	Token string `json:"token"`
}

// ValidateTokenHandler recibe un token y devuelve true/false según validez
func ValidateTokenHandler(c fiber.Ctx) error {
	var req TokenValidationRequest

	if err := c.Bind().Body(&req); err != nil {
		logger.Log.Warnf("Body inválido: %v", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"valid":   false,
			"message": "Body inválido",
		})
	}

	result := security.ValidateToken(req.Token)

	if result.Code == 0 {
		logger.Log.Infof("Token válido para subject: %s", result.Claims.Subject)
		return c.JSON(fiber.Map{
			"valid":   true,
			"subject": result.Claims.Subject,
			"message": "Token válido",
		})
	}

	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
		"valid":   false,
		"message": result.Message,
	})
}
