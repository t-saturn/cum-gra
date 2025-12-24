package handlers

import (
	"fmt"
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	services "server/internal/services/users"
	"server/pkg/logger"
	"strings"

	"github.com/gofiber/fiber/v3"
	"github.com/xuri/excelize/v2"
)

func BulkCreateUsersHandler(c fiber.Ctx) error {
	// Obtener email del contexto
	email, ok := c.Locals("email").(string)
	if !ok || email == "" {
		return c.Status(fiber.StatusUnauthorized).
			JSON(dto.ErrorResponse{Error: "Usuario no autenticado"})
	}

	// Buscar el usuario en la BD
	var user models.User
	db := config.DB
	if err := db.Where("email = ? AND is_deleted = FALSE", email).First(&user).Error; err != nil {
		return c.Status(fiber.StatusUnauthorized).
			JSON(dto.ErrorResponse{Error: "Usuario no encontrado en el sistema"})
	}

	// Obtener access token del header Authorization
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(fiber.StatusUnauthorized).
			JSON(dto.ErrorResponse{Error: "Token no proporcionado"})
	}

	accessToken := strings.TrimPrefix(authHeader, "Bearer ")
	if accessToken == authHeader {
		return c.Status(fiber.StatusUnauthorized).
			JSON(dto.ErrorResponse{Error: "Formato de token inválido"})
	}

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

	// Parsear usuarios desde Excel
	users := make([]dto.CreateUserRequest, 0, len(rows)-1)

	for i, row := range rows {
		if i == 0 {
			// Skip header
			continue
		}

		if len(row) < 4 { // Mínimo: email, dni, first_name, last_name
			logger.Log.Warn(fmt.Sprintf("Fila %d ignorada: datos insuficientes", i+1))
			continue
		}

		email := strings.TrimSpace(row[0])
		dni := strings.TrimSpace(row[1])
		firstName := strings.TrimSpace(row[2])
		lastName := strings.TrimSpace(row[3])

		// Si no hay contraseña en la columna 4, usar el DNI como contraseña
		password := dni
		if len(row) > 4 && row[4] != "" {
			password = strings.TrimSpace(row[4])
		}

		user := dto.CreateUserRequest{
			Email:     email,
			DNI:       dni,
			FirstName: firstName,
			LastName:  lastName,
			Password:  password, // Usar DNI o la contraseña proporcionada
		}

		// Campos opcionales
		if len(row) > 5 && row[5] != "" {
			phone := strings.TrimSpace(row[5])
			user.Phone = &phone
		}
		if len(row) > 6 && row[6] != "" {
			status := strings.TrimSpace(row[6])
			user.Status = &status
		}
		if len(row) > 7 && row[7] != "" {
			codEmpSGD := strings.TrimSpace(row[7])
			user.CodEmpSGD = &codEmpSGD
		}
		if len(row) > 8 && row[8] != "" {
			posID := strings.TrimSpace(row[8])
			user.StructuralPositionID = &posID
		}
		if len(row) > 9 && row[9] != "" {
			ouID := strings.TrimSpace(row[9])
			user.OrganicUnitID = &ouID
		}
		if len(row) > 10 && row[10] != "" {
			ubID := strings.TrimSpace(row[10])
			user.UbigeoID = &ubID
		}

		users = append(users, user)
	}

	if len(users) == 0 {
		return c.Status(fiber.StatusBadRequest).
			JSON(dto.ErrorResponse{Error: "No se encontraron usuarios válidos en el archivo"})
	}

	// Procesar carga masiva
	result := services.BulkCreateUsers(users, user.ID, accessToken)

	return c.Status(fiber.StatusOK).JSON(result)
}