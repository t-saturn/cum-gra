package mapper

import (
	"fmt"

	"server/internal/dto"
	"server/internal/models"
)

func ToModuleDTO(m *models.Module) dto.ModuleDTO {
	if m == nil {
		return dto.ModuleDTO{}
	}

	var parent *dto.SimpleModuleDTO
	if m.Parent != nil {
		parent = &dto.SimpleModuleDTO{
			ID:   m.Parent.ID,
			Name: m.Parent.Name,
		}
	}

	children := make([]dto.SimpleModuleDTO, 0, len(m.Children))
	for _, child := range m.Children {
		children = append(children, dto.SimpleModuleDTO{
			ID:   child.ID,
			Name: child.Name,
		})
	}

	return dto.ModuleDTO{
		ID:        m.ID,
		Item:      m.Item,
		Name:      m.Name,
		Route:     m.Route,
		Icon:      m.Icon,
		ParentID:  m.ParentID,
		SortOrder: m.SortOrder,
		Status:    m.Status,
		CreatedAt: m.CreatedAt,
		UpdatedAt: m.UpdatedAt,
		Parent:    parent,
		Children:  children,
	}
}

func ToModuleWithAppDTO(
	m models.Module,
	app *models.Application,
	usersCount int64,
) dto.ModuleWithAppDTO {
	var deletedBy *string
	if m.DeletedBy != nil {
		s := fmt.Sprint(*m.DeletedBy)
		deletedBy = &s
	}

	var parentID *string
	if m.ParentID != nil {
		s := fmt.Sprint(*m.ParentID)
		parentID = &s
	}

	var applicationID *string
	var applicationName *string
	var applicationClientID *string
	
	if m.ApplicationID != nil {
		s := fmt.Sprint(*m.ApplicationID)
		applicationID = &s
	}

	if app != nil {
		applicationName = &app.Name
		applicationClientID = &app.ClientID
	}

	return dto.ModuleWithAppDTO{
		ID:                  fmt.Sprint(m.ID),
		Item:                m.Item,
		Name:                m.Name,
		Route:               m.Route,
		Icon:                m.Icon,
		ParentID:            parentID,
		ApplicationID:       applicationID,
		SortOrder:           m.SortOrder,
		Status:              m.Status,
		CreatedAt:           m.CreatedAt,
		UpdatedAt:           m.UpdatedAt,
		DeletedAt:           m.DeletedAt,
		DeletedBy:           deletedBy,
		ApplicationName:     applicationName,
		ApplicationClientID: applicationClientID,
		UsersCount:          usersCount,
	}
}