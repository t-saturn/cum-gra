package repositories

import (
	"context"

	"github.com/t-saturn/auth-service-server/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// CaptchaRepository gestiona los logs de CAPTCHA.
type CaptchaRepository struct {
	col *mongo.Collection
}

func NewCaptchaRepository(db *mongo.Database) *CaptchaRepository {
	return &CaptchaRepository{col: db.Collection("captcha_logs")}
}

func (r *CaptchaRepository) Insert(ctx context.Context, c *models.Captcha) error {
	_, err := r.col.InsertOne(ctx, c)
	return err
}
