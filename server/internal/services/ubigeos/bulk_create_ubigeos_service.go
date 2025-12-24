package services

import (
	"fmt"
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	"time"
)

type BulkUbigeoItem struct {
	Row        int
	UbigeoCode string
	IneiCode   string
	Department string
	Province   string
	District   string
}

func BulkCreateUbigeos(ubigeos []BulkUbigeoItem) *dto.BulkUploadResponse {
	db := config.DB
	response := &dto.BulkUploadResponse{
		Errors:  []dto.BulkUploadError{},
		Details: []map[string]any{},
	}

	for _, item := range ubigeos {
		// Validar campos requeridos
		if item.UbigeoCode == "" || item.Department == "" || item.Province == "" || item.District == "" {
			response.Failed++
			response.Errors = append(response.Errors, dto.BulkUploadError{
				Row:     item.Row,
				Message: "Campos obligatorios vacíos",
			})
			continue
		}

		// Verificar duplicado
		var exists int64
		db.Model(&models.Ubigeo{}).Where("ubigeo_code = ?", item.UbigeoCode).Count(&exists)
		if exists > 0 {
			response.Skipped++
			response.Errors = append(response.Errors, dto.BulkUploadError{
				Row:     item.Row,
				Field:   "ubigeo_code",
				Message: fmt.Sprintf("Ubigeo %s ya existe", item.UbigeoCode),
			})
			continue
		}

		ubigeo := models.Ubigeo{
			UbigeoCode: item.UbigeoCode,
			IneiCode:   item.IneiCode,
			Department: item.Department,
			Province:   item.Province,
			District:   item.District,
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
		}

		if err := db.Create(&ubigeo).Error; err != nil {
			response.Failed++
			response.Errors = append(response.Errors, dto.BulkUploadError{
				Row:     item.Row,
				Message: err.Error(),
			})
			continue
		}

		response.Created++
		response.Details = append(response.Details, map[string]any{
			"id":          ubigeo.ID,
			"ubigeo_code": ubigeo.UbigeoCode,
			"district":    ubigeo.District,
		})
	}

	return response
}