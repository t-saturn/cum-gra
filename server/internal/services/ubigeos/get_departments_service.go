package services

import (
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
)

func GetDepartments() ([]dto.DepartmentDTO, error) {
	db := config.DB

	var departments []string
	if err := db.Model(&models.Ubigeo{}).
		Distinct("department").
		Order("department ASC").
		Pluck("department", &departments).Error; err != nil {
		return nil, err
	}

	result := make([]dto.DepartmentDTO, 0, len(departments))
	for _, dept := range departments {
		result = append(result, dto.DepartmentDTO{
			Name: dept,
		})
	}

	return result, nil
}