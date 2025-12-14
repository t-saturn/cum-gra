package services

import (
	"errors"

	"server/internal/config"
	"server/internal/dto"
	"server/internal/mapper"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUserByID(id string) (*dto.UserDetailDTO, error) {
	db := config.DB

	userID, err := uuid.Parse(id)
	if err != nil {
		return nil, errors.New("ID inválido")
	}

	var user models.User
	if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("usuario no encontrado")
		}
		return nil, err
	}

	var userDetail models.UserDetail
	db.Preload("StructuralPosition").
		Preload("OrganicUnit").
		Preload("Ubigeo").
		Where("user_id = ?", userID).
		First(&userDetail)

	result := mapper.ToUserDetailDTO(user, &userDetail)
	return &result, nil
}