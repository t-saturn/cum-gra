package mapper

import (
	"fmt"

	"server/internal/dto"
	"server/internal/models"
)

func ToStructuralPositionItemDTO(r models.StructuralPositionRow) dto.StructuralPositionItemDTO {
	var deletedByStr *string
	if r.DeletedBy != nil {
		s := fmt.Sprint(*r.DeletedBy)
		deletedByStr = &s
	}

	return dto.StructuralPositionItemDTO{
		ID:          fmt.Sprint(r.ID),
		Name:        r.Name,
		Code:        r.Code,
		Level:       r.Level,
		Description: r.Description,
		IsActive:    r.IsActive,
		CreatedAt:   r.CreatedAt,
		UpdatedAt:   r.UpdatedAt,
		IsDeleted:   r.IsDeleted,
		DeletedAt:   r.DeletedAt,
		DeletedBy:   deletedByStr,
		CodCarSGD:   r.CodCarSGD,
		UsersCount:  r.UsersCount,
	}
}

func ToStructuralPositionListDTO(rows []models.StructuralPositionRow) []dto.StructuralPositionItemDTO {
	out := make([]dto.StructuralPositionItemDTO, 0, len(rows))
	for _, r := range rows {
		out = append(out, ToStructuralPositionItemDTO(r))
	}
	return out
}

func StructuralPositionToRow(pos models.StructuralPosition, usersCount int64) models.StructuralPositionRow {
	return models.StructuralPositionRow{
		StructuralPosition: pos,
		UsersCount:         usersCount,
	}
}