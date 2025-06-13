package repositories

import (
	"github.com/t-saturn/central-user-manager/server/internal/models"
	"github.com/t-saturn/central-user-manager/server/pkg/database"
)

type ApplicationRepository interface {
	Create(app *models.Application) error
}

type applicationRepo struct{}

func NewApplicationRepository() ApplicationRepository {
	return &applicationRepo{}
}

func (r *applicationRepo) Create(app *models.Application) error {
	return database.DB.Create(app).Error
}
