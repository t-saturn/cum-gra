package postgres

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/t-saturn/central-user-manager/internal/core/domain"
	"github.com/t-saturn/central-user-manager/internal/core/ports/repositories"
	"github.com/t-saturn/central-user-manager/internal/infrastructure/database"
	"gorm.io/gorm"
)

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type applicationRepository struct{}

func NewApplicationRepository() repositories.ApplicationRepository {
	return &applicationRepository{}
}

func (r *applicationRepository) ExistsByName(name string) (bool, error) {
	var count int64
	err := database.DB.Model(&domain.Application{}).Where("name = ? AND is_deleted = false", name).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *applicationRepository) ExistsByNameExceptID(name string, excludeID uuid.UUID) (bool, error) {
	var count int64
	err := database.DB.Model(&domain.Application{}).Where("LOWER(name) = LOWER(?) AND id <> ? AND is_deleted = false", name, excludeID).Count(&count).Error

	return count > 0, err
}

func (r *applicationRepository) Create(ctx context.Context, app *domain.Application) (*domain.Application, error) {
	app.ID = uuid.New()
	app.CreatedAt = time.Now()
	app.UpdatedAt = time.Now()
	app.Status = "active"
	app.IsDeleted = false

	if err := database.DB.WithContext(ctx).Create(app).Error; err != nil {
		return nil, err
	}

	return app, nil
}

func (r *applicationRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Application, error) {
	var app domain.Application
	err := database.DB.WithContext(ctx).Where("id = ? AND is_deleted = false", id).First(&app).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	return &app, nil
}

func (r *applicationRepository) Update(ctx context.Context, id uuid.UUID, app *domain.Application) (*domain.Application, error) {
	var existing domain.Application
	db := database.DB.WithContext(ctx)

	// Verifica si existe
	if err := db.First(&existing, "id = ? AND is_deleted = false", id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}

	// Solo actualiza los campos no nulos
	if app.Name != "" {
		existing.Name = app.Name
	}
	if *app.Description != "" {
		existing.Description = app.Description
	}
	if app.Status != "" {
		existing.Status = app.Status
	}
	if app.ClientID != "" {
		existing.ClientID = app.ClientID
	}
	if app.ClientSecret != "" {
		existing.ClientSecret = app.ClientSecret
	}
	if app.Domain != "" {
		existing.Domain = app.Domain
	}
	if *app.Logo != "" {
		existing.Logo = app.Logo
	}
	if app.CallbackUrls != nil {
		existing.CallbackUrls = app.CallbackUrls
	}
	if app.IsFirstParty != false {
		existing.IsFirstParty = app.IsFirstParty
	}

	existing.UpdatedAt = time.Now()

	if err := db.Save(&existing).Error; err != nil {
		return nil, err
	}

	return &existing, nil
}

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type applicationRoleRepository struct{}

func NewApplicationRoleRepository() repositories.ApplicationRoleRepository {
	return &applicationRoleRepository{}
}

func (r *applicationRoleRepository) ExistsByName(name string) (bool, error) {
	var count int64
	err := database.DB.Model(&domain.ApplicationRole{}).Where("name = ? AND is_deleted = false", name).Count(&count).Error
	if err != nil {
		return false, err
	}

	return count > 0, nil
}

func (r *applicationRoleRepository) ExistsByNameExceptID(name string, excludeID uuid.UUID) (bool, error) {
	var count int64
	err := database.DB.Model(&domain.ApplicationRole{}).Where("LOWER(name) = LOWER(?) AND id <> ? AND is_deleted = false", name, excludeID).Count(&count).Error

	return count > 0, err
}

func (r *applicationRoleRepository) Create(ctx context.Context, data *domain.ApplicationRole) (*domain.ApplicationRole, error) {
	data.ID = uuid.New()
	data.CreatedAt = time.Now()
	data.UpdatedAt = time.Now()
	data.IsDeleted = false

	if err := database.DB.WithContext(ctx).Create(data).Error; err != nil {
		return nil, err
	}

	return data, nil
}

/** ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------- */
type userApplicationRoleRepository struct{}

func NewUserApplicationRepository() repositories.UserApplicationRoleRepository {
	return &userApplicationRoleRepository{}
}
