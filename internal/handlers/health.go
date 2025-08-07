package handlers

import (
	"context"
	"net/http"

	"github.com/gofiber/fiber/v3"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/utils"
)

// HealthService define la interfaz para el chequeo de salud.
type HealthService interface {
	Check(ctx context.Context) (*dto.HealthResponseDTO, error)
}

// HealthHandler maneja las solicitudes de chequeo de salud.
type HealthHandler struct {
	healthService HealthService
}

// NewHealthHandler crea una nueva instancia de HealthHandler.
func NewHealthHandler(healthService HealthService) *HealthHandler {
	return &HealthHandler{
		healthService: healthService,
	}
}

// Health maneja GET /auth/health: verifica el estado del servicio.
func (h *HealthHandler) Health(c fiber.Ctx) error {
	// 1. Ejecutar el chequeo de salud usando el contexto de Fiber
	respDTO, err := h.healthService.Check(c)
	if err != nil {
		logger.Log.Errorf("Health check failed: %v", err)
		return utils.JSONError(c, http.StatusInternalServerError, "HEALTH_CHECK_FAILED", "Error interno al verificar salud")
	}

	// 2. Devolver la respuesta
	return utils.JSONResponse(c, http.StatusOK, true, "Servicio operativo", respDTO, nil)
}
