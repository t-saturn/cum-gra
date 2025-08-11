package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/pkg/security"
	"github.com/t-saturn/auth-service-server/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// InsertVerify registra un intento de /auth/verify en la colección verify_attempts.
// - status: "success" | "failed" | "locked" ...
// - reason: "invalid_password" | "user_not_found" | "account_inactive" | ...
// - userID: vacía si no se encontró usuario
func (s *AuthService) InsertVerify(ctx context.Context, input dto.AuthVerifyRequestDTO, status, reason, userID string, validationTimeMs int64) (primitive.ObjectID, error) {
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

// InsertAttempt registra un AuthAttempt con status y reason.
func (s *AuthService) InsertAttempt(ctx context.Context, input dto.AuthLoginRequestDTO, status, reason, userID string) (primitive.ObjectID, error) {
	now := utils.NowUTC()
	attempt := &models.AuthAttempt{
		Method:        models.AuthMethodCredentials,
		Status:        status,
		ApplicationID: input.ApplicationID,
		Email:         input.Email,
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
			ServiceResponse: reason,
			ValidatedBy:     models.AuthMethodCredentials,
			ValidationTime:  0,
		},
	}
	return s.authAttemptRepo.Insert(ctx, attempt)
}

// InsertSession encapsula la creación de una Session y retorna su ObjectID y DTO.
func (s *AuthService) InsertSession(ctx context.Context, input dto.AuthLoginRequestDTO, userID string, authAttemptID primitive.ObjectID, now time.Time) (primitive.ObjectID, dto.SessionDTO, error) {
	// Calcular expiración
	expires := now.Add(24 * time.Hour)

	// Calcular duración máxima de sesión (refresh)
	maxRefresh := now.Add(7 * 24 * time.Hour) // 7 días sin reautenticarse

	sess := &models.Session{
		SessionID:    uuid.New().String(),
		UserID:       userID,
		Status:       models.SessionStatusActive,
		IsActive:     true,
		CreatedAt:    now,
		LastActivity: now,
		ExpiresAt:    expires,
		MaxRefreshAt: maxRefresh,
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
		AuthAttemptID: &authAttemptID,
	}

	// Persistir en Mongo y obtener ObjectID
	objID, err := s.sessionRepo.Insert(ctx, sess)
	if err != nil {
		return primitive.NilObjectID, dto.SessionDTO{}, err
	}

	// Construir DTO para la respuesta
	sessDTO := dto.SessionDTO{
		SessionID: sess.SessionID,
		Status:    sess.Status,
		CreatedAt: sess.CreatedAt,
		ExpiresAt: sess.ExpiresAt,
	}

	return objID, sessDTO, nil
}

// InsertToken encapsula la creación y persistencia de un token.
// - recibe sólo lo mínimo necesario: userID, sessionID, deviceInfo, tipo y duración.
// - devuelve el ObjectID recién insertado.

func (s *AuthService) InsertToken(ctx context.Context, userID string, sessionID string, deviceInfo dto.DeviceInfoDTO, tokenType string, duration time.Duration, parentID *primitive.ObjectID) (primitive.ObjectID, string, error) {
	// 1) Generar el JWT
	var jwtStr string
	var err error
	switch tokenType {
	case models.TokenTypeAccess:
		jwtStr, err = security.GenerateAccessToken(userID)
	case models.TokenTypeRefresh:
		jwtStr, err = security.GenerateRefreshToken(userID)
	default:
		return primitive.NilObjectID, "", errors.New("tipo de token no soportado")
	}
	if err != nil {
		return primitive.NilObjectID, "", err
	}

	// 2) Calcular hash del token crudo
	tokenHash := security.HashTokenHex(jwtStr) // asegúrate de usar el mismo algoritmo siempre

	now := time.Now().UTC()
	tokenUUID := uuid.New().String()

	tokModel := &models.Token{
		TokenID:       tokenUUID,
		TokenHash:     tokenHash, // <-- GUARDAMOS EL HASH, NO EL TOKEN
		UserID:        userID,
		SessionID:     sessionID,
		Status:        models.TokenStatusActive,
		TokenType:     tokenType,
		IssuedAt:      now,
		ExpiresAt:     now.Add(duration),
		CreatedAt:     now,
		UpdatedAt:     now,
		DeviceInfo:    deviceInfo.ToModel(),
		ParentTokenID: parentID,
		Alg:           "RS256", // opcional
		// Kid:        "current-key-id",   // opcional, si usas kid en el header
	}

	oid, err := s.tokenRepo.Insert(ctx, tokModel)
	return oid, jwtStr, err // retornas el token crudo al cliente
}

// deref convierte *string a string, devolviendo cadena vacía si es nil.
func deref(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}
