package services

import (
	"context"

	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/pkg/logger"
	"github.com/t-saturn/auth-service-server/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// LogVerify registra un intento de /auth/verify en la colección verify_attempts.
// - status: "success" | "failed" | "locked" ...
// - reason: "invalid_password" | "user_not_found" | "account_inactive" | ...
// - userID: vacía si no se encontró usuario
func (s *AuthService) LogVerify(ctx context.Context, input dto.AuthVerifyRequestDTO, status, reason, userID string, validationTimeMs int64) (primitive.ObjectID, error) {
	now := utils.NowUTC()

	attempt := &models.VerifyAttempt{
		Email:          deref(input.Email),
		UserID:         userID,
		ApplicationID:  input.ApplicationID,
		Status:         status,
		Reason:         reason,
		Method:         models.AuthMethodCredentials,
		CreatedAt:      now,
		ValidatedAt:    &now,
		ValidationTime: validationTimeMs,
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
		ValidationResponse: &models.ValidationResponse{
			UserID:          userID,
			ServiceResponse: status,
			ValidatedBy:     models.AuthMethodCredentials,
			ValidationTime:  validationTimeMs,
		},
	}

	return s.verifyAttemptRepo.Insert(ctx, attempt)
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
