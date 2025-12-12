package mapper

import (
	"fmt"

	"server/internal/dto"
	"server/internal/models"
)

func ToApplicationRoleDTO(
	role models.ApplicationRole,
	app *models.Application,
	modulesCount int64,
	usersCount int64,
) dto.ApplicationRoleDTO {
	var deletedBy *string
	if role.DeletedBy != nil {
		s := fmt.Sprint(*role.DeletedBy)
		deletedBy = &s
	}

	var appMinimal *dto.AppMinimalDTO
	if app != nil {
		appMinimal = &dto.AppMinimalDTO{
			ID:       app.ID,
			Name:     app.Name,
			ClientID: app.ClientID,
		}
	}

	return dto.ApplicationRoleDTO{
		ID:            fmt.Sprint(role.ID),
		Name:          role.Name,
		Description:   role.Description,
		ApplicationID: fmt.Sprint(role.ApplicationID),
		Application:   appMinimal,
		CreatedAt:     role.CreatedAt,
		UpdatedAt:     role.UpdatedAt,
		IsDeleted:     role.IsDeleted,
		DeletedAt:     role.DeletedAt,
		DeletedBy:     deletedBy,
		ModulesCount:  modulesCount,
		UsersCount:    usersCount,
	}
}