package handlers

import (
	"fmt"
	"server/internal/dto"
	services "server/internal/services/organic-units"
	"server/pkg/logger"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/xuri/excelize/v2"
)

func BulkCreateOrganicUnitsHandler(c fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "No se proporcionó archivo Excel"})
	}

	if !strings.HasSuffix(strings.ToLower(file.Filename), ".xlsx") &&
		!strings.HasSuffix(strings.ToLower(file.Filename), ".xls") {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "El archivo debe ser formato Excel (.xlsx o .xls)"})
	}

	openedFile, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(dto.ErrorResponse{Error: "Error abriendo archivo"})
	}
	defer openedFile.Close()

	f, err := excelize.OpenReader(openedFile)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "Error leyendo archivo Excel"})
	}
	defer f.Close()

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

	units := make([]services.BulkOrganicUnitItem, 0, len(rows)-1)

	for i, row := range rows {
		if i == 0 {
			continue
		}

		if len(row) < 2 {
			logger.Log.Warn(fmt.Sprintf("Fila %d ignorada: datos insuficientes", i+1))
			continue
		}

		unit := services.BulkOrganicUnitItem{
			Row:     i + 1,
			Name:    strings.TrimSpace(row[0]),
			Acronym: strings.TrimSpace(row[1]),
		}

		// Campos opcionales
		if len(row) > 2 && row[2] != "" {
			brand := strings.TrimSpace(row[2])
			unit.Brand = &brand
		}
		if len(row) > 3 && row[3] != "" {
			desc := strings.TrimSpace(row[3])
			unit.Description = &desc
		}
		if len(row) > 4 && row[4] != "" {
			if parentID, err := strconv.ParseUint(strings.TrimSpace(row[4]), 10, 32); err == nil {
				pid := uint(parentID)
				unit.ParentID = &pid
			}
		}
		if len(row) > 5 && row[5] != "" {
			isActive := strings.ToLower(strings.TrimSpace(row[5])) == "true"
			unit.IsActive = &isActive
		}
		if len(row) > 6 && row[6] != "" {
			codSgd := strings.TrimSpace(row[6])
			unit.CodDepSGD = &codSgd
		}

		units = append(units, unit)
	}

	if len(units) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "No se encontraron unidades válidas en el archivo"})
	}

	if len(units) > 500 {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "Máximo 500 registros por carga"})
	}

	result := services.BulkCreateOrganicUnits(units)

	return c.Status(fiber.StatusOK).JSON(result)
}