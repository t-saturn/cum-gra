package services

import (
	"server/internal/dto"
	"server/internal/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SessionMeService interface {
	Execute(userID string) (*dto.SessionMeResponse, error)
}

type sessionMeServiceImpl struct {
	db *gorm.DB
}

func NewSessionMeService(db *gorm.DB) SessionMeService {
	return &sessionMeServiceImpl{db: db}
}

func (s *sessionMeServiceImpl) Execute(userID string) (*dto.SessionMeResponse, error) {
	var user models.User
	if err := s.db.First(&user, "id = ?", uuid.MustParse(userID)).Error; err != nil {
		return nil, err
	}

	return &dto.SessionMeResponse{
		ID:        user.ID.String(),
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Phone:     user.Phone,
		DNI:       user.DNI,
		Status:    user.Status,
		IsDeleted: user.IsDeleted,
	}, nil
}
