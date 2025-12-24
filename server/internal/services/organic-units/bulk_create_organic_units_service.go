package services

import (
	"fmt"
	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"
	"time"
)

type BulkOrganicUnitItem struct {
	Row         int
	Name        string
	Acronym     string
	Brand       *string
	Description *string
	ParentID    *uint
	IsActive    *bool
	CodDepSGD   *string
}

func BulkCreateOrganicUnits(units []BulkOrganicUnitItem) *dto.BulkUploadResponse {
	db := config.DB
	response := &dto.BulkUploadResponse{
		Errors:  []dto.BulkUploadError{},
		Details: []map[string]any{},
	}

	for _, item := range units {
		if item.Name == "" || item.Acronym == "" {
			response.Failed++
			response.Errors = append(response.Errors, dto.BulkUploadError{
				Row:     item.Row,
				Message: "name y acronym son obligatorios",
			})
			continue
		}

		// Verificar nombre único
		var exists int64
		db.Model(&models.OrganicUnit{}).Where("name = ? AND is_deleted = FALSE", item.Name).Count(&exists)
		if exists > 0 {
			response.Skipped++
			response.Errors = append(response.Errors, dto.BulkUploadError{
				Row:     item.Row,
				Field:   "name",
				Message: fmt.Sprintf("Unidad '%s' ya existe", item.Name),
			})
			continue
		}

		// Verificar acrónimo único
		db.Model(&models.OrganicUnit{}).Where("acronym = ? AND is_deleted = FALSE", item.Acronym).Count(&exists)
		if exists > 0 {
			response.Skipped++
			response.Errors = append(response.Errors, dto.BulkUploadError{
				Row:     item.Row,
				Field:   "acronym",
				Message: fmt.Sprintf("Acrónimo '%s' ya existe", item.Acronym),
			})
			continue
		}

		// Verificar parent_id si existe
		if item.ParentID != nil {
			var parent models.OrganicUnit
			if err := db.First(&parent, *item.ParentID).Error; err != nil {
				response.Failed++
				response.Errors = append(response.Errors, dto.BulkUploadError{
					Row:     item.Row,
					Field:   "parent_id",
					Message: fmt.Sprintf("Unidad padre %d no existe", *item.ParentID),
				})
				continue
			}
		}

		isActive := true
		if item.IsActive != nil {
			isActive = *item.IsActive
		}

		unit := models.OrganicUnit{
			Name:        item.Name,
			Acronym:     item.Acronym,
			Brand:       item.Brand,
			Description: item.Description,
			ParentID:    item.ParentID,
			IsActive:    isActive,
			CodDepSGD:   item.CodDepSGD,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}

		if err := db.Create(&unit).Error; err != nil {
			response.Failed++
			response.Errors = append(response.Errors, dto.BulkUploadError{
				Row:     item.Row,
				Message: err.Error(),
			})
			continue
		}

		response.Created++
		response.Details = append(response.Details, map[string]any{
			"id":      unit.ID,
			"name":    unit.Name,
			"acronym": unit.Acronym,
		})
	}

	return response
}