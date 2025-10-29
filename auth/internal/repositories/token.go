package repositories

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/t-saturn/auth-service-server/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// ErrSessionNotFound indica que no hay ningún token registrado con ese valor.
var (
	ErrSessionNotFound      = errors.New("session not found for given token")
	ErrTokenNotFound        = errors.New("token not found")
	ErrNoTokensInSession    = errors.New("no tokens_generated in session")
	ErrAccessTokenNotFound  = errors.New("access token not found for session")
	ErrRefreshTokenNotFound = errors.New("refresh token not found for session")
)

// MarkExpired marca un token como expirado (status = expired) y actualiza updated_at.
// Nota: usamos filtro por _id y evitamos tocar si ya está en expired.
func (r *TokenRepository) MarkExpired(ctx context.Context, id primitive.ObjectID, at time.Time) error {
	filter := bson.M{
		"_id":    id,
		"status": bson.M{"$ne": models.TokenStatusExpired},
	}
	update := bson.M{
		"$set": bson.M{
			"status":     models.TokenStatusExpired,
			"updated_at": at,
		},
	}
	res, err := r.col.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		// Verificamos si existe; si no existe -> ErrTokenNotFound, si existe ya estaba expirado -> OK
		err := r.col.FindOne(ctx, bson.M{"_id": id}).Err()
		if errors.Is(err, mongo.ErrNoDocuments) {
			return ErrTokenNotFound
		}
		// ya estaba expirado: no es error
	}
	return nil
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

// FindSessionIDByTokenValue busca el token real en el campo token_hash
// y devuelve únicamente el session_id asociado.
func (r *TokenRepository) FindSessionID(ctx context.Context, tokenValue string) (string, error) {
	// Ajusta aquí al campo correcto donde guardas el token sin hash,
	// o al hash si has hasheado la cadena.
	filter := bson.M{"token_hash": tokenValue}

	var result struct {
		SessionID string `bson:"session_id"`
	}
	err := r.col.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return "", ErrSessionNotFound
		}
		return "", err
	}
	if result.SessionID == "" {
		return "", ErrSessionNotFound
	}
	return result.SessionID, nil
}

// FindBySessionID recupera todos los tokens asociados a una sesión.
func (r *TokenRepository) FindBySessionID(ctx context.Context, sessionID string) ([]models.Token, error) {
	filter := bson.M{"session_id": sessionID}
	cur, err := r.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)

	var tokens []models.Token
	for cur.Next(ctx) {
		var t models.Token
		if err := cur.Decode(&t); err != nil {
			return nil, err
		}
		tokens = append(tokens, t)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}
	return tokens, nil
}

// UpdateStatus revoca un token individual actualizando su estado, motivo y metadatos
func (r *TokenRepository) UpdateStatus(ctx context.Context, id primitive.ObjectID, status string, reason string, revokedBy string, revokedByApp string) error {
	if !models.IsValidTokenReason(reason) {
		return fmt.Errorf("razón inválida para revocación de token: %q", reason)
	}

	now := time.Now()

	update := bson.M{
		"$set": bson.M{
			"status":         status,
			"revoked_at":     now,
			"reason":         reason,
			"revoked_by":     revokedBy,
			"revoked_by_app": revokedByApp,
			"updated_at":     now,
		},
	}

	_, err := r.col.UpdateByID(ctx, id, update)
	return err
}

