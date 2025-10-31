package mapper

import (
	"fmt"

	"server/internal/dto"
	"server/internal/models"
)

func ToOrganicUnitItemDTO(ou models.OrganicUnit, usersCount int64) dto.OrganicUnitItemDTO {
	var parentID *string
	if ou.ParentID != nil {
		s := fmt.Sprint(*ou.ParentID)
		parentID = &s
	}
	var deletedBy *string
	if ou.DeletedBy != nil {
		s := fmt.Sprint(*ou.DeletedBy)
		deletedBy = &s
	}

	return dto.OrganicUnitItemDTO{
		ID:          fmt.Sprint(ou.ID),
		Name:        ou.Name,
		Acronym:     ou.Acronym,
		Brand:       ou.Brand,
		Description: ou.Description,
		ParentID:    parentID,
		IsActive:    ou.IsActive,
		CreatedAt:   ou.CreatedAt,
		UpdatedAt:   ou.UpdatedAt,
		IsDeleted:   ou.IsDeleted,
		DeletedAt:   ou.DeletedAt,
		DeletedBy:   deletedBy,
		UsersCount:  usersCount,
	}
}
