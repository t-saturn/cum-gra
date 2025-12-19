package handlers

import (
	"server/internal/dto"
	services "server/internal/services/ubigeos"
	"server/pkg/logger"

	"github.com/gofiber/fiber/v3"
)

func UpdateUbigeoHandler(c fiber.Ctx) error {
	id := c.Params("id")

	var req dto.UpdateUbigeoRequest
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "Datos inválidos"})
	}

	result, err := services.UpdateUbigeo(id, req)
	if err != nil {
		if err.Error() == "ubigeo no encontrado" || err.Error() == "ID inválido" {
			return c.Status(fiber.StatusNotFound).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		if err.Error() == "ya existe un ubigeo con este código" {
			return c.Status(fiber.StatusConflict).
				JSON(dto.ErrorResponse{Error: err.Error()})
		}
		logger.Log.Error("Error actualizando ubigeo:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	return c.Status(fiber.StatusOK).JSON(result)
}