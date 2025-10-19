package handlers

import (
	"time"

	"central-user-manager/internal/config"
	"central-user-manager/internal/dto"
	"central-user-manager/internal/mapper"
	"central-user-manager/internal/models"
	"central-user-manager/pkg/logger"
	"central-user-manager/pkg/validator"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func AuthRoleHandler(c fiber.Ctx) error {
	var input dto.AuthRoleRequest

	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "Datos mal formateados"})
	}
	if err := validator.Validate.Struct(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ValidationErrorResponse{Errors: validator.FormatValidationError(err)})
	}

	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(dto.ErrorResponse{Error: "UserID inválido"})
	}

	db := config.DB

	var app models.Application
	if err := db.Where("client_id = ?", input.ClientID).First(&app).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "Aplicación no encontrada"})
		}
		logger.Log.Error("Error buscando la aplicación:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}

	var userAppRole models.UserApplicationRole
	if err := db.
		Preload("ApplicationRole").
		Where("user_id = ? AND application_id = ? AND is_deleted = false", userID, app.ID).
		First(&userAppRole).Error; err != nil {

		if err == gorm.ErrRecordNotFound {
			return c.Status(fiber.StatusNotFound).JSON(dto.ErrorResponse{Error: "No se encontró un rol asignado para el usuario en esta aplicación"})
		}
		logger.Log.Error("Error consultando rol del usuario:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "Error interno del servidor"})
	}
	role := userAppRole.ApplicationRole
	if role == nil {
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "Inconsistencia: rol de aplicación no disponible"})
	}

	var modulePerms []models.ModuleRolePermission
	if err := db.
		Preload("Module.Children").
		Preload("Module.Parent").
		Where("application_role_id = ? AND is_deleted = false", role.ID).
		Find(&modulePerms).Error; err != nil {

		logger.Log.Error("Error obteniendo módulos del rol:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "Error obteniendo módulos del rol"})
	}

	now := time.Now()
	var restrictions []models.UserModuleRestriction
	if err := db.
		Where("user_id = ? AND application_id = ? AND is_deleted = false AND (expires_at IS NULL OR expires_at > ?)", userID, app.ID, now).
		Find(&restrictions).Error; err != nil {

		logger.Log.Error("Error obteniendo restricciones de usuario:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(dto.ErrorResponse{Error: "Error obteniendo restricciones de usuario"})
	}

	restricted := make(map[uuid.UUID]bool, len(restrictions))
	for _, r := range restrictions {
		restricted[r.ModuleID] = true
	}

	modulesMap := make(map[uuid.UUID]dto.ModuleDTO)
	for _, mp := range modulePerms {
		if mp.Module == nil {
			continue
		}
		if restricted[mp.ModuleID] {
			continue
		}
		mod := mapper.ToModuleDTO(mp.Module)
		modulesMap[mod.ID] = mod
	}

	modules := make([]dto.ModuleDTO, 0, len(modulesMap))
	for _, m := range modulesMap {
		modules = append(modules, m)
	}

	return c.Status(fiber.StatusOK).JSON(dto.AuthRoleResponse{
		RoleID:  role.ID.String(),
		Modules: modules,
	})
}
