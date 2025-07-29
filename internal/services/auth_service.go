package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
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
	status := models.AuthStatusSuccess
	var userID uuid.UUID
	tx := s.DB.Table("users").
		Where("is_deleted = false").
		Scopes(func(db *gorm.DB) *gorm.DB {
			if input.Email != nil && *input.Email != "" {
				return db.Where("email = ?", *input.Email)
			} else if input.DNI != nil && *input.DNI != "" {
				return db.Where("dni = ?", *input.DNI)
			}
			status = models.AuthStatusFailed
			return db
		}).
		First(&user)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			status = models.AuthStatusInvalid
		} else {
			status = models.AuthStatusFailed
		}
	} else {
		userID = user.ID
		if user.Status != "active" {
			status = models.AuthStatusFailed
		} else {
			argon := security.NewArgon2Service()
			if !argon.CheckPasswordHash(input.Password, user.PasswordHash) {
				status = models.AuthStatusInvalid
			} else {
				status = models.AuthStatusSuccess
			}
		}
	}

	now := time.Now()
	authAttempt := models.AuthAttempt{
		Method:        models.AuthMethodCredentials,
		Status:        status,
		ApplicationID: input.ApplicationID,
		Email:         deref(input.Email),
		DeviceInfo: models.DeviceInfo{
			UserAgent:   input.DeviceInfo.UserAgent,
			IP:          input.DeviceInfo.IP,
			DeviceID:    input.DeviceInfo.DeviceID,
			OS:          input.DeviceInfo.OS,
			BrowserName: input.DeviceInfo.BrowserName,
		},
		CreatedAt:   now,
		ValidatedAt: &now,
		ValidationResponse: &models.ValidationResponse{
			UserID:          "",
			ServiceResponse: status,
			ValidatedBy:     models.AuthMethodCredentials,
			ValidationTime:  0,
		},
	}

	if status == models.AuthStatusSuccess {
		authAttempt.ValidationResponse.UserID = userID.String()
	}
	authCol := config.GetMongoCollection("auth_attempts")
	_, err := authCol.InsertOne(context.TODO(), authAttempt)
	if err != nil {
		return nil, err
	}

	if status != models.AuthStatusSuccess {
		switch status {
		case models.AuthStatusInvalid:
			return nil, ErrInvalidCredentials
		case models.AuthStatusFailed:
			return nil, ErrInactiveAccount
		}
		return nil, errors.New("fallo de autenticación")
	}

	accessToken, err := security.GenerateAccessToken(userID.String())
	if err != nil {
		return nil, err
	}
	refreshToken, err := security.GenerateRefreshToken(userID.String())
	if err != nil {
		return nil, err
	}

	return &AuthResult{
		UserID:       userID.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
