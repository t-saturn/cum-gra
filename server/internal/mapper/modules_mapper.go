package mapper

import (
	"fmt"

	"central-user-manager/internal/dto"
	"central-user-manager/internal/models"
)

func ToModuleWithAppDTO(m models.Module, usersCount int64) dto.ModuleWithAppDTO {
	var parentID *string
	if m.ParentID != nil {
		s := fmt.Sprint(*m.ParentID)
		parentID = &s
	}
	var appID *string
	if m.ApplicationID != nil {
		s := fmt.Sprint(*m.ApplicationID)
		appID = &s
	}
	var deletedBy *string
	if m.DeletedBy != nil {
		s := fmt.Sprint(*m.DeletedBy)
		deletedBy = &s
	}

	var appName *string
	var appClientID *string
	if m.Application != nil {
		if m.Application.Name != "" {
			n := m.Application.Name
			appName = &n
		}
		if m.Application.ClientID != "" {
			c := m.Application.ClientID
			appClientID = &c
		}
	}

	return dto.ModuleWithAppDTO{
		ID:                  fmt.Sprint(m.ID),
		Item:                m.Item,
		Name:                m.Name,
		Route:               m.Route,
		Icon:                m.Icon,
		ParentID:            parentID,
		ApplicationID:       appID,
		SortOrder:           m.SortOrder,
		Status:              m.Status,
		CreatedAt:           m.CreatedAt,
		UpdatedAt:           m.UpdatedAt,
		DeletedAt:           m.DeletedAt,
		DeletedBy:           deletedBy,
		ApplicationName:     appName,
		ApplicationClientID: appClientID,
		UsersCount:          usersCount,
	}
}