// Insert inserta un nuevo token y devuelve su ObjectID.
func (r *TokenRepository) Insert(ctx context.Context, t *models.Token) (primitive.ObjectID, error) {
	res, err := r.col.InsertOne(ctx, t)
	if err != nil {
		return primitive.NilObjectID, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

// Añade childID al slice ChildTokens del token parentID.
func (r *TokenRepository) AddChildToken(ctx context.Context, parentID, childID primitive.ObjectID) error {
	_, err := r.col.UpdateByID(ctx, parentID, bson.M{
		"$push": bson.M{"child_tokens": childID},
	})
	return err
}

// Guarda el ID de un token “pareado” (access ↔ refresh) en PairedTokenID.
func (r *TokenRepository) SetPairedTokenID(ctx context.Context, tokenID, pairedID primitive.ObjectID) error {
	_, err := r.col.UpdateByID(ctx, tokenID, bson.M{
		"$set": bson.M{"paired_token_id": pairedID},
	})
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

// ListTokenIDsBySession devuelve los IDs de tokens activos asociados a una sesión.
func (r *TokenRepository) ListActiveTokensIDsBySession(ctx context.Context, sessionID string) ([]primitive.ObjectID, error) {
	filter := bson.M{
		"session_id": sessionID,
		"status":     "active",
	}

	cursor, err := r.col.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var tokenIDs []primitive.ObjectID
	for cursor.Next(ctx) {
		var token struct {
			ID primitive.ObjectID `bson:"_id"`
		}
		if err := cursor.Decode(&token); err != nil {
			return nil, err
		}
		tokenIDs = append(tokenIDs, token.ID)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return tokenIDs, nil
}

// RevokeTokensByIDs revoca múltiples tokens por sus IDs.
func (r *TokenRepository) RevokeTokensByIDs(ctx context.Context, tokenIDs []primitive.ObjectID, reason, revokedBy, revokedByApp string) error {
	if !models.IsValidTokenReason(reason) {
		return fmt.Errorf("razón inválida para revocación de token: %q", reason)
	}

	now := time.Now()

	filter := bson.M{
		"_id":    bson.M{"$in": tokenIDs},
		"status": "active",
	}

	update := bson.M{
		"$set": bson.M{
			"status":         "revoked",
			"revoked_at":     now,
			"reason":         reason,
			"revoked_by":     revokedBy,
			"revoked_by_app": revokedByApp,
			"updated_at":     now,
		},
	}

	_, err := r.col.UpdateMany(ctx, filter, update)
	return err
}

func (r *TokenRepository) FindByHashAndType(ctx context.Context, hash, tokenType string) (*models.Token, error) {
	var tok models.Token
	err := r.col.FindOne(ctx, bson.M{"token_hash": hash, "token_type": tokenType}).Decode(&tok)
	if err != nil {
		return nil, err
	}
	return &tok, nil
}

func (r *TokenRepository) MarkRevoked(ctx context.Context, id primitive.ObjectID, reason string, at time.Time) error {
	update := bson.M{
		"$set": bson.M{
			"status":      models.TokenStatusRevoked,
			"revoked_at":  at,
			"revoked_msg": reason,
			"updated_at":  at,
		},
	}
	_, err := r.col.UpdateByID(ctx, id, update)
	return err
}

// GetTokenExpiriesBySessionID:
// 1. Lee la sesión -> tokens_generated
// 2. Busca en "tokens" por esos _id y session_id
// 3. Devuelve el expires_at más reciente para access y refresh (solo activos, opcional)
func (r *SessionRepository) GetTokenExpiriesBySessionID(ctx context.Context, sessionID string) (accessExpiresAt time.Time, refreshExpiresAt time.Time, err error) {
	// 1. Obtener tokens_generated
	var s struct {
		TokensGenerated []primitive.ObjectID `bson:"tokens_generated"`
	}
	if err = r.col.FindOne(
		ctx,
		bson.M{"session_id": sessionID},
		options.FindOne().SetProjection(bson.M{"tokens_generated": 1, "_id": 0}),
	).Decode(&s); err != nil {
		return time.Time{}, time.Time{}, err
	}
	if len(s.TokensGenerated) == 0 {
		return time.Time{}, time.Time{}, ErrNoTokensInSession
	}

	// 2) Consultar tokens
	tokensCol := r.col.Database().Collection("tokens")
	cur, err := tokensCol.Find(
		ctx,
		bson.M{
			"_id":        bson.M{"$in": s.TokensGenerated},
			"session_id": sessionID,
			// Si quieres solo activos, descomenta:
			"status": "active",
		},
		options.Find().SetProjection(bson.M{
			"token_type": 1,
			"expires_at": 1,
			"status":     1,
			"_id":        0,
		}),
	)
	if err != nil {
		return time.Time{}, time.Time{}, err
	}
	defer cur.Close(ctx)

	var (
		accessFound, refreshFound bool
		bestAccess, bestRefresh   time.Time
	)
	for cur.Next(ctx) {
		var t struct {
			TokenType string    `bson:"token_type"`
			ExpiresAt time.Time `bson:"expires_at"`
			Status    string    `bson:"status"`
		}
		if err := cur.Decode(&t); err != nil {
			return time.Time{}, time.Time{}, err
		}
		// Si activaste "status: active" en el filtro, esto sobra. Si no, añade defensa:
		if t.Status != "active" {
			continue
		}

		switch t.TokenType {
		case "access":
			if !accessFound || t.ExpiresAt.After(bestAccess) {
				bestAccess = t.ExpiresAt.UTC()
				accessFound = true
			}
		case "refresh":
			if !refreshFound || t.ExpiresAt.After(bestRefresh) {
				bestRefresh = t.ExpiresAt.UTC()
				refreshFound = true
			}
		}
	}
	if err := cur.Err(); err != nil {
		return time.Time{}, time.Time{}, err
	}

	if !accessFound {
		return time.Time{}, time.Time{}, ErrAccessTokenNotFound
	}
	if !refreshFound {
		return time.Time{}, time.Time{}, ErrRefreshTokenNotFound
	}

	return bestAccess, bestRefresh, nil
}
