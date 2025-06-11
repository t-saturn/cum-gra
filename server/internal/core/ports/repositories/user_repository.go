package repositories

import "github.com/t-saturn/central-user-manager/server/internal/core/domain/entities"

type UserRepository interface {
	Create(user *entities.User) error
	// Puedes agregar más métodos luego: GetByID, GetAll, etc.
}
