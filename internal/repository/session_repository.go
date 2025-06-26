package repository

import (
	"context"

	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/models"
)

func InsertSession(ctx context.Context, session *models.Session) error {
	_, err := config.MongoDatabase.Collection("sessions").InsertOne(ctx, session)
	return err
}
