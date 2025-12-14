package mapper

import (
	"fmt"

	"server/internal/dto"
	"server/internal/models"
)

func ToUserModuleRestrictionDTO(
	restriction models.UserModuleRestriction,
	user *models.User,
	userDetail *models.UserDetail,
	module *models.Module,
	app *models.Application,
) dto.UserModuleRestrictionDTO {
	var updatedBy *string
	if restriction.UpdatedBy != nil {
		s := fmt.Sprint(*restriction.UpdatedBy)
		updatedBy = &s
	}

	var deletedBy *string
	if restriction.DeletedBy != nil {
		s := fmt.Sprint(*restriction.DeletedBy)
		deletedBy = &s
	}

	userEmail := ""
	userFullName := ""
	if user != nil {
		userEmail = user.Email
		if userDetail != nil {
			firstName := ""
			lastName := ""
			if userDetail.FirstName != nil {
				firstName = *userDetail.FirstName
			}
			if userDetail.LastName != nil {
				lastName = *userDetail.LastName
			}
			userFullName = fmt.Sprintf("%s %s", firstName, lastName)
		}
	}

	moduleName := ""
	var moduleRoute *string
	if module != nil {
		moduleName = module.Name
		moduleRoute = module.Route
	}

	applicationName := ""
	applicationClientID := ""
	if app != nil {
		applicationName = app.Name
		applicationClientID = app.ClientID
	}

	return dto.UserModuleRestrictionDTO{
		ID:                  fmt.Sprint(restriction.ID),
		UserID:              fmt.Sprint(restriction.UserID),
		ModuleID:            fmt.Sprint(restriction.ModuleID),
		ApplicationID:       fmt.Sprint(restriction.ApplicationID),
		RestrictionType:     restriction.RestrictionType,
		MaxPermissionLevel:  restriction.MaxPermissionLevel,
		Reason:              restriction.Reason,
		ExpiresAt:           restriction.ExpiresAt,
		CreatedAt:           restriction.CreatedAt,
		CreatedBy:           fmt.Sprint(restriction.CreatedBy),
		UpdatedAt:           restriction.UpdatedAt,
		UpdatedBy:           updatedBy,
		IsDeleted:           restriction.IsDeleted,
		DeletedAt:           restriction.DeletedAt,
		DeletedBy:           deletedBy,
		UserEmail:           userEmail,
		UserFullName:        userFullName,
		ModuleName:          moduleName,
		ModuleRoute:         moduleRoute,
		ApplicationName:     applicationName,
		ApplicationClientID: applicationClientID,
	}
}