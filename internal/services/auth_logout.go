package services

import (
	"context"

	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/pkg/utils"
)

// Logout cierra una sesión y revoca los tokens asociados.
func (s *AuthService) Logout(ctx context.Context, input dto.LogoutRequestDTO) (*dto.LogoutResponseDTO, error) {
	// 1. Obtener la sesión
	sess, err := s.sessionRepo.FindBySessionID(ctx, input.SessionID)
	if err != nil {
		return nil, ErrSessionNotFound
	}

	// 2. Verificar que esté activa
	if sess.Status != models.SessionStatusActive || !sess.IsActive {
		return nil, ErrSessionInactive
	}

	// 3. Registrar timestamp de revocación
	revokedAt := utils.NowUTC()

	// 4. Marcar la sesión como inactiva
	if err := s.sessionRepo.UpdateStatus(ctx, sess.ID, models.SessionStatusInactive, &revokedAt); err != nil {
		return nil, err
	}

	// 5. Buscar tokens asociados a la sesión
	tokens, err := s.tokenRepo.FindBySessionID(ctx, input.SessionID)
	if err != nil {
		return nil, err
	}

	revoked := make([]string, 0, len(tokens))
	for _, t := range tokens {
		if err := s.tokenRepo.UpdateStatus(ctx, t.ID, models.TokenStatusRevoked, &revokedAt, nil); err != nil {
			return nil, err
		}
		revoked = append(revoked, t.ID.Hex())
	}

	resp := &dto.LogoutResponseDTO{
		SessionID:     input.SessionID,
		RevokedAt:     revokedAt,
		TokensRevoked: revoked,
	}
	return resp, nil
}
