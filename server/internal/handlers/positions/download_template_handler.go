package handlers

import (
	"fmt"
	"server/internal/dto"
	services "server/internal/services/positions"
	"server/pkg/logger"
	"time"

	"github.com/gofiber/fiber/v3"
)

func DownloadPositionsTemplateHandler(c fiber.Ctx) error {
	file, err := services.GeneratePositionsTemplateExcel()
	if err != nil {
		logger.Log.Error("Error generando plantilla:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error generando plantilla"})
	}
	defer file.Close()

	buffer, err := file.WriteToBuffer()
	if err != nil {
		logger.Log.Error("Error escribiendo archivo:", err)
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error generando archivo"})
	}

	filename := fmt.Sprintf("plantilla_posiciones_%s.xlsx", time.Now().Format("20060102_150405"))
	c.Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	c.Set("Content-Length", fmt.Sprintf("%d", buffer.Len()))

	return c.Send(buffer.Bytes())
}