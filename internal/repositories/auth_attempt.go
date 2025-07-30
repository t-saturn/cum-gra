package repositories

import (
	"context"

	"github.com/t-saturn/auth-service-server/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// AuthAttemptRepository gestiona la persistencia de AuthAttempt en MongoDB.
type AuthAttemptRepository struct {
	col *mongo.Collection
}

// NewAuthAttemptRepository construye un repositorio a partir de la conexión Mongo.
func NewAuthAttemptRepository(db *mongo.Database) *AuthAttemptRepository {
	return &AuthAttemptRepository{
		col: db.Collection("auth_attempts"),
	}
}

// Insert guarda un AuthAttempt en la colección.
// Asume que los datos ya vienen validados en capas superiores.
func (r *AuthAttemptRepository) Insert(ctx context.Context, a *models.AuthAttempt) error {
	_, err := r.col.InsertOne(ctx, a)
	return err
}
