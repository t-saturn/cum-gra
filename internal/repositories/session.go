package repositories

import (
	"context"
	"time"

	"github.com/t-saturn/auth-service-server/internal/dto"
	"github.com/t-saturn/auth-service-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// Insert inserta una nueva sesión y devuelve su ObjectID.
func (r *SessionRepository) Insert(ctx context.Context, s *models.Session) (primitive.ObjectID, error) {
	res, err := r.col.InsertOne(ctx, s)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

// FindBySessionID busca una sesión por su SessionID (campo session_id).
func (r *SessionRepository) FindBySessionID(ctx context.Context, sessionID string) (*models.Session, error) {
	var sess models.Session
	// Filtrar por el campo "session_id"
	err := r.col.FindOne(ctx, bson.M{"session_id": sessionID}).Decode(&sess)
	if err != nil {
		return nil, err
	}
	return &sess, nil
}

// SetRevocationInfo actualiza la sesión con información de revocación
func (r *SessionRepository) SetRevocationInfo(ctx context.Context, id primitive.ObjectID, reason, revokedBy, revokedByApp string) error {
	update := bson.M{"$set": bson.M{
		"revocation_reason": reason,
		"revoked_by":        revokedBy,
		"revoked_by_app":    revokedByApp,
	}}
	_, err := r.col.UpdateByID(ctx, id, update)
	return err
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

// FindByUserID retorna todas las sesiones de un usuario según filtros.
func (r *SessionRepository) FindByUserID(ctx context.Context, userID string, params dto.ListSessionsQueryDTO) ([]models.Session, error) {
	// Construir filtro base
	filter := bson.M{"user_id": userID}

	// Filtros opcionales
	if params.Status != nil {
		filter["status"] = *params.Status
	}
	if params.IsActive != nil {
		filter["is_active"] = *params.IsActive
	}

	// Opciones de consulta: orden por defecto de sesiones más recientes primero
	findOpts := options.Find()
	findOpts.SetSort(bson.D{{Key: "created_at", Value: -1}})

	// Ejecutar consulta
	cur, err := r.col.Find(ctx, filter, findOpts)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	// Leer resultados
	sessions := make([]models.Session, 0)
	for cur.Next(ctx) {
		var s models.Session
		if err := cur.Decode(&s); err != nil {
			return nil, err
		}
		sessions = append(sessions, s)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return sessions, nil
}

// CountByUserID cuenta cuántas sesiones tiene un usuario, opcionalmente filtrando por estado.
func (r *SessionRepository) CountByUserID(ctx context.Context, userID string, status *string) (int64, error) {
	filter := bson.M{"user_id": userID}
	if status != nil {
		filter["status"] = *status
	}
	count, err := r.col.CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, nil
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

// AddTokenToSession añade el ObjectID de un token al campo TokensGenerated de una sesión.
func (r *SessionRepository) AddTokenToSession(ctx context.Context, sessionID string, tokenID primitive.ObjectID) error {
	filter := bson.M{"session_id": sessionID}
	update := bson.M{"$push": bson.M{"tokens_generated": tokenID}}
	_, err := r.col.UpdateOne(ctx, filter, update)
	return err
}
