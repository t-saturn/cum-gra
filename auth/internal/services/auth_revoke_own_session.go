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

// RevokeOwnSession valida token (Authorization) y revoca la sesión objetivo (meta.SessionID) + tokens activos.
func (s *AuthService) RevokeOwnSession(ctx context.Context, accessToken string, meta dto.RevokeOwnSessionQueryDTO) (*dto.RevokeOwnSessionResponseDTO, error) {

	// 0. Validación mínima
	if accessToken == "" || meta.SessionID == "" {
		return nil, ErrInvalidToken
	}

	// 1. Lookup por hash del token → sesión del ejecutor
	hash := security.HashTokenHex(accessToken)
	tokModel, err := s.tokenRepo.FindByHash(ctx, hash)
	if err != nil {
		return nil, ErrInvalidToken
	}

	// 2. Token activo y de tipo access
	if tokModel.Status != models.TokenStatusActive || tokModel.TokenType != models.TokenTypeAccess {
		return nil, ErrInvalidToken
	}

	// 2.5. Expiración por DB
	now := time.Now().UTC()
	if now.After(tokModel.ExpiresAt) {
		_ = s.tokenRepo.MarkExpired(ctx, tokModel.ID, now)
		return nil, security.ErrTokenExpired
	}

	// 3. Cargar sesión del ejecutor (derivada del token)
	execSess, err := s.sessionRepo.FindBySessionID(ctx, tokModel.SessionID)
	if err != nil || execSess == nil {
		return nil, ErrSessionNotFound
	}
	if execSess.Status != models.SessionStatusActive || !execSess.IsActive {
		return nil, ErrSessionInactive
	}

	// 4. Verificar firma JWS RS256
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
	if claims.Subject != "" && claims.Subject != tokModel.UserID {
		return nil, ErrInvalidToken
	}

	// 5. Cargar sesión objetivo (viene en meta.SessionID)
	targetSID := meta.SessionID
	targetSess, err := s.sessionRepo.FindBySessionID(ctx, targetSID)
	if err != nil || targetSess == nil {
		return nil, ErrSessionNotFound
	}

	// 6. Debe pertenecer al mismo usuario
	if targetSess.UserID != execSess.UserID {
		return nil, ErrForbiddenRevoke
	}

	// 7. Idempotencia
	if targetSess.Status == models.SessionStatusRevoked {
		return &dto.RevokeOwnSessionResponseDTO{
			SessionID: targetSess.SessionID,
			Status:    models.SessionStatusRevoked,
			RevokedAt: targetSess.RevokedAt,
			Message:   "Sesión ya estaba revocada",
		}, nil
	}

	// 8. Revocar sesión + tokens
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

	// 8.3) Revocar todos los tokens activos y obtener la lista
	revokedTokens, _, err := s.tokenRepo.RevokeAllActiveBySessionIDReturn(ctx, targetSess.SessionID, revokedAt, reason, revokedBy, revokedByApp)
	if err != nil {
		return nil, err
	}

	// 8.4) Mapear tokens a DTO
	tokensDTO := make([]dto.RevokedTokenDetailDTO, 0, len(revokedTokens))
	for _, t := range revokedTokens {
		// Nota: usa el campo que tengas disponible como "identificador visible".
		// Si no tienes token_id poblado, usa el ObjectID:
		tokenID := t.TokenID
		if tokenID == "" && !t.ID.IsZero() {
			tokenID = t.ID.Hex()
		}

		var expPtr *time.Time
		if !t.ExpiresAt.IsZero() {
			expCopy := t.ExpiresAt
			expPtr = &expCopy
		}

		tokensDTO = append(tokensDTO, dto.RevokedTokenDetailDTO{
			TokenID:   tokenID,
			TokenType: t.TokenType,
			RevokedAt: t.RevokedAt,
			Reason:    t.Reason,
			ExpiresAt: expPtr,
		})
	}

	// 9) Respuesta
	return &dto.RevokeOwnSessionResponseDTO{
		SessionID: targetSess.SessionID,
		Status:    models.SessionStatusRevoked,
		RevokedAt: &revokedAt,
		Message:   "Sesión revocada exitosamente",
		Tokens:    tokensDTO,
	}, nil
}
