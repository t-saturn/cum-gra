package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/security"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	logger.Log.Debugf("Iniciando verificación de credenciales")

	type UserData struct {
		ID           uuid.UUID
		Email        string
		PasswordHash string
		DNI          string
		Status       string
		IsDeleted    bool
	}

	var user UserData
	var status string = "success"
	var userID uuid.UUID

	tx := s.DB.Table("users").
		Where("is_deleted = false").
		Scopes(func(db *gorm.DB) *gorm.DB {
			if input.Email != nil && *input.Email != "" {
				logger.Log.Debugf("Consultando usuario por email: %s", *input.Email)
				return db.Where("email = ?", *input.Email)
			} else if input.DNI != nil && *input.DNI != "" {
				logger.Log.Debugf("Consultando usuario por DNI: %s", *input.DNI)
				return db.Where("dni = ?", *input.DNI)
			}
			logger.Log.Warn("No se proporcionó ni email ni DNI")
			status = "missing_identifier"
			return db
		}).
		First(&user)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			status = "invalid_credentials"
		} else {
			logger.Log.Errorf("Error consultando usuario: %v", tx.Error)
			status = "db_error"
		}
	} else {
		userID = user.ID
		if user.Status != "active" {
			logger.Log.Warnf("Usuario con estado inactivo: %s", user.Status)
			status = "inactive_account"
		} else {
			argon := security.NewArgon2Service()
			if !argon.CheckPasswordHash(input.Password, user.PasswordHash) {
				logger.Log.Warn("Contraseña inválida")
				status = "invalid_credentials"
			} else {
				logger.Log.Debug("Contraseña verificada correctamente")
				status = "success"
			}
		}
	}

	now := time.Now()
	userObjectID := primitive.NewObjectID()

	authAttempt := models.AuthAttempt{
		Method:        "credentials",
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
			UserID:         userObjectID,
			ValidatedBy:    "credentials",
			ValidationTime: 0,
		},
	}

	authCol := config.GetMongoCollection("auth_attempts")
	_, err := authCol.InsertOne(context.TODO(), authAttempt)
	if err != nil {
		logger.Log.Errorf("Error insertando AuthAttempt: %v", err)
		return nil, err
	}

	if status != "success" {
		if status == "invalid_credentials" {
			return nil, ErrInvalidCredentials
		} else if status == "inactive_account" {
			return nil, ErrInactiveAccount
		}
		return nil, errors.New("fallo de autenticación")
	}

	accessToken, err := security.GenerateToken(userID.String(), 15*time.Minute)
	if err != nil {
		logger.Log.Errorf("Error generando access token: %v", err)
		return nil, err
	}
	refreshToken, err := security.GenerateToken(userID.String(), 7*24*time.Hour)
	if err != nil {
		logger.Log.Errorf("Error generando refresh token: %v", err)
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
