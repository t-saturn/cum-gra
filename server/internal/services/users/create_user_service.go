package services

import (
	"errors"
	"strconv"
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUser(req dto.CreateUserRequest, createdBy uuid.UUID) (*dto.UserDetailDTO, error) {
	db := config.DB

	// Verificar email único
	var exists int64
	if err := db.Model(&models.User{}).
		Where("email = ? AND is_deleted = FALSE", req.Email).
		Count(&exists).Error; err != nil {
		return nil, err
	}
	if exists > 0 {
		return nil, errors.New("ya existe un usuario con este email")
	}

	// Verificar DNI único
	if err := db.Model(&models.User{}).
		Where("dni = ? AND is_deleted = FALSE", req.DNI).
		Count(&exists).Error; err != nil {
		return nil, err
	}
	if exists > 0 {
		return nil, errors.New("ya existe un usuario con este DNI")
	}

	status := "active"
	if req.Status != nil {
		status = *req.Status
	}

	// Validar structural position si se proporciona
	var structuralPositionID *uint
	if req.StructuralPositionID != nil && *req.StructuralPositionID != "" {
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
		structuralPositionID = &pid
	}

	// Validar organic unit si se proporciona
	var organicUnitID *uint
	if req.OrganicUnitID != nil && *req.OrganicUnitID != "" {
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
		organicUnitID = &oid
	}

	// Validar ubigeo si se proporciona
	var ubigeoID *uint
	if req.UbigeoID != nil && *req.UbigeoID != "" {
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
		ubigeoID = &uid
	}

	// Crear usuario
	user := models.User{
		ID:        uuid.New(),
		Email:     req.Email,
		DNI:       req.DNI,
		Status:    status,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsDeleted: false,
	}

	if err := db.Create(&user).Error; err != nil {
		return nil, err
	}

	// Crear user detail
	userDetail := models.UserDetail{
		UserID:               user.ID,
		FirstName:            &req.FirstName,
		LastName:             &req.LastName,
		Phone:                req.Phone,
		CodEmpSGD:            req.CodEmpSGD,
		StructuralPositionID: structuralPositionID,
		OrganicUnitID:        organicUnitID,
		UbigeoID:             ubigeoID,
	}

	if err := db.Create(&userDetail).Error; err != nil {
		return nil, err
	}

	// Cargar relaciones
	db.Preload("StructuralPosition").
		Preload("OrganicUnit").
		Preload("Ubigeo").
		First(&userDetail, userDetail.ID)

	result := mapper.ToUserDetailDTO(user, &userDetail)
	return &result, nil
}
