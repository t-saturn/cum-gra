package repositories

import (
	"context"

	"github.com/t-saturn/auth-service-server/internal/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// VerifyAttemptRepository gestiona la persistencia de los intentos de verificación de credenciales.
type VerifyAttemptRepository struct {
	col *mongo.Collection
}

// NewVerifyAttemptRepository crea un repositorio para la colección "verify_attempts".
func NewVerifyAttemptRepository(db *mongo.Database) *VerifyAttemptRepository {
	return &VerifyAttemptRepository{col: db.Collection("verify_attempts")}
}

func (r *VerifyAttemptRepository) Insert(ctx context.Context, a *models.VerifyAttempt) (primitive.ObjectID, error) {
	res, err := r.col.InsertOne(ctx, a)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}
