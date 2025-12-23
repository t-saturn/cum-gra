package services

import (
	"errors"
	"fmt"
	"strconv"
	"time"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"
	keycloakService "server/internal/services/keycloak"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func CreateUser(req dto.CreateUserRequest, createdBy uuid.UUID, accessToken string) (*dto.UserDetailDTO, error) {
	db := config.DB

	// 1. Verificar email único en BD local
	var exists int64
	if err := db.Model(&models.User{}).
		Where("email = ? AND is_deleted = FALSE", req.Email).
		Count(&exists).Error; err != nil {
		return nil, err
	}
	if exists > 0 {
		return nil, errors.New("ya existe un usuario con este email")
	}

	// 2. Verificar DNI único en BD local
	if err := db.Model(&models.User{}).
		Where("dni = ? AND is_deleted = FALSE", req.DNI).
		Count(&exists).Error; err != nil {
		return nil, err
	}
	if exists > 0 {
		return nil, errors.New("ya existe un usuario con este DNI")
	}

	// 3. Verificar si existe en Keycloak
	existsInKC, existingKCUserID, err := keycloakService.UserExistsInKeycloak(accessToken, req.Email, req.DNI)
	if err != nil {
		return nil, fmt.Errorf("error verificando usuario en keycloak: %w", err)
	}

	var keycloakUserID string

	if existsInKC {
		// Usuario ya existe en Keycloak, usar ese UUID
		keycloakUserID = existingKCUserID
	} else {
		// 4. Crear usuario en Keycloak primero
		kcInput := keycloakService.CreateKeycloakUserInput{
			Email:     req.Email,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			DNI:       req.DNI,
			Password:  req.Password,
		}

		kcResult, errKC := keycloakService.CreateKeycloakUser(accessToken, kcInput)

		if errKC != nil {
			return nil, fmt.Errorf("error creando usuario en keycloak: %w", err)
		}

		keycloakUserID = kcResult.UserID
	}

	// 5. Parsear UUID de Keycloak
	userUUID, err := uuid.Parse(keycloakUserID)
	if err != nil {
		return nil, fmt.Errorf("uuid inválido de keycloak: %w", err)
	}

	// 6. Validaciones de relaciones
	status := "active"
	if req.Status != nil {
		status = *req.Status
	}

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

	// 7. Crear usuario en BD local con UUID de Keycloak
	user := models.User{
		ID:        userUUID, // UUID de Keycloak
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

	// 8. Crear user detail
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

	// 9. Cargar relaciones
	db.Preload("StructuralPosition").
		Preload("OrganicUnit").
		Preload("Ubigeo").
		First(&userDetail, userDetail.ID)

	userDTO := mapper.ToUserDetailDTO(user, &userDetail)

	return &userDTO, nil
}