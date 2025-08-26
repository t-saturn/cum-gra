package services

import (
	"context"
	"errors"
	"math"
	"time"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
		logrus.WithFields(logrus.Fields{
			"where":     "Me.validate",
			"sessionID": input.SessionID,
			"token_len": len(accessToken),
		}).Warn("token o session_id vacío")
		return nil, ErrInvalidToken
	}

	if input.ClientID == "" {
		logrus.WithFields(logrus.Fields{
			"where": "Me.validate",
		}).Warn("AppID vacío")
		return nil, ErrInvalidAppID
	}

	// 1. Lookup rápido por hash del token crudo
	hash := security.HashTokenHex(accessToken)
	logrus.WithFields(logrus.Fields{
		"where":        "Me.token.findByHash",
		"hash_prefix":  hash[:12],
		"sessionID_in": input.SessionID,
		"client_id_in": input.ClientID,
	}).Debug("buscando token por hash")
	tokModel, err := s.tokenRepo.FindByHash(ctx, hash)
	if err != nil {
		logrus.WithError(err).WithFields(logrus.Fields{
			"where":       "Me.token.findByHash",
			"hash_prefix": hash[:12],
		}).Warn("token no encontrado / inválido")
		return nil, ErrInvalidToken
	}

	logrus.WithFields(logrus.Fields{
		"where":      "Me.token.loaded",
		"token_id":   tokModel.ID,
		"status":     tokModel.Status,
		"session_id": tokModel.SessionID,
		"user_id":    tokModel.UserID,
		"expires_at": tokModel.ExpiresAt,
	}).Debug("token cargado")

	// 2. Debe estar activo
	if tokModel.Status != models.TokenStatusActive {
		logrus.WithFields(logrus.Fields{
			"where":  "Me.token.status",
			"status": tokModel.Status,
		}).Warn("token no activo")
		return nil, ErrInvalidToken
	}

	// 2.5. Expiración por DB
	now := time.Now().UTC()
	if now.After(tokModel.ExpiresAt) {
		_ = s.tokenRepo.MarkExpired(ctx, tokModel.ID, now)
		logrus.WithFields(logrus.Fields{
			"where":      "Me.token.expiredByDB",
			"now":        now,
			"expires_at": tokModel.ExpiresAt,
		}).Warn("token expirado por DB")
		return nil, security.ErrTokenExpired
	}

	// 3. La sesión del token debe coincidir con la solicitada
	if tokModel.SessionID != input.SessionID {
		logrus.WithFields(logrus.Fields{
			"where":      "Me.token.sessionMismatch",
			"token_sess": tokModel.SessionID,
			"input_sess": input.SessionID,
		}).Warn("sessionID mismatch")
		return nil, ErrSessionMismatch
	}

	// 4. Cargar sesión
	logrus.WithFields(logrus.Fields{
		"where":      "Me.session.find",
		"session_id": input.SessionID,
	}).Debug("buscando sesión")
	sessModel, err := s.sessionRepo.FindBySessionID(ctx, input.SessionID)
	if err != nil || sessModel == nil {
		logrus.WithError(err).WithFields(logrus.Fields{
			"where":      "Me.session.find",
			"session_id": input.SessionID,
		}).Warn("sesión no encontrada")
		return nil, ErrSessionNotFound
	}
	logrus.WithFields(logrus.Fields{
		"where":       "Me.session.loaded",
		"status":      sessModel.Status,
		"is_active":   sessModel.IsActive,
		"user_id":     sessModel.UserID,
		"expires_at":  sessModel.ExpiresAt,
		"max_refresh": sessModel.MaxRefreshAt,
	}).Debug("sesión cargada")

	// 5. Verificar estado de sesión
	if sessModel.Status != models.SessionStatusActive || !sessModel.IsActive {
		logrus.WithFields(logrus.Fields{
			"where":     "Me.session.status",
			"status":    sessModel.Status,
			"is_active": sessModel.IsActive,
		}).Warn("sesión inactiva")
		return nil, ErrSessionInactive
	}

	// 6. Verificar firma JWS RS256 (y exp del JWS primero)
	claims, vErr := security.VerifyTokenRS256(accessToken)
	if vErr != nil {
		if errors.Is(vErr, security.ErrTokenExpired) {
			if tokModel.Status == models.TokenStatusActive {
				_ = s.tokenRepo.MarkExpired(ctx, tokModel.ID, time.Now().UTC())
			}
			logrus.WithFields(logrus.Fields{
				"where": "Me.jwt.verify",
			}).Warn("JWS expirado")
			return nil, security.ErrTokenExpired
		}
		logrus.WithError(vErr).WithFields(logrus.Fields{
			"where": "Me.jwt.verify",
		}).Warn("JWS inválido")
		return nil, security.ErrTokenInvalid
	}
	logrus.WithFields(logrus.Fields{
		"where":  "Me.jwt.claims",
		"sub":    claims.Subject,
		"aud":    claims.Audience,
		"issued": claims.IssuedAt,
	}).Debug("claims verificados")

	// 6.1. Defensa adicional: subject del JWS debe coincidir con el user del token en DB (si viene)
	if claims.Subject != "" && claims.Subject != tokModel.UserID {
		logrus.WithFields(logrus.Fields{
			"where":     "Me.jwt.subjectMismatch",
			"jwt_sub":   claims.Subject,
			"db_userid": tokModel.UserID,
		}).Warn("subject del JWT no coincide con token DB")
		return nil, ErrInvalidToken
	}

	// 7. Cargar usuario + datos por client_id (input.AppID)
	userUUID, err := uuid.Parse(sessModel.UserID)
	if err != nil {
		logrus.WithError(err).WithFields(logrus.Fields{
			"where":   "Me.user.parse",
			"user_id": sessModel.UserID,
		}).Warn("userID inválido en sesión")
		return nil, repositories.ErrUserNotFound
	}

	logrus.WithFields(logrus.Fields{
		"where":     "Me.user.findData",
		"user_uuid": userUUID,
		"client_id": input.ClientID, // aquí tu client_id
	}).Debug("buscando datos de usuario + rol/permisos/restricciones")
	uview, err := s.userRepo.FindDataUser(ctx, userUUID, input.ClientID)
	if err != nil {
		logrus.WithError(err).WithFields(logrus.Fields{
			"where":     "Me.user.findData",
			"user_uuid": userUUID,
			"client_id": input.ClientID,
		}).Error("falló FindDataUser")
		return nil, err
	}

	accessAt, refreshAt, err := s.sessionRepo.GetTokenExpiriesBySessionID(ctx, input.SessionID)
	if err != nil {
		logrus.WithError(err).WithFields(logrus.Fields{
			"where":      "Me.session.expiries",
			"session_id": input.SessionID,
		}).Warn("no se pudo obtener expiraciones")
		return nil, ErrSessionNotFound
	}

	const isoMillis = "2006-01-02T15:04:05.000Z07:00"
	resp := &dto.AuthMeResponseDTO{
		UserID:             uview.ID.String(),
		Email:              uview.Email,
		Name:               uview.FirstName + " " + uview.LastName,
		DNI:                uview.DNI,
		Status:             uview.Status,
		StructuralPosition: derefStr(uview.StructuralPositionName),
		OrganicUnit:        derefStr(uview.OrganicUnitName),
		Session:            toSessionViewDTO(sessModel),
		Role:               uview.Role,
		ModulePermissions:  uview.ModulePermissions,
		ModuleRestriccions: uview.ModuleRestriccions,
		AccessExpiresAt:    accessAt.Format(isoMillis),
		RefreshExpiresAt:   refreshAt.Format(isoMillis),
		Exp:                accessAt.Unix(),
		RemainingSeconds:   int64(math.Max(0, accessAt.Sub(now).Seconds())),
	}
	if uview.Phone != nil {
		resp.Phone = *uview.Phone
	}

	logrus.WithFields(logrus.Fields{
		"where":              "Me.response",
		"user_id":            resp.UserID,
		"client_id":          input.ClientID,
		"role":               resp.Role,
		"perm_count":         len(resp.ModulePermissions),
		"restr_count":        len(resp.ModuleRestriccions),
		"access_expires_at":  resp.AccessExpiresAt,
		"refresh_expires_at": resp.RefreshExpiresAt,
	}).Info("AuthMe OK")

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
