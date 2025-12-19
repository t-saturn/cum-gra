package mapper

import (
	"fmt"

	"server/internal/dto"
	"server/internal/models"
)

func ToUbigeoDTO(ubigeo models.Ubigeo) dto.UbigeoDTO {
	return dto.UbigeoDTO{
		ID:         fmt.Sprint(ubigeo.ID),
		UbigeoCode: ubigeo.UbigeoCode,
		IneiCode:   ubigeo.IneiCode,
		Department: ubigeo.Department,
		Province:   ubigeo.Province,
		District:   ubigeo.District,
		CreatedAt:  ubigeo.CreatedAt,
		UpdatedAt:  ubigeo.UpdatedAt,
	}
}