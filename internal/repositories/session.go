package repositories

import (
	"context"
	"time"

	"github.com/t-saturn/auth-service-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// SessionRepository gestiona las sesiones en MongoDB.
type SessionRepository struct {
	col *mongo.Collection
}

// NewSessionRepository crea un SessionRepository apuntando a la colección "sessions".
func NewSessionRepository(db *mongo.Database) *SessionRepository {
	return &SessionRepository{
		col: db.Collection("sessions"),
	}
}

// Create inserta una nueva sesión y devuelve su ObjectID.
func (r *SessionRepository) Create(ctx context.Context, s *models.Session) (primitive.ObjectID, error) {
	res, err := r.col.InsertOne(ctx, s)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

// FindByUUID busca una sesión por su UUID.
func (r *SessionRepository) FindByUUID(ctx context.Context, uuid string) (*models.Session, error) {
	var sess models.Session
	err := r.col.FindOne(ctx, bson.M{"uuid": uuid}).Decode(&sess)
	if err != nil {
		return nil, err
	}
	return &sess, nil
}

// UpdateStatus actualiza el status y la fecha de revocación de una sesión.
func (r *SessionRepository) UpdateStatus(ctx context.Context, id primitive.ObjectID, status string, revokedAt *time.Time) error {
	update := bson.M{"$set": bson.M{"status": status}}
	if revokedAt != nil {
		update["$set"].(bson.M)["revoked_at"] = *revokedAt
	}
	_, err := r.col.UpdateByID(ctx, id, update)
	return err
}
