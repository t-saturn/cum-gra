package repository

import (
	"context"

	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
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

// CreateInvalidToken inserts a new invalid token document
func CreateInvalidToken(ctx context.Context, token *models.InvalidToken) error {
	_, err := getInvalidTokenCollection().InsertOne(ctx, token)
	if err != nil {
		config.Logger.WithError(err).Error("Error inserting InvalidToken")
	}
	return err
}

// GetAllInvalidTokens returns all invalid tokens
func GetAllInvalidTokens(ctx context.Context) ([]*models.InvalidToken, error) {
	cursor, err := getInvalidTokenCollection().Find(ctx, bson.M{})
	if err != nil {
		config.Logger.WithError(err).Error("Error fetching InvalidTokens")
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []*models.InvalidToken
	for cursor.Next(ctx) {
		var token models.InvalidToken
		if err := cursor.Decode(&token); err != nil {
			config.Logger.WithError(err).Warn("Error decoding InvalidToken")
			continue
		}
		results = append(results, &token)
	}

	return results, nil
}
