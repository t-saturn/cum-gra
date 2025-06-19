package postgres

import (
	"time"

	"github.com/central-user-manager/internal/core/domain"
	"github.com/central-user-manager/internal/core/ports/repositories"
	"github.com/central-user-manager/internal/infrastructure/database"
	"github.com/google/uuid"
)

type applicationRepository struct{}

func NewApplicationRepository() repositories.ApplicationRepository {
	return &applicationRepository{}
}

func (r *applicationRepository) Create(app *domain.Application) error {
	return database.DB.Create(app).Error
}

func (r *applicationRepository) GetAll() ([]domain.Application, error) {
	var apps []domain.Application
	err := database.DB.Where("is_deleted = false").Find(&apps).Error
	return apps, err
}

func (r *applicationRepository) GetByID(id uuid.UUID) (*domain.Application, error) {
	var app domain.Application
	err := database.DB.First(&app, "id = ? AND is_deleted = false", id).Error
	return &app, err
}

func (r *applicationRepository) Update(app *domain.Application) error {
	return database.DB.Save(app).Error
}

func (r *applicationRepository) Delete(id uuid.UUID) error {
	return database.DB.Model(&domain.Application{}).
		Where("id = ?", id).
		Updates(map[string]interface{}{
			"is_deleted": true,
			"deleted_at": time.Now(),
		}).Error
}

func (r *applicationRepository) ExistsByClientID(clientID string) (bool, error) {
	var count int64
	err := database.DB.Model(&domain.Application{}).
		Where("client_id = ?", clientID).
		Count(&count).Error
	return count > 0, err
}
