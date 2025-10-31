package mapper

import (
	"fmt"

	"server/internal/dto"
	"server/internal/models"
)

func ToSimpleStructuralPositionDTO(sp *models.StructuralPosition) *dto.SimpleStructuralPositionDTO {
	if sp == nil {
		return nil
	}
	return &dto.SimpleStructuralPositionDTO{
		ID:    fmt.Sprint(sp.ID),
		Name:  sp.Name,
		Code:  sp.Code,
		Level: sp.Level,
	}
}

func ToSimpleOrganicUnitDTO(ou *models.OrganicUnit) *dto.SimpleOrganicUnitDTO {
	if ou == nil {
		return nil
	}
	var parentID *string
	if ou.ParentID != nil {
		s := fmt.Sprint(*ou.ParentID)
		parentID = &s
	}
	return &dto.SimpleOrganicUnitDTO{
		ID:       fmt.Sprint(ou.ID),
		Name:     ou.Name,
		Acronym:  ou.Acronym,
		ParentID: parentID,
	}
}

func ToUserListItemDTO(u models.User) dto.UserListItemDTO {
	var deletedBy *string
	if u.DeletedBy != nil {
		s := fmt.Sprint(*u.DeletedBy)
		deletedBy = &s
	}
	return dto.UserListItemDTO{
		ID:        fmt.Sprint(u.ID),
		Email:     u.Email,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Phone:     u.Phone,
		DNI:       u.DNI,
		Status:    u.Status,

		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		IsDeleted: u.IsDeleted,
		DeletedAt: u.DeletedAt,
		DeletedBy: deletedBy,

		OrganicUnit:        ToSimpleOrganicUnitDTO(u.OrganicUnit),
		StructuralPosition: ToSimpleStructuralPositionDTO(u.StructuralPosition),
	}
}
