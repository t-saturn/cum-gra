package repository

import (
	"context"
	"time"

	"github.com/t-saturn/auth-service-server/internal/config"
	"github.com/t-saturn/auth-service-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertToken(ctx context.Context, token *models.Token) error {
	_, err := config.MongoDatabase.Collection("tokens").InsertOne(ctx, token)
	if err != nil {
		return err
	}

	return nil
}

func FindTokenByJTI(ctx context.Context, jti string) (*models.Token, error) {
	var token models.Token
	err := config.MongoDatabase.Collection("tokens").FindOne(ctx, bson.M{
		"tokenId": jti,
	}).Decode(&token)
	return &token, err
}

func IncrementRefreshCount(ctx context.Context, jti string) error {
	_, err := config.MongoDatabase.Collection("tokens").UpdateOne(ctx,
		bson.M{"tokenId": jti},
		bson.M{"$inc": bson.M{"refreshCount": 1}, "$set": bson.M{"updatedAt": time.Now()}})
	return err
}
