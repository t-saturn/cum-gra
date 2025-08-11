package services

import (
	"context"
	"errors"

	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/pkg/utils"
)

var (
	ErrSessionAlreadyRevoked = errors.New("session_already_revoked")
)

// RevokeSession revoca una sesión y todos sus tokens asociados.
func (s *AuthService) RevokeSession(ctx context.Context, sessionID, reason, revokedBy, revokedByApp string) (*dto.RevokeOwnSessionQueryDTO, error) {
	now := utils.NowUTC()

	// 1) Buscar la sesión
	sess, err := s.sessionRepo.FindBySessionID(ctx, sessionID)
	if err != nil {
		return nil, ErrSessionNotFound
	}

	// 1b) Si ya está revocada
	if sess.Status == models.SessionStatusRevoked {
		return nil, ErrSessionAlreadyRevoked
	}

	// 2) Revocar la sesión
	if err := s.sessionRepo.UpdateStatus(ctx, sess.ID, models.SessionStatusRevoked, &now); err != nil {
		return nil, err
	}
	_ = s.sessionRepo.SetRevocationInfo(ctx, sess.ID, reason, revokedBy, revokedByApp)

	// 3) Listar y revocar tokens activos...
	tokenIDs, err := s.tokenRepo.ListActiveTokensIDsBySession(ctx, sessionID)
	if err != nil {
		return nil, err
	}
	if len(tokenIDs) > 0 {
		if err := s.tokenRepo.RevokeTokensByIDs(ctx, tokenIDs, reason, revokedBy, revokedByApp); err != nil {
			return nil, err
		}
	}

	// 4) Construir DTO
	hexIDs := make([]string, len(tokenIDs))
	for i, oid := range tokenIDs {
		hexIDs[i] = oid.Hex()
	}
	resp := &dto.RevokeOwnSessionQueryDTO{
		SessionID: sessionID,
	}
	return resp, nil
}
