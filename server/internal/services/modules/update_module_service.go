package services

import (
	"errors"
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UpdateModule(id string, req dto.UpdateModuleRequest, updatedBy uuid.UUID) (*dto.ModuleWithAppDTO, error) {
	db := config.DB

	moduleID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var module models.Module
	if delErr := db.Where("id = ? AND deleted_at IS NULL", moduleID).First(&module).Error; delErr != nil {
		if errors.Is(delErr, gorm.ErrRecordNotFound) {
			return nil, errors.New("módulo no encontrado")
		}
		return nil, delErr
	}

	// Flag para saber si hay que actualizar parent_id a NULL
	setParentToNull := false

	// Validar parent_id si se está actualizando
	if req.ParentID != nil {
		if *req.ParentID == "" {
			// Convertir a módulo raíz
			module.ParentID = nil
			setParentToNull = true
		} else {
			parsedParentID, parseErr := uuid.Parse(*req.ParentID)
			if parseErr != nil {
				return nil, errors.New("parent_id inválido")
			}
			if err != nil {
				return nil, errors.New("parent_id inválido")
			}
			
			// No puede ser su propio padre
			if parsedParentID == moduleID {
				return nil, errors.New("un módulo no puede ser su propio padre")
			}
			
			var parent models.Module
			if err = db.Where("id = ? AND deleted_at IS NULL", parsedParentID).First(&parent).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, errors.New("módulo padre no encontrado")
				}
				return nil, err
			}
			module.ParentID = &parsedParentID
		}
	}

	// Validar nombre único si se está actualizando (solo para módulos raíz)
	if req.Name != nil && *req.Name != module.Name {
		// Determinar si el módulo será raíz después de la actualización
		willBeRoot := module.ParentID == nil || setParentToNull
		
		if willBeRoot {
			var exists int64
			query := db.Model(&models.Module{}).
				Where("name = ? AND id != ? AND deleted_at IS NULL AND parent_id IS NULL", *req.Name, moduleID)
			
			if module.ApplicationID != nil {
				query = query.Where("application_id = ?", *module.ApplicationID)
			} else {
				query = query.Where("application_id IS NULL")
			}
			
			if countErr := query.Count(&exists).Error; countErr != nil {
				return nil, err
			}
			if exists > 0 {
				return nil, errors.New("ya existe un módulo raíz con este nombre en esta aplicación")
			}
		}
		module.Name = *req.Name
	}

	if req.Item != nil {
		module.Item = req.Item
	}
	if req.Route != nil {
		module.Route = req.Route
	}
	if req.Icon != nil {
		module.Icon = req.Icon
	}
	if req.SortOrder != nil {
		module.SortOrder = *req.SortOrder
	}
	if req.Status != nil {
		module.Status = *req.Status
	}

	module.UpdatedAt = time.Now()

	// Usar transacción para manejar el caso de parent_id = NULL
	err = db.Transaction(func(tx *gorm.DB) error {
		// Si necesitamos poner parent_id en NULL, hacerlo explícitamente
		if setParentToNull {
			if err = tx.Model(&module).Update("parent_id", nil).Error; err != nil {
				return err
			}
		}
		
		// Guardar el resto de los campos
		if err = tx.Save(&module).Error; err != nil {
			return err
		}
		
		return nil
	})

	if err != nil {
		return nil, err
	}

	return GetModuleByID(id)
}