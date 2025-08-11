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

// ListSessions valida el access token + session y retorna las sesiones del usuario autenticado.
func (s *AuthService) ListSessions(ctx context.Context, auth dto.AuthRequestDTO, q dto.ListSessionsQueryDTO) (*dto.ListSessionsResponseDTO, error) {
	// 0) Validación mínima
	if auth.Token == "" || auth.SessionID == "" {
		return nil, ErrInvalidToken
	}

	// 1) Lookup rápido por hash
	hash := security.HashTokenHex(auth.Token)
	tokModel, err := s.tokenRepo.FindByHash(ctx, hash)
	if err != nil {
		return nil, ErrInvalidToken
	}

	// 2) Debe estar activo y ser access token
	if tokModel.Status != models.TokenStatusActive {
		return nil, ErrInvalidToken
	}
	if tokModel.TokenType != models.TokenTypeAccess {
		return nil, ErrInvalidToken
	}

	// 2.5) Expiración por DB
	now := time.Now().UTC()
	if now.After(tokModel.ExpiresAt) {
		_ = s.tokenRepo.MarkExpired(ctx, tokModel.ID, now)
		return nil, security.ErrTokenExpired
	}

	// 3) Coincidencia de session_id
	if tokModel.SessionID != auth.SessionID {
		return nil, ErrSessionMismatch
	}

	// 4) Cargar sesión
	sessModel, err := s.sessionRepo.FindBySessionID(ctx, auth.SessionID)
	if err != nil || sessModel == nil {
		return nil, ErrSessionNotFound
	}

	// 5) Verificar estado de sesión
	if sessModel.Status != models.SessionStatusActive || !sessModel.IsActive {
		return nil, ErrSessionInactive
	}

	// 6) Verificar firma JWS RS256 (y exp del JWS)
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
	// Defensa adicional: subject debe coincidir con user del token
	if claims.Subject != "" && claims.Subject != tokModel.UserID {
		return nil, ErrInvalidToken
	}

	// 7) Listar sesiones del usuario
	//    user_id sale de la sesión ya validada
	//    (si prefieres usar claims.Subject, también es válido si coincide)
	_, err = uuid.Parse(sessModel.UserID)
	if err != nil {
		return nil, repositories.ErrUserNotFound
	}

	sessions, total, err := s.sessionRepo.FindByUserIDPaged(ctx, sessModel.UserID, q)
	if err != nil {
		return nil, err
	}

	// 8) Mapear a DTO
	data := make([]dto.SessionViewDTO, 0, len(sessions))
	for i := range sessions {
		data = append(data, toSessionViewDTO(&sessions[i]))
	}

	// 9) Paginar
	page := q.Page
	if page < 1 {
		page = 1
	}
	limit := q.Limit
	if limit <= 0 {
		limit = 20
	}
	hasPrev := page > 1
	hasNext := int64(page*limit) < total

	resp := &dto.ListSessionsResponseDTO{
		Data: data,
		Pagination: dto.PaginationMeta{
			Page:    page,
			Limit:   limit,
			Total:   total,
			HasPrev: hasPrev,
			HasNext: hasNext,
		},
	}

	return resp, nil
}
