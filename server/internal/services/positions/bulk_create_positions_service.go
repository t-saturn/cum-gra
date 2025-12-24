package services

import (
	"fmt"
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	"time"
)

type BulkPositionItem struct {
	Row         int
	Name        string
	Code        string
	Level       *int
	Description *string
	IsActive    *bool
	CodCarSGD   *string
}

func BulkCreatePositions(positions []BulkPositionItem) *dto.BulkUploadResponse {
	db := config.DB
	response := &dto.BulkUploadResponse{
		Errors:  []dto.BulkUploadError{},
		Details: []map[string]any{},
	}

	for _, item := range positions {
		if item.Name == "" || item.Code == "" {
			response.Failed++
			response.Errors = append(response.Errors, dto.BulkUploadError{
				Row:     item.Row,
				Message: "name y code son obligatorios",
			})
			continue
		}

		// Verificar nombre único
		var exists int64
		db.Model(&models.StructuralPosition{}).Where("name = ? AND is_deleted = FALSE", item.Name).Count(&exists)
		if exists > 0 {
			response.Skipped++
			response.Errors = append(response.Errors, dto.BulkUploadError{
				Row:     item.Row,
				Field:   "name",
				Message: fmt.Sprintf("Posición '%s' ya existe", item.Name),
			})
			continue
		}

		// Verificar código único
		db.Model(&models.StructuralPosition{}).Where("code = ? AND is_deleted = FALSE", item.Code).Count(&exists)
		if exists > 0 {
			response.Skipped++
			response.Errors = append(response.Errors, dto.BulkUploadError{
				Row:     item.Row,
				Field:   "code",
				Message: fmt.Sprintf("Código '%s' ya existe", item.Code),
			})
			continue
		}

		isActive := true
		if item.IsActive != nil {
			isActive = *item.IsActive
		}

		position := models.StructuralPosition{
			Name:        item.Name,
			Code:        item.Code,
			Level:       item.Level,
			Description: item.Description,
			IsActive:    isActive,
			CodCarSGD:   item.CodCarSGD,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if err := db.Create(&position).Error; err != nil {
			response.Failed++
			response.Errors = append(response.Errors, dto.BulkUploadError{
				Row:     item.Row,
				Message: err.Error(),
			})
			continue
		}

		response.Created++
		response.Details = append(response.Details, map[string]any{
			"id":   position.ID,
			"name": position.Name,
			"code": position.Code,
		})
	}

	return response
}