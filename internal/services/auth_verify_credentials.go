package services

import (
	"context"
	"errors"
	"time"

	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/internal/repositories"
	"github.com/t-saturn/auth-service-server/pkg/security"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

var (
	ErrInvalidCredentials = errors.New("credenciales inválidas")
	ErrInactiveAccount    = errors.New("cuenta inactiva")
)

// AuthService gestiona la lógica de autenticación.
type AuthService struct {
	userRepo          *repositories.UserRepository
	authAttemptRepo   *repositories.AuthAttemptRepository
	verifyAttemptRepo *repositories.VerifyAttemptRepository
}

// NewAuthService construye un AuthService con Postgres y Mongo ya conectados.
func NewAuthService(pgDB *gorm.DB, mongoDB *mongo.Database) *AuthService {
	return &AuthService{
		userRepo:          repositories.NewUserRepository(pgDB),
		authAttemptRepo:   repositories.NewAuthAttemptRepository(mongoDB),
		verifyAttemptRepo: repositories.NewVerifyAttemptRepository(mongoDB),
	}
}

// VerifyCredentials verifica email/DNI + contraseña y retorna los tokens.
func (s *AuthService) VerifyCredentials(ctx context.Context, input dto.AuthVerifyRequestDTO) (*dto.AuthVerifyResponseDTO, error) {
	start := time.Now()
	// 1 Cargar usuario
	userData, err := s.userRepo.FindActiveByEmailOrDNI(ctx, input.Email, input.DNI)
	if err != nil {
		switch {
		case errors.Is(err, repositories.ErrUserDeleted), errors.Is(err, repositories.ErrUserDisabled):
			// Cuenta eliminada o deshabilitada
			s.LogVerify(ctx, input, models.AuthStatusFailed, "account_inactive", "", 0)
			return nil, ErrInactiveAccount
		case errors.Is(err, gorm.ErrRecordNotFound), errors.Is(err, repositories.ErrUserNotFound):
			// Usuario no existe
			s.LogVerify(ctx, input, models.AuthStatusFailed, "user_not_found", "", 0)
			return nil, ErrInvalidCredentials
		default:
			return nil, err
		}
	}

	// 2 Verificar la contraseña
	argon := security.NewArgon2Service()
	if !argon.CheckPasswordHash(input.Password, userData.PasswordHash) {
		s.LogVerify(ctx, input, models.AuthStatusFailed, "invalid_password", "", 0)
		return nil, ErrInvalidCredentials
	}
	elapsed := time.Since(start).Milliseconds()

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
	s.LogVerify(ctx, input, models.AuthStatusSuccess, "correct", userData.ID.String(), elapsed)

	// 5 Devolver respuesta DTO
	return &dto.AuthVerifyResponseDTO{
		UserID:       userData.ID.String(),
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
