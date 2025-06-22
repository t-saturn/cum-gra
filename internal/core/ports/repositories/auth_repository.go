// internal/core/ports/repositories/auth_repository.go
package repositories

import (
	"github.com/t-saturn/central-user-manager/internal/core/domain"
)

type AuthRepository interface {
	FindByEmail(email string) (*domain.User, error)
}
