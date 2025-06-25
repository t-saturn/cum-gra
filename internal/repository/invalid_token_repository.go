package repository

import (
	"context"

	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

const invalidTokenCollection = "tokens_invalid"

func getInvalidTokenCollection() *mongo.Collection {
	return config.MongoDatabase.Collection(invalidTokenCollection)
}

func InsertInvalidToken(ctx context.Context, token *models.InvalidToken) error {
	_, err := getInvalidTokenCollection().InsertOne(ctx, token)
	if err != nil {
		config.Logger.WithError(err).Error("Error inserting InvalidToken")
	}
	return err
}
