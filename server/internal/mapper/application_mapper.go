package mapper

import (
	"fmt"

	"server/internal/dto"
	"server/internal/models"
)

func ToAdminUserDTO(u models.User) dto.AdminUserDTO {
	// Este método ya no se usa realmente, pero lo dejamos por compatibilidad
	return dto.AdminUserDTO{
		FullName: u.Email,
		DNI:      u.DNI,
		Email:    u.Email,
	}
}

func ToApplicationDTO(
	a models.Application,
	admins []dto.AdminUserDTO,
	usersCount int64,
) dto.ApplicationDTO {
	var deletedBy *string
	if a.DeletedBy != nil {
		s := fmt.Sprint(*a.DeletedBy)
		deletedBy = &s
	}

	return dto.ApplicationDTO{
		ID:          fmt.Sprint(a.ID),
		Name:        a.Name,
		ClientID:    a.ClientID,
		Domain:      a.Domain,
		Logo:        a.Logo,
		Description: a.Description,
		Status:      a.Status,
		CreatedAt:   a.CreatedAt,
		UpdatedAt:   a.UpdatedAt,
		IsDeleted:   a.IsDeleted,
		DeletedAt:   a.DeletedAt,
		DeletedBy:   deletedBy,
		Admins:      admins,
		UsersCount:  usersCount,
	}
}
