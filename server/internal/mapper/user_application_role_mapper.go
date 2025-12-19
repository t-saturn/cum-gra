package mapper

import (
	"fmt"

	"server/internal/dto"
	"server/internal/models"
)

func ToUserApplicationRoleDTO(
	uar models.UserApplicationRole,
	user *models.User,
	userDetail *models.UserDetail,
	app *models.Application,
	role *models.ApplicationRole,
	grantedByUser *models.User,
	revokedByUser *models.User,
) dto.UserApplicationRoleDTO {
	var revokedBy *string
	if uar.RevokedBy != nil {
		s := fmt.Sprint(*uar.RevokedBy)
		revokedBy = &s
	}

	var deletedBy *string
	if uar.DeletedBy != nil {
		s := fmt.Sprint(*uar.DeletedBy)
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

	applicationName := ""
	applicationClientID := ""
	if app != nil {
		applicationName = app.Name
		applicationClientID = app.ClientID
	}

	roleName := ""
	if role != nil {
		roleName = role.Name
	}

	grantedByEmail := ""
	if grantedByUser != nil {
		grantedByEmail = grantedByUser.Email
	}

	var revokedByEmail *string
	if revokedByUser != nil {
		revokedByEmail = &revokedByUser.Email
	}

	return dto.UserApplicationRoleDTO{
		ID:                  fmt.Sprint(uar.ID),
		UserID:              fmt.Sprint(uar.UserID),
		ApplicationID:       fmt.Sprint(uar.ApplicationID),
		ApplicationRoleID:   fmt.Sprint(uar.ApplicationRoleID),
		GrantedAt:           uar.GrantedAt,
		GrantedBy:           fmt.Sprint(uar.GrantedBy),
		RevokedAt:           uar.RevokedAt,
		RevokedBy:           revokedBy,
		IsDeleted:           uar.IsDeleted,
		DeletedAt:           uar.DeletedAt,
		DeletedBy:           deletedBy,
		CreatedAt:           uar.CreatedAt,
		UpdatedAt:           uar.UpdatedAt,
		UserEmail:           userEmail,
		UserFullName:        userFullName,
		ApplicationName:     applicationName,
		ApplicationClientID: applicationClientID,
		RoleName:            roleName,
		GrantedByEmail:      grantedByEmail,
		RevokedByEmail:      revokedByEmail,
	}
}