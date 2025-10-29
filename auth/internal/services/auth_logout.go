package services

import (
	"context"
	"fmt"

	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/pkg/utils"
)

// Logout cierra una sesión y revoca los tokens asociados.
func (s *AuthService) Logout(ctx context.Context, refreshToken, reason, sessionID string) (*dto.LogoutResponseDTO, error) {
	// 1. Verificar si se proporcionó un sessionID
	if sessionID != "" {
		// Si hay sessionID, construir DTO con el sessionID proporcionado
		input := dto.LogoutRequestDTO{
			SessionID:  sessionID,
			LogoutType: reason,
		}
		// Llamar directamente a logoutBySession con el sessionID proporcionado
		return s.logoutBySession(ctx, input)
	}

	// 2. Si no hay sessionID, seguir el flujo actual: extraer sessionID del token
	sessionID, err := s.tokenRepo.FindSessionID(ctx, refreshToken)
	if err != nil {
		return nil, ErrSessionNotFound
	}
	fmt.Print("SessionID:", sessionID)

	// 3. Construir un DTO interno con sessionID y reason
	input := dto.LogoutRequestDTO{
		SessionID:  sessionID,
		LogoutType: reason,
	}

	// 4. Llamar a la lógica de logout usando el sessionID extraído
	return s.logoutBySession(ctx, input)
}

func (s *AuthService) logoutBySession(ctx context.Context, input dto.LogoutRequestDTO) (*dto.LogoutResponseDTO, error) {
	// 1. Obtener la sesión
	sess, err := s.sessionRepo.FindBySessionID(ctx, input.SessionID)
	if err != nil {
		return nil, ErrSessionNotFound
	}

	// 2. Verificar que esté activa
	if sess.Status != models.SessionStatusActive || !sess.IsActive {
		return nil, ErrSessionInactive
	}

	// 3. Validar que el logout_type (reason) sea válido
	if !models.IsValidTokenReason(input.LogoutType) {
		return nil, fmt.Errorf("tipo de logout inválido: %q", input.LogoutType)
	}

	// 4. Registrar timestamp de revocación
	revokedAt := utils.NowUTC()

	// 5. Marcar la sesión como inactiva
	if err := s.sessionRepo.UpdateStatus(ctx, sess.ID, models.SessionStatusInactive, &revokedAt); err != nil {
		return nil, err
	}

	// 6. Buscar tokens asociados a la sesión
	tokens, err := s.tokenRepo.FindBySessionID(ctx, input.SessionID)
	if err != nil {
		return nil, err
	}

	// 7. Revocar cada token
	revoked := make([]string, 0, len(tokens))
	for _, t := range tokens {
		revokedBy := t.UserID
		revokedByApp := t.SessionID

		if err := s.tokenRepo.UpdateStatus(ctx, t.ID, models.TokenStatusRevoked, input.LogoutType, revokedBy, revokedByApp); err != nil {
			return nil, err
		}
		revoked = append(revoked, t.ID.Hex())
	}

	// 8. Retornar respuesta
	resp := &dto.LogoutResponseDTO{
		SessionID:     input.SessionID,
		RevokedAt:     revokedAt,
		TokensRevoked: revoked,
	}
	return resp, nil
}
