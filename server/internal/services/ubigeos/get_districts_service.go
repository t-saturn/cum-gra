package services

import (
	"errors"
	"fmt"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
)

func GetDistrictsByProvince(department, province string) ([]dto.DistrictDTO, error) {
	db := config.DB

	if department == "" {
		return nil, errors.New("departamento requerido")
	}

	if province == "" {
		return nil, errors.New("provincia requerida")
	}

	var ubigeos []models.Ubigeo
	if err := db.Where("department = ? AND province = ?", department, province).
		Order("district ASC").
		Find(&ubigeos).Error; err != nil {
		return nil, err
	}

	districts := make([]dto.DistrictDTO, 0, len(ubigeos))
	for _, u := range ubigeos {
		districts = append(districts, dto.DistrictDTO{
			ID:         fmt.Sprint(u.ID),
			Name:       u.District,
			Province:   u.Province,
			Department: u.Department,
			UbigeoCode: u.UbigeoCode,
		})
	}

	return districts, nil
}