package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/pkg/security"
	"gorm.io/gorm"
)

var (
	ErrInvalidCredentials = errors.New("credenciales inválidas")
	ErrInactiveAccount    = errors.New("cuenta inactiva")
)

type AuthService struct {
	DB *gorm.DB
}

type AuthResult struct {
	UserID       string
	AccessToken  string
	RefreshToken string
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{DB: db}
}

func (s *AuthService) VerifyCredentials(input dto.AuthVerifyRequest) (*AuthResult, error) {
	type UserData struct {
		ID           uuid.UUID
		Email        string
		PasswordHash string
		DNI          string
		Status       string
		IsDeleted    bool
	}

	var user UserData

	tx := s.DB.Table("users").
		Where("is_deleted = ?", false).
		Scopes(func(db *gorm.DB) *gorm.DB {
			if input.Email != nil && *input.Email != "" {
				return db.Where("email = ?", *input.Email)
			} else if input.DNI != nil && *input.DNI != "" {
				return db.Where("dni = ?", *input.DNI)
			}
			return db
		}).
		First(&user)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, ErrInvalidCredentials
		}
		return nil, tx.Error
	}

	if user.Status != "active" {
		return nil, ErrInactiveAccount
	}

	argon := security.NewArgon2Service()
	if !argon.CheckPasswordHash(input.Password, user.PasswordHash) {
		return nil, ErrInvalidCredentials
	}

	accessToken, err := security.GenerateToken(user.ID.String(), 15*time.Minute)
	if err != nil {
		return nil, err
	}

	refreshToken, err := security.GenerateToken(user.ID.String(), 7*24*time.Hour)
	if err != nil {
		return nil, err
	}

	return &AuthResult{
		UserID:       user.ID.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
