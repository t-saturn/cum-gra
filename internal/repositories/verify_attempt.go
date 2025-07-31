package repositories

import (
	"context"

	"github.com/t-saturn/auth-service-server/internal/models"
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

// Insert guarda un VerifyAttempt en la colección.
// Se asume que los datos ya han sido validados en capas superiores.
func (r *VerifyAttemptRepository) Insert(ctx context.Context, a *models.VerifyAttempt) error {
	_, err := r.col.InsertOne(ctx, a)
	return err
}
