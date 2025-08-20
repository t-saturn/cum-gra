package services

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/internal/repositories"
	"github.com/t-saturn/auth-service-server/pkg/security"
)

// Me: GET /auth/me
// Me valida el access token (desde Authorization) y devuelve datos del usuario + sesión actual.
func (s *AuthService) Me(ctx context.Context, accessToken string, input dto.AuthMeQueryDTO) (*dto.AuthMeResponseDTO, error) {
	// 0. Validación mínima
	if accessToken == "" || input.SessionID == "" {
		return nil, ErrInvalidToken
	}

	// 1. Lookup rápido por hash del token crudo
	hash := security.HashTokenHex(accessToken)
	tokModel, err := s.tokenRepo.FindByHash(ctx, hash)
	if err != nil {
		return nil, ErrInvalidToken
	}

	// 2. Debe estar activo
	if tokModel.Status != models.TokenStatusActive {
		return nil, ErrInvalidToken
	}

	// 2.5. Expiración por DB
	now := time.Now().UTC()
	if now.After(tokModel.ExpiresAt) {
		_ = s.tokenRepo.MarkExpired(ctx, tokModel.ID, now)
		return nil, security.ErrTokenExpired
	}

	// 3. La sesión del token debe coincidir con la solicitada
	if tokModel.SessionID != input.SessionID {
		return nil, ErrSessionMismatch
	}

	// 4. Cargar sesión
	sessModel, err := s.sessionRepo.FindBySessionID(ctx, input.SessionID)
	if err != nil || sessModel == nil {
		return nil, ErrSessionNotFound
	}

	// 5. Verificar estado de sesión
	if sessModel.Status != models.SessionStatusActive || !sessModel.IsActive {
		return nil, ErrSessionInactive
	}

	// 6. Verificar firma JWS RS256 (y exp del JWS primero)
	claims, vErr := security.VerifyTokenRS256(accessToken)
	if vErr != nil {
		if errors.Is(vErr, security.ErrTokenExpired) {
			if tokModel.Status == models.TokenStatusActive {
				_ = s.tokenRepo.MarkExpired(ctx, tokModel.ID, time.Now().UTC())
			}
			return nil, security.ErrTokenExpired
		}
		return nil, security.ErrTokenInvalid
	}

	// 6.1. Defensa adicional: subject del JWS debe coincidir con el user del token en DB (si viene)
	if claims.Subject != "" && claims.Subject != tokModel.UserID {
		return nil, ErrInvalidToken
	}

	// 7. Cargar usuario (Postgres) con nombres de posición/unidad
	userUUID, err := uuid.Parse(sessModel.UserID)
	if err != nil {
		return nil, repositories.ErrUserNotFound
	}
	uview, err := s.userRepo.FindActiveByIDWithOrgNames(ctx, userUUID)
	if err != nil {
		return nil, err
	}

	// 8. Mapear respuesta
	resp := &dto.AuthMeResponseDTO{
		UserID:             uview.ID.String(),
		Email:              uview.Email,
		Name:               uview.FirstName + " " + uview.LastName,
		DNI:                uview.DNI,
		Status:             uview.Status,
		StructuralPosition: derefStr(uview.StructuralPositionName),
		OrganicUnit:        derefStr(uview.OrganicUnitName),
		Session:            toSessionViewDTO(sessModel),
		Role:               "admin",
		ModulePermissions:  []string{"module1", "module2", "module3"},
		ModuleRestriccions: []string{"module1"},
	}
	if uview.Phone != nil {
		resp.Phone = *uview.Phone
	}

	return resp, nil
}

// Mapeadores a DTO
func toSessionViewDTO(s *models.Session) dto.SessionViewDTO {
	ids := make([]string, 0, len(s.TokensGenerated))
	for _, id := range s.TokensGenerated {
		ids = append(ids, id.Hex())
	}

	return dto.SessionViewDTO{
		SessionID:       s.SessionID,
		UserID:          s.UserID,
		Status:          s.Status,
		IsActive:        s.IsActive,
		MaxRefreshAt:    s.MaxRefreshAt,
		CreatedAt:       s.CreatedAt,
		LastActivity:    s.LastActivity,
		ExpiresAt:       s.ExpiresAt,
		RevokedAt:       s.RevokedAt,
		DeviceInfo:      toDeviceInfoDTO(s.DeviceInfo), // <- Model → DTO
		TokensGenerated: ids,
	}
}

func toDeviceInfoDTO(mi models.DeviceInfo) dto.DeviceInfoDTO {
	var loc *dto.LocationDetailDTO
	if mi.Location != nil {
		loc = &dto.LocationDetailDTO{
			Country:     mi.Location.Country,
			CountryCode: mi.Location.CountryCode,
			Region:      mi.Location.Region,
			City:        mi.Location.City,
			Coordinates: dto.CoordinatesDTO{
				mi.Location.Coordinates[0],
				mi.Location.Coordinates[1],
			},
			ISP:          mi.Location.ISP,
			Organization: mi.Location.Organization,
		}
	}

	return dto.DeviceInfoDTO{
		UserAgent:      mi.UserAgent,
		IP:             mi.IP,
		DeviceID:       mi.DeviceID,
		BrowserName:    mi.BrowserName,
		BrowserVersion: mi.BrowserVersion,
		OS:             mi.OS,
		OSVersion:      mi.OSVersion,
		DeviceType:     mi.DeviceType,
		Timezone:       mi.Timezone,
		Language:       mi.Language,
		Location:       loc,
	}
}

func derefStr(s *string) string {
	if s != nil {
		return *s
	}
	return ""
}
