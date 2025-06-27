package repositories

import (
	"context"
	"time"

	"github.com/t-saturn/auth-service-server/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

type TokenRepository interface {
	Insert(ctx context.Context, token *models.Token) error
}

type tokenRepository struct {
	collection *mongo.Collection
}

func NewTokenRepository(db *mongo.Database) TokenRepository {
	return &tokenRepository{
		collection: db.Collection("tokens"),
	}
}

func (r *tokenRepository) Insert(ctx context.Context, token *models.Token) error {
	token.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, token)
	return err
}
