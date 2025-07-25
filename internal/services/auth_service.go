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
			return db
		}).
		First(&user)

	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			logger.Log.Warn("Usuario no encontrado")
			return nil, ErrInvalidCredentials
		}
		logger.Log.Errorf("Error consultando usuario: %v", tx.Error)
		return nil, tx.Error
	}

	logger.Log.Debugf("Usuario encontrado: %s", user.ID)

	if user.Status != "active" {
		logger.Log.Warnf("Usuario con estado inactivo: %s", user.Status)
		return nil, ErrInactiveAccount
	}

	argon := security.NewArgon2Service()
	if !argon.CheckPasswordHash(input.Password, user.PasswordHash) {
		logger.Log.Warn("Contraseña inválida")
		return nil, ErrInvalidCredentials
	}
	logger.Log.Debug("Contraseña verificada correctamente")

	accessToken, err := security.GenerateToken(user.ID.String(), 15*time.Minute)
	if err != nil {
		logger.Log.Errorf("Error generando access token: %v", err)
		return nil, err
	}
	refreshToken, err := security.GenerateToken(user.ID.String(), 7*24*time.Hour)
	if err != nil {
		logger.Log.Errorf("Error generando refresh token: %v", err)
		return nil, err
	}
	logger.Log.Debug("Tokens generados exitosamente")

	now := time.Now()
	userObjectID := primitive.NewObjectID()

	authAttempt := models.AuthAttempt{
		Method:        "credentials",
		Status:        "success",
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

	logger.Log.Debug("Insertando AuthAttempt en MongoDB")
	authCol := config.GetMongoCollection("auth_attempts")
	insertRes, err := authCol.InsertOne(context.TODO(), authAttempt)
	if err != nil {
		logger.Log.Errorf("Error insertando AuthAttempt: %v", err)
		return nil, err
	}
	authAttemptID := insertRes.InsertedID.(primitive.ObjectID)
	logger.Log.Debugf("AuthAttempt creado con ID: %s", authAttemptID.Hex())

	authLog := models.AuthLog{
		UserID:        userObjectID,
		AuthAttemptID: &authAttemptID,
		Action:        "login",
		Success:       true,
		ApplicationID: input.ApplicationID,
		Timestamp:     now,
		DeviceInfo:    authAttempt.DeviceInfo,
	}

	logCol := config.GetMongoCollection("auth_logs")
	_, err = logCol.InsertOne(context.TODO(), authLog)
	if err != nil {
		logger.Log.Errorf("Error insertando AuthLog: %v", err)
		return nil, err
	}
	logger.Log.Debug("AuthLog insertado correctamente")

	return &AuthResult{
		UserID:       user.ID.String(),
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
