package services

import (
	"errors"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
)

func GetProvincesByDepartment(department string) ([]dto.ProvinceDTO, error) {
	db := config.DB

	if department == "" {
		return nil, errors.New("departamento requerido")
	}

	var results []struct {
		Province   string
		Department string
	}

	if err := db.Model(&models.Ubigeo{}).
		Select("DISTINCT province, department").
		Where("department = ?", department).
		Order("province ASC").
		Scan(&results).Error; err != nil {
		return nil, err
	}

	provinces := make([]dto.ProvinceDTO, 0, len(results))
	for _, r := range results {
		provinces = append(provinces, dto.ProvinceDTO{
			Name:       r.Province,
			Department: r.Department,
		})
	}

	return provinces, nil
}