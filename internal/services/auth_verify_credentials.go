package services

import (
	"context"
	"errors"

	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/internal/repositories"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/security"
	"github.com/t-saturn/auth-service-server/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	ErrInvalidCredentials = errors.New("credenciales inválidas")
	ErrInactiveAccount    = errors.New("cuenta inactiva")
)

// AuthService gestiona la lógica de autenticación.
type AuthService struct {
	userRepo        *repositories.UserRepository
	authAttemptRepo *repositories.AuthAttemptRepository
}

// NewAuthService construye un AuthService con Postgres y Mongo ya conectados.
func NewAuthService(pgDB *gorm.DB, mongoDB *mongo.Database) *AuthService {
	return &AuthService{
		userRepo:        repositories.NewUserRepository(pgDB),
		authAttemptRepo: repositories.NewAuthAttemptRepository(mongoDB),
	}
}

// VerifyCredentials verifica email/DNI + contraseña y retorna los tokens.
func (s *AuthService) VerifyCredentials(ctx context.Context, input dto.AuthVerifyRequestDTO) (*dto.AuthVerifyResponseDTO, error) {
	// 1 Cargar usuario
	userData, err := s.userRepo.FindActiveByEmailOrDNI(ctx, input.Email, input.DNI)
	if err != nil {
		switch {
		case errors.Is(err, repositories.ErrUserDeleted), errors.Is(err, repositories.ErrUserDisabled):
			// Cuenta eliminada o deshabilitada
			s.LogAttempt(ctx, input, models.AuthStatusFailed, "")
			return nil, ErrInactiveAccount
		case errors.Is(err, gorm.ErrRecordNotFound), errors.Is(err, repositories.ErrUserNotFound):
			// Usuario no existe
			s.LogAttempt(ctx, input, models.AuthStatusInvalid, "")
			return nil, ErrInvalidCredentials
		default:
			return nil, err
		}
	}

	// 2 Verificar la contraseña
	argon := security.NewArgon2Service()
	if !argon.CheckPasswordHash(input.Password, userData.PasswordHash) {
		s.LogAttempt(ctx, input, models.AuthStatusInvalid, "")
		return nil, ErrInvalidCredentials
	}

	// 3 Generar tokens JWE
	accessToken, err := security.GenerateAccessToken(userData.ID.String())
	if err != nil {
		return nil, err
	}
	refreshToken, err := security.GenerateRefreshToken(userData.ID.String())
	if err != nil {
		return nil, err
	}

	// 4 Registrar intento exitoso
	s.LogAttempt(ctx, input, models.AuthStatusSuccess, userData.ID.String())

	// 5 Devolver respuesta DTO
	return &dto.AuthVerifyResponseDTO{
		UserID:       userData.ID.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

// logAttempt inserta un AuthAttempt usando el repositorio.
// status debe ser uno de los modelos.AuthStatus* y userID solo si es éxito.
func (s *AuthService) LogAttempt(ctx context.Context, input dto.AuthVerifyRequestDTO, status, userID string) {
	now := utils.NowUTC()
	attempt := &models.AuthAttempt{
		Method:        models.AuthMethodCredentials,
		Status:        status,
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
			UserID:          userID,
			ServiceResponse: status,
			ValidatedBy:     models.AuthMethodCredentials,
			ValidationTime:  0,
		},
	}

	if err := s.authAttemptRepo.Insert(ctx, attempt); err != nil {
		logger.Log.Errorf("Error guardando AuthAttempt: %v", err)
	}
}

// deref convierte *string a string, devolviendo cadena vacía si es nil.
func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
