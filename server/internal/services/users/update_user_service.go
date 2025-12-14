package services

import (
	"errors"
	"strconv"
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func UpdateUser(id string, req dto.UpdateUserRequest, updatedBy uuid.UUID) (*dto.UserDetailDTO, error) {
	db := config.DB

	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var user models.User
	if err := db.Where("id = ? AND is_deleted = FALSE", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("usuario no encontrado")
		}
		return nil, err
	}

	// Validar email único si se está actualizando
	if req.Email != nil && *req.Email != user.Email {
		var exists int64
		if err := db.Model(&models.User{}).
			Where("email = ? AND id != ? AND is_deleted = FALSE", *req.Email, userID).
			Count(&exists).Error; err != nil {
			return nil, err
		}
		if exists > 0 {
			return nil, errors.New("ya existe un usuario con este email")
		}
		user.Email = *req.Email
	}

	// Validar DNI único si se está actualizando
	if req.DNI != nil && *req.DNI != user.DNI {
		var exists int64
		if err := db.Model(&models.User{}).
			Where("dni = ? AND id != ? AND is_deleted = FALSE", *req.DNI, userID).
			Count(&exists).Error; err != nil {
			return nil, err
		}
		if exists > 0 {
			return nil, errors.New("ya existe un usuario con este DNI")
		}
		user.DNI = *req.DNI
	}

	if req.Status != nil {
		user.Status = *req.Status
	}

	user.UpdatedAt = time.Now()

	if err := db.Save(&user).Error; err != nil {
		return nil, err
	}

	// Actualizar user detail
	var userDetail models.UserDetail
	if err := db.Where("user_id = ?", userID).First(&userDetail).Error; err != nil {
		// Si no existe, crear uno nuevo
		if errors.Is(err, gorm.ErrRecordNotFound) {
			userDetail = models.UserDetail{
				UserID: userID,
			}
			if errDB := db.Create(&userDetail).Error; errDB != nil {
				return nil, errDB
			}
		} else {
			return nil, err
		}
	}

	if req.FirstName != nil {
		userDetail.FirstName = req.FirstName
	}
	if req.LastName != nil {
		userDetail.LastName = req.LastName
	}
	if req.Phone != nil {
		userDetail.Phone = req.Phone
	}
	if req.CodEmpSGD != nil {
		userDetail.CodEmpSGD = req.CodEmpSGD
	}

	// Validar structural position si se está actualizando
	if req.StructuralPositionID != nil {
		if *req.StructuralPositionID == "" {
			userDetail.StructuralPositionID = nil
		} else {
			posID, err := strconv.ParseUint(*req.StructuralPositionID, 10, 32)
			if err != nil {
				return nil, errors.New("structural_position_id inválido")
			}

			var position models.StructuralPosition
			if err := db.Where("id = ? AND is_deleted = FALSE", uint(posID)).First(&position).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, errors.New("posición estructural no encontrada")
				}
				return nil, err
			}
			pid := uint(posID)
			userDetail.StructuralPositionID = &pid
		}
	}

	// Validar organic unit si se está actualizando
	if req.OrganicUnitID != nil {
		if *req.OrganicUnitID == "" {
			userDetail.OrganicUnitID = nil
		} else {
			ouID, err := strconv.ParseUint(*req.OrganicUnitID, 10, 32)
			if err != nil {
				return nil, errors.New("organic_unit_id inválido")
			}

			var organicUnit models.OrganicUnit
			if err := db.Where("id = ? AND is_deleted = FALSE", uint(ouID)).First(&organicUnit).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, errors.New("unidad orgánica no encontrada")
				}
				return nil, err
			}
			oid := uint(ouID)
			userDetail.OrganicUnitID = &oid
		}
	}

	// Validar ubigeo si se está actualizando
	if req.UbigeoID != nil {
		if *req.UbigeoID == "" {
			userDetail.UbigeoID = nil
		} else {
			ubID, err := strconv.ParseUint(*req.UbigeoID, 10, 32)
			if err != nil {
				return nil, errors.New("ubigeo_id inválido")
			}

			var ubigeo models.Ubigeo
			if err := db.Where("id = ?", uint(ubID)).First(&ubigeo).Error; err != nil {
				if errors.Is(err, gorm.ErrRecordNotFound) {
					return nil, errors.New("ubigeo no encontrado")
				}
				return nil, err
			}
			uid := uint(ubID)
			userDetail.UbigeoID = &uid
		}
	}

	if err := db.Save(&userDetail).Error; err != nil {
		return nil, err
	}

	return GetUserByID(id)
}