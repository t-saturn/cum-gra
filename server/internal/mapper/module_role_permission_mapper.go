package mapper

import (
	"fmt"

	"server/internal/dto"
	"server/internal/models"
)

func ToModuleRolePermissionDTO(
	mrp models.ModuleRolePermission,
	module *models.Module,
	role *models.ApplicationRole,
	app *models.Application,
) dto.ModuleRolePermissionDTO {
	var deletedBy *string
	if mrp.DeletedBy != nil {
		s := fmt.Sprint(*mrp.DeletedBy)
		deletedBy = &s
	}

	moduleName := ""
	var moduleRoute *string
	if module != nil {
		moduleName = module.Name
		moduleRoute = module.Route
	}

	roleName := ""
	if role != nil {
		roleName = role.Name
	}

	applicationName := ""
	applicationClientID := ""
	if app != nil {
		applicationName = app.Name
		applicationClientID = app.ClientID
	}

	return dto.ModuleRolePermissionDTO{
		ID:                  fmt.Sprint(mrp.ID),
		ModuleID:            fmt.Sprint(mrp.ModuleID),
		ApplicationRoleID:   fmt.Sprint(mrp.ApplicationRoleID),
		PermissionType:      mrp.PermissionType,
		CreatedAt:           mrp.CreatedAt,
		IsDeleted:           mrp.IsDeleted,
		DeletedAt:           mrp.DeletedAt,
		DeletedBy:           deletedBy,
		ModuleName:          moduleName,
		ModuleRoute:         moduleRoute,
		RoleName:            roleName,
		ApplicationName:     applicationName,
		ApplicationClientID: applicationClientID,
	}
}