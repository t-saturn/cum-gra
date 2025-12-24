package handlers

import (
	"fmt"
	"server/internal/dto"
	services "server/internal/services/ubigeos"
	"server/pkg/logger"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/xuri/excelize/v2"
)

func BulkCreateUbigeosHandler(c fiber.Ctx) error {
	// Obtener el archivo Excel
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "No se proporcionó archivo Excel"})
	}

	// Validar extensión
	if !strings.HasSuffix(strings.ToLower(file.Filename), ".xlsx") &&
		!strings.HasSuffix(strings.ToLower(file.Filename), ".xls") {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "El archivo debe ser formato Excel (.xlsx o .xls)"})
	}

	// Abrir archivo
	openedFile, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error abriendo archivo"})
	}
	defer openedFile.Close()

	// Leer Excel
	f, err := excelize.OpenReader(openedFile)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "Error leyendo archivo Excel"})
	}
	defer f.Close()

	// Obtener primera hoja
	sheets := f.GetSheetList()
	if len(sheets) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "El archivo Excel no contiene hojas"})
	}

	rows, err := f.GetRows(sheets[0])
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "Error leyendo filas del Excel"})
	}

	if len(rows) < 2 {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "El archivo debe contener al menos una fila de datos"})
	}

	// Parsear ubigeos desde Excel
	ubigeos := make([]services.BulkUbigeoItem, 0, len(rows)-1)

	for i, row := range rows {
		if i == 0 {
			continue // Skip header
		}

		if len(row) < 5 {
			logger.Log.Warn(fmt.Sprintf("Fila %d ignorada: datos insuficientes", i+1))
			continue
		}

		ubigeo := services.BulkUbigeoItem{
			Row:        i + 1,
			UbigeoCode: strings.TrimSpace(row[0]),
			IneiCode:   strings.TrimSpace(row[1]),
			Department: strings.TrimSpace(row[2]),
			Province:   strings.TrimSpace(row[3]),
			District:   strings.TrimSpace(row[4]),
		}

		ubigeos = append(ubigeos, ubigeo)
	}

	if len(ubigeos) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "No se encontraron ubigeos válidos en el archivo"})
	}

	if len(ubigeos) > 1000 {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "Máximo 1000 registros por carga"})
	}

	// Procesar carga masiva
	result := services.BulkCreateUbigeos(ubigeos)

	return c.Status(fiber.StatusOK).JSON(result)
}