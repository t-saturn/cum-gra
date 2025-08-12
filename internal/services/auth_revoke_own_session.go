package services

import (
	"context"
	"errors"
	"time"

	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/pkg/security"
)

var (
	ErrForbiddenRevoke = errors.New("forbidden_revoke") // intentar revocar una sesión de otro usuario
)

// RevokeOwnSession valida token+sesión y revoca la sesión objetivo (más todos sus tokens activos).
func (s *AuthService) RevokeOwnSession(
	ctx context.Context,
	auth dto.AuthRequestDTO, // body: token + session_id (OBJETIVO a revocar)
	meta dto.RevokeOwnSessionQueryDTO, // body: reason / revoked_by_app (opcionales)
) (*dto.RevokeOwnSessionResponseDTO, error) {

	// 0) Validación mínima
	if auth.Token == "" || auth.SessionID == "" {
		return nil, ErrInvalidToken
	}

	// 1) Lookup rápido por hash del token para obtener la sesión del ejecutor
	hash := security.HashTokenHex(auth.Token)
	tokModel, err := s.tokenRepo.FindByHash(ctx, hash)
	if err != nil {
		return nil, ErrInvalidToken
	}

	// 2) Token debe estar activo y ser access
	if tokModel.Status != models.TokenStatusActive || tokModel.TokenType != models.TokenTypeAccess {
		return nil, ErrInvalidToken
	}

	// 2.5) Expiración por DB
	now := time.Now().UTC()
	if now.After(tokModel.ExpiresAt) {
		_ = s.tokenRepo.MarkExpired(ctx, tokModel.ID, now)
		return nil, security.ErrTokenExpired
	}

	// 3) Cargar sesión del ejecutor (derivada del token)
	execSess, err := s.sessionRepo.FindBySessionID(ctx, tokModel.SessionID)
	if err != nil || execSess == nil {
		return nil, ErrSessionNotFound
	}
	if execSess.Status != models.SessionStatusActive || !execSess.IsActive {
		return nil, ErrSessionInactive
	}

	// 4) Verificar firma JWS RS256 (y validez del JWS)
	claims, vErr := security.VerifyTokenRS256(auth.Token)
	if vErr != nil {
		if errors.Is(vErr, security.ErrTokenExpired) {
			if tokModel.Status == models.TokenStatusActive {
				_ = s.tokenRepo.MarkExpired(ctx, tokModel.ID, time.Now().UTC())
			}
			return nil, security.ErrTokenExpired
		}
		return nil, security.ErrTokenInvalid
	}
	// Defensa: subject debe coincidir con user del token guardado
	if claims.Subject != "" && claims.Subject != tokModel.UserID {
		return nil, ErrInvalidToken
	}

	// 5) Cargar sesión objetivo a revocar (viene en el body como auth.SessionID)
	targetSID := auth.SessionID
	targetSess, err := s.sessionRepo.FindBySessionID(ctx, targetSID)
	if err != nil || targetSess == nil {
		return nil, ErrSessionNotFound
	}

	// 6) La sesión objetivo debe pertenecer al mismo usuario que el ejecutor
	if targetSess.UserID != execSess.UserID {
		return nil, ErrForbiddenRevoke
	}

	// 7) Idempotencia: si ya está revocada, responde OK
	if targetSess.Status == models.SessionStatusRevoked {
		return &dto.RevokeOwnSessionResponseDTO{
			SessionID: targetSess.SessionID,
			Status:    models.SessionStatusRevoked,
			RevokedAt: targetSess.RevokedAt,
			Message:   "Sesión ya estaba revocada",
		}, nil
	}

	// 8) Revocar: actualizar sesión + revocar tokens activos de esa sesión
	revokedAt := time.Now().UTC()
	reason := "user_revoked"
	if meta.Reason != nil && *meta.Reason != "" {
		reason = *meta.Reason
	}
	revokedBy := execSess.UserID
	revokedByApp := ""
	if meta.RevokedByApp != nil {
		revokedByApp = *meta.RevokedByApp
	}

	// 8.1) Cambiar estado de la sesión
	if err := s.sessionRepo.UpdateStatus(ctx, targetSess.ID, models.SessionStatusRevoked, &revokedAt); err != nil {
		return nil, err
	}
	// 8.2) Guardar metadata de revocación
	if err := s.sessionRepo.SetRevocationInfo(ctx, targetSess.ID, reason, revokedBy, revokedByApp); err != nil {
		return nil, err
	}
	// 8.3) Revocar todos los tokens activos de la sesión objetivo
	if _, err := s.tokenRepo.RevokeAllActiveBySessionID(ctx, targetSess.SessionID, revokedAt, reason, revokedBy, revokedByApp); err != nil {
		return nil, err
	}

	// 9) Respuesta
	return &dto.RevokeOwnSessionResponseDTO{
		SessionID: targetSess.SessionID,
		Status:    models.SessionStatusRevoked,
		RevokedAt: &revokedAt,
		Message:   "Sesión revocada exitosamente",
	}, nil
}
