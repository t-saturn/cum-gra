package services

import (
	"context"
	"errors"

	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	repository "github.com/t-saturn/auth-service-server/internal/repositories"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/security"
	"github.com/t-saturn/auth-service-server/pkg/utils"
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

func (s *AuthService) VerifyCredentials(input dto.AuthVerifyRequestDTO) (*AuthResult, error) {
	// 1) Buscar usuario usando el repositorio
	userRepo := repository.NewUserRepository(s.DB)
	userData, err := userRepo.FindActiveByEmailOrDNI(context.Background(), input.Email, input.DNI)
	if err != nil {
		switch {
		case errors.Is(err, repository.ErrUserDeleted):
			// El repo detectó is_deleted = true
			return nil, ErrInactiveAccount
		case errors.Is(err, repository.ErrUserDisabled):
			// El repo detectó status != "active"
			return nil, ErrInactiveAccount
		case errors.Is(err, gorm.ErrRecordNotFound):
			// No existe ningún usuario con ese email/DNI
			return nil, ErrInvalidCredentials
		default:
			return nil, err
		}
	}

	// 2) Verificar la contraseña
	argon := security.NewArgon2Service()
	if !argon.CheckPasswordHash(input.Password, userData.PasswordHash) {
		return nil, ErrInvalidCredentials
	}

	// 3) Generar tokens JWE
	accessToken, err := security.GenerateAccessToken(userData.ID.String())
	if err != nil {
		return nil, err
	}
	refreshToken, err := security.GenerateRefreshToken(userData.ID.String())
	if err != nil {
		return nil, err
	}

	// 4) Registrar el intento en MongoDB
	now := utils.NowUTC()

	authAttempt := models.AuthAttempt{
		Method:        models.AuthMethodCredentials,
		Status:        models.AuthStatusSuccess,
		ApplicationID: input.ApplicationID,
		Email:         deref(input.Email),
		DeviceInfo: models.DeviceInfo{
			UserAgent:      input.DeviceInfo.UserAgent,
			IP:             input.DeviceInfo.IP,
			DeviceID:       input.DeviceInfo.DeviceID,
			OS:             input.DeviceInfo.OS,
			OSVersion:      input.DeviceInfo.OSVersion,
			BrowserName:    input.DeviceInfo.BrowserName,
			BrowserVersion: input.DeviceInfo.BrowserVersion,
			DeviceType:     input.DeviceInfo.DeviceType,
			Timezone:       input.DeviceInfo.Timezone,
			Language:       input.DeviceInfo.Language,
			Location: &models.LocationDetail{
				Country:     input.DeviceInfo.Location.Country,
				CountryCode: input.DeviceInfo.Location.CountryCode,
				Region:      input.DeviceInfo.Location.Region,
				City:        input.DeviceInfo.Location.City,
				Coordinates: models.Coordinates{
					input.DeviceInfo.Location.Coordinates[0],
					input.DeviceInfo.Location.Coordinates[1]},
				ISP:          input.DeviceInfo.Location.ISP,
				Organization: input.DeviceInfo.Location.Organization,
			},
		},
		CreatedAt:   now,
		ValidatedAt: &now,
		ValidationResponse: &models.ValidationResponse{
			UserID:          userData.ID.String(),
			ServiceResponse: models.AuthStatusSuccess,
			ValidatedBy:     models.AuthMethodCredentials,
			ValidationTime:  0,
		},
	}
	authCol := config.GetMongoCollection("auth_attempts")
	if _, err := authCol.InsertOne(context.Background(), authAttempt); err != nil {
		logger.Log.Errorf("Error guardando auth attempt en Mongo: %v", err)
	}

	// 5) Devolver resultado exitoso
	return &AuthResult{
		UserID:       userData.ID.String(),
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
