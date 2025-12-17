package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"
)

func GetUbigeos(page, pageSize int, department, province, district *string) (*dto.UbigeosListResponse, error) {
	db := config.DB

	query := db.Model(&models.Ubigeo{})

	if department != nil && *department != "" {
		query = query.Where("department ILIKE ?", "%"+*department+"%")
	}

	if province != nil && *province != "" {
		query = query.Where("province ILIKE ?", "%"+*province+"%")
	}

	if district != nil && *district != "" {
		query = query.Where("district ILIKE ?", "%"+*district+"%")
	}

	var total int64
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	var ubigeos []models.Ubigeo
	if err := query.
		Order("department ASC, province ASC, district ASC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&ubigeos).Error; err != nil {
		return nil, err
	}

	if len(ubigeos) == 0 {
		return &dto.UbigeosListResponse{
			Data:     []dto.UbigeoDTO{},
			Total:    total,
			Page:     page,
			PageSize: pageSize,
		}, nil
	}

	out := make([]dto.UbigeoDTO, 0, len(ubigeos))
	for _, ubigeo := range ubigeos {
		out = append(out, mapper.ToUbigeoDTO(ubigeo))
	}

	return &dto.UbigeosListResponse{
		Data:     out,
		Total:    total,
		Page:     page,
		PageSize: pageSize,
	}, nil
}