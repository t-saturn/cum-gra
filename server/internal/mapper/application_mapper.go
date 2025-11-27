package mapper

import (
	"fmt"
	"strings"

	"server/internal/dto"
	"server/internal/models"
)

func ToAdminUserDTO(u models.User) dto.AdminUserDTO {
	first := ""
	last := ""

	full := strings.TrimSpace(strings.Join([]string{first, last}, " "))
	if full == "" {
		// fallback si no hay nombres
		full = u.Email
	}

	return dto.AdminUserDTO{
		FullName: full,
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

		CreatedAt: a.CreatedAt,
		UpdatedAt: a.UpdatedAt,
		IsDeleted: a.IsDeleted,
		DeletedAt: a.DeletedAt,
		DeletedBy: deletedBy,

		Admins:     admins,
		UsersCount: usersCount,
	}
}
