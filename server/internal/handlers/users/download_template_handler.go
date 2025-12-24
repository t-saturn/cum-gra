// internal/handlers/users/download_template_handler.go
package handlers

import (
	"fmt"
	"server/internal/dto"
	services "server/internal/services/users"
	"server/pkg/logger"
	"time"

	"github.com/gofiber/fiber/v3"
)

func DownloadUsersTemplateHandler(c fiber.Ctx) error {
	// Generar el archivo Excel
	file, err := services.GenerateUsersTemplateExcel()
	if err != nil {
		logger.Log.Error("Error generando plantilla Excel:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error generando plantilla"})
	}
	defer file.Close()

	// Crear buffer para el archivo
	buffer, err := file.WriteToBuffer()
	if err != nil {
		logger.Log.Error("Error escribiendo archivo Excel:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error generando archivo"})
	}

	// Generar nombre de archivo con timestamp
	filename := fmt.Sprintf("plantilla_usuarios_%s.xlsx", time.Now().Format("20060102_150405"))

	// Configurar headers para descarga
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Set("Content-Length", fmt.Sprintf("%d", buffer.Len()))

	return c.Send(buffer.Bytes())
}