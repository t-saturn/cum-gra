package repository

import (
	"context"

	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/models"
)

func InsertToken(ctx context.Context, token *models.Token) error {
	_, err := config.MongoDatabase.Collection("tokens").InsertOne(ctx, token)
	if err != nil {
		return err
	}

	return nil
}
