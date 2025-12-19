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

func ToSimpleUbigeoDTO(ub *models.Ubigeo) *dto.SimpleUbigeoDTO {
	if ub == nil {
		return nil
	}
	return &dto.SimpleUbigeoDTO{
		ID:         fmt.Sprint(ub.ID),
		UbigeoCode: ub.UbigeoCode,
		Department: ub.Department,
		Province:   ub.Province,
		District:   ub.District,
	}
}

func ToUserListItemDTO(u models.User, detail *models.UserDetail) dto.UserListItemDTO {
	var deletedBy *string
	if u.DeletedBy != nil {
		s := fmt.Sprint(*u.DeletedBy)
		deletedBy = &s
	}

	result := dto.UserListItemDTO{
		ID:        fmt.Sprint(u.ID),
		Email:     u.Email,
		DNI:       u.DNI,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		IsDeleted: u.IsDeleted,
		DeletedAt: u.DeletedAt,
		DeletedBy: deletedBy,
	}

	if detail != nil {
		result.FirstName = detail.FirstName
		result.LastName = detail.LastName
		result.Phone = detail.Phone
		result.StructuralPosition = ToSimpleStructuralPositionDTO(detail.StructuralPosition)
		result.OrganicUnit = ToSimpleOrganicUnitDTO(detail.OrganicUnit)
		result.Ubigeo = ToSimpleUbigeoDTO(detail.Ubigeo)
	}

	return result
}

func ToUserDetailDTO(u models.User, detail *models.UserDetail) dto.UserDetailDTO {
	var deletedBy *string
	if u.DeletedBy != nil {
		s := fmt.Sprint(*u.DeletedBy)
		deletedBy = &s
	}

	result := dto.UserDetailDTO{
		ID:        fmt.Sprint(u.ID),
		Email:     u.Email,
		DNI:       u.DNI,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
		IsDeleted: u.IsDeleted,
		DeletedAt: u.DeletedAt,
		DeletedBy: deletedBy,
	}

	if detail != nil {
		result.FirstName = detail.FirstName
		result.LastName = detail.LastName
		result.Phone = detail.Phone
		result.CodEmpSGD = detail.CodEmpSGD

		if detail.StructuralPositionID != nil {
			s := fmt.Sprint(*detail.StructuralPositionID)
			result.StructuralPositionID = &s
		}
		if detail.OrganicUnitID != nil {
			s := fmt.Sprint(*detail.OrganicUnitID)
			result.OrganicUnitID = &s
		}
		if detail.UbigeoID != nil {
			s := fmt.Sprint(*detail.UbigeoID)
			result.UbigeoID = &s
		}

		result.StructuralPosition = ToSimpleStructuralPositionDTO(detail.StructuralPosition)
		result.OrganicUnit = ToSimpleOrganicUnitDTO(detail.OrganicUnit)
		result.Ubigeo = ToSimpleUbigeoDTO(detail.Ubigeo)
	}

	return result
}