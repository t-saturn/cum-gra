package repositories

import (
	"context"
	"time"

	"github.com/t-saturn/auth-service-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// TokenRepository gestiona los tokens en MongoDB.
type TokenRepository struct {
	col *mongo.Collection
}

// NewTokenRepository crea un repositorio para la colección "tokens".
func NewTokenRepository(db *mongo.Database) *TokenRepository {
	return &TokenRepository{
		col: db.Collection("tokens"),
	}
}

// Insert inserta un nuevo token y devuelve su ObjectID.
func (r *TokenRepository) Insert(ctx context.Context, t *models.Token) (primitive.ObjectID, error) {
	res, err := r.col.InsertOne(ctx, t)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

// FindByHash busca un token por su hash (token_hash) y lo devuelve.
func (r *TokenRepository) FindByHash(ctx context.Context, hash string) (*models.Token, error) {
	var tok models.Token
	err := r.col.FindOne(ctx, bson.M{"token_hash": hash}).Decode(&tok)
	if err != nil {
		return nil, err
	}
	return &tok, nil
}

// UpdateStatus modifica el estado y la fecha de revocación
func (r *TokenRepository) UpdateStatus(ctx context.Context, id primitive.ObjectID, status string, revokedAt, lastUsed *time.Time) error {
	set := bson.M{"status": status}
	if revokedAt != nil {
		set["revoked_at"] = *revokedAt
	}

	_, err := r.col.UpdateByID(ctx, id, bson.M{"$set": set})
	return err
}

// IncrementRefreshCount incrementa en 1 el contador de refresh.
func (r *TokenRepository) IncrementRefreshCount(ctx context.Context, id primitive.ObjectID) error {
	_, err := r.col.UpdateByID(ctx, id, bson.M{"$inc": bson.M{"refresh_count": 1}})
	return err
}

// FindByID recupera un token a partir de su ID en hex string.
func (r *TokenRepository) FindByID(ctx context.Context, tokenID string) (*models.Token, error) {
	oid, err := primitive.ObjectIDFromHex(tokenID)
	if err != nil {
		return nil, err
	}
	var tok models.Token
	if err := r.col.FindOne(ctx, bson.M{"_id": oid}).Decode(&tok); err != nil {
		return nil, err
	}
	return &tok, nil
}
