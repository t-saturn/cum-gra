package repositories

import (
	"context"

	"github.com/t-saturn/auth-service-server/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
)

// CaptchaLogRepository gestiona los logs de CAPTCHA.
type CaptchaLogRepository struct {
	col *mongo.Collection
}

func NewCaptchaLogRepository(db *mongo.Database) *CaptchaLogRepository {
	return &CaptchaLogRepository{col: db.Collection("captcha_logs")}
}

func (r *CaptchaLogRepository) Insert(ctx context.Context, c *models.CaptchaLog) error {
	_, err := r.col.InsertOne(ctx, c)
	return err
}
