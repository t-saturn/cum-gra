package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/models"
	"github.com/t-saturn/auth-service-server/internal/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LoginResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	SessionID    string    `json:"session_id"`
	AccessExp    time.Time `json:"access_exp"`
	RefreshExp   time.Time `json:"refresh_exp"`
}

// Simulación de validación de usuario (reemplazar por lógica real)
func validateUser(email, password string) (primitive.ObjectID, error) {
	// TODO: Buscar en colección de usuarios reales y verificar hash bcrypt
	if email == "usuario@dominio.com" && password == "secreta123" {
		return primitive.NewObjectID(), nil // ID falso por ahora
	}
	return primitive.NilObjectID, errors.New("credenciales incorrectas")
}

func LoginUser(email, password, appID, appURL string, device models.SessionDeviceInfo) (*LoginResponse, error) {
	userID, err := validateUser(email, password)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 1. Crear sesión
	sessionID := uuid.New().String()
	now := time.Now()
	expires := now.Add(2 * time.Hour) // duración fija para la sesión

	session := &models.Session{
		SessionID:    sessionID,
		UserID:       userID,
		IsActive:     true,
		CreatedAt:    now,
		LastActivity: now,
		ExpiresAt:    expires,
		DeviceInfo:   device,
	}

	if err := repository.InsertSession(ctx, session); err != nil {
		return nil, errors.New("no se pudo crear la sesión")
	}

	// 2. Generar tokens
	accessToken, accessJti, accessExp, _ := config.GenerateJWT(userID.Hex(), "access")
	refreshToken, refreshJti, refreshExp, _ := config.GenerateJWT(userID.Hex(), "refresh")

	// 3. Guardar tokens
	hash := func(token string) string {
		h := sha256.Sum256([]byte(token))
		return hex.EncodeToString(h[:])
	}

	accessDoc := &models.Token{
		TokenID:        accessJti,
		TokenHash:      hash(accessToken),
		UserID:         userID,
		SessionID:      &session.ID,
		Status:         "active",
		TokenType:      "access",
		IssuedAt:       now,
		ExpiresAt:      accessExp,
		CreatedAt:      now,
		UpdatedAt:      now,
		ApplicationID:  appID,
		ApplicationURL: appURL,
		DeviceInfo:     device.ToDeviceInfo(),
	}

	refreshDoc := &models.Token{
		TokenID:         refreshJti,
		TokenHash:       hash(refreshToken),
		UserID:          userID,
		SessionID:       &session.ID,
		Status:          "active",
		TokenType:       "refresh",
		IssuedAt:        now,
		ExpiresAt:       refreshExp,
		CreatedAt:       now,
		UpdatedAt:       now,
		ApplicationID:   appID,
		ApplicationURL:  appURL,
		DeviceInfo:      device.ToDeviceInfo(),
		RefreshCount:    0,
		MaxRefreshCount: 10,
	}

	if err := repository.InsertToken(ctx, accessDoc); err != nil {
		return nil, err
	}
	if err := repository.InsertToken(ctx, refreshDoc); err != nil {
		return nil, err
	}

	// 4. Devolver respuesta
	return &LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		SessionID:    sessionID,
		AccessExp:    accessExp,
		RefreshExp:   refreshExp,
	}, nil
}
