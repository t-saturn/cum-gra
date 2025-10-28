package mapper

import (
	"central-user-manager/internal/dto"
	"central-user-manager/internal/models"
)

// ToModuleDTO convierte un *models.Module en un dto.ModuleDTO
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
