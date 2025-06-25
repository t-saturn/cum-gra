package repository

import (
	"context"

	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const activeTokenCollection = "tokens_activos"

func getActiveTokenCollection() *mongo.Collection {
	return config.MongoDatabase.Collection(activeTokenCollection)
}

// CreateActiveToken inserts a new token into the collection
func CreateActiveToken(ctx context.Context, token *models.ActiveToken) error {
	_, err := getActiveTokenCollection().InsertOne(ctx, token)
	if err != nil {
		config.Logger.WithError(err).Error("Error inserting ActiveToken")
		return err
	}
	config.Logger.WithField("token_id", token.TokenID).Info("✅ ActiveToken inserted")
	return nil
}

// GetAllActiveTokens returns all active tokens in the system
func GetAllActiveTokens(ctx context.Context) ([]*models.ActiveToken, error) {
	cursor, err := getActiveTokenCollection().Find(ctx, bson.M{})
	if err != nil {
		config.Logger.WithError(err).Error("Error fetching ActiveTokens")
		return nil, err
	}
	defer cursor.Close(ctx)

	var tokens []*models.ActiveToken
	for cursor.Next(ctx) {
		var token models.ActiveToken
		if err := cursor.Decode(&token); err != nil {
			config.Logger.WithError(err).Warn("Error decoding ActiveToken")
			continue
		}
		tokens = append(tokens, &token)
	}
	return tokens, nil
}
