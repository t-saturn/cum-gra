package handlers

import (
	"server/internal/dto"
	services "server/internal/services/ubigeos"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func CreateUbigeoHandler(c fiber.Ctx) error {
	var req dto.CreateUbigeoRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "Datos inválidos"})
	}

	result, err := services.CreateUbigeo(req)
	if err != nil {
		if err.Error() == "ya existe un ubigeo con este código" {
			return c.Status(fiber.StatusConflict).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error creando ubigeo:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusCreated).JSON(result)
}