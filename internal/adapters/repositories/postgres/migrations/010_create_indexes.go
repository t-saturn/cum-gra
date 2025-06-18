package migrations

import (
	"log"

	"github.com/central-user-manager/internal/infrastructure/database"
)

func CreateIndexes() {
	indexes := []string{
		// Índices para users
		"CREATE INDEX IF NOT EXISTS idx_users_email ON users(email);",
		"CREATE INDEX IF NOT EXISTS idx_users_dni ON users(dni);",
		"CREATE INDEX IF NOT EXISTS idx_users_status ON users(status);",
		"CREATE INDEX IF NOT EXISTS idx_users_is_deleted ON users(is_deleted);",
		"CREATE INDEX IF NOT EXISTS idx_users_structural_position_id ON users(structural_position_id);",
		"CREATE INDEX IF NOT EXISTS idx_users_organic_unit_id ON users(organic_unit_id);",

		// Índices para applications
		"CREATE INDEX IF NOT EXISTS idx_applications_client_id ON applications(client_id);",
		"CREATE INDEX IF NOT EXISTS idx_applications_status ON applications(status);",
		"CREATE INDEX IF NOT EXISTS idx_applications_is_deleted ON applications(is_deleted);",

		// Índices para application_roles
		"CREATE INDEX IF NOT EXISTS idx_application_roles_application_id ON application_roles(application_id);",
		"CREATE INDEX IF NOT EXISTS idx_application_roles_is_deleted ON application_roles(is_deleted);",

		// Índices para user_application_roles
		"CREATE INDEX IF NOT EXISTS idx_user_application_roles_user_id ON user_application_roles(user_id);",
		"CREATE INDEX IF NOT EXISTS idx_user_application_roles_application_id ON user_application_roles(application_id);",
		"CREATE INDEX IF NOT EXISTS idx_user_application_roles_role_id ON user_application_roles(application_role_id);",
		"CREATE INDEX IF NOT EXISTS idx_user_application_roles_is_deleted ON user_application_roles(is_deleted);",

		// Índices para modules
		"CREATE INDEX IF NOT EXISTS idx_modules_parent_id ON modules(parent_id);",
		"CREATE INDEX IF NOT EXISTS idx_modules_application_id ON modules(application_id);",
		"CREATE INDEX IF NOT EXISTS idx_modules_status ON modules(status);",
		"CREATE INDEX IF NOT EXISTS idx_modules_is_deleted ON modules(is_deleted);",

		// Índices para module_role_permissions
		"CREATE INDEX IF NOT EXISTS idx_module_role_permissions_module_id ON module_role_permissions(module_id);",
		"CREATE INDEX IF NOT EXISTS idx_module_role_permissions_role_id ON module_role_permissions(application_role_id);",
		"CREATE INDEX IF NOT EXISTS idx_module_role_permissions_is_deleted ON module_role_permissions(is_deleted);",

		// Índices para user_module_restrictions
		"CREATE INDEX IF NOT EXISTS idx_user_module_restrictions_user_id ON user_module_restrictions(user_id);",
		"CREATE INDEX IF NOT EXISTS idx_user_module_restrictions_module_id ON user_module_restrictions(module_id);",
		"CREATE INDEX IF NOT EXISTS idx_user_module_restrictions_application_id ON user_module_restrictions(application_id);",
		"CREATE INDEX IF NOT EXISTS idx_user_module_restrictions_is_deleted ON user_module_restrictions(is_deleted);",

		// Índices para structural_positions
		"CREATE INDEX IF NOT EXISTS idx_structural_positions_code ON structural_positions(code);",
		"CREATE INDEX IF NOT EXISTS idx_structural_positions_is_deleted ON structural_positions(is_deleted);",

		// Índices para organic_units
		"CREATE INDEX IF NOT EXISTS idx_organic_units_parent_id ON organic_units(parent_id);",
		"CREATE INDEX IF NOT EXISTS idx_organic_units_is_deleted ON organic_units(is_deleted);",

		// Índices para password_histories (GORM pluraliza automáticamente)
		"CREATE INDEX IF NOT EXISTS idx_password_histories_user_id ON password_histories(user_id);",
		"CREATE INDEX IF NOT EXISTS idx_password_histories_is_deleted ON password_histories(is_deleted);",
	}

	for _, index := range indexes {
		err := database.DB.Exec(index).Error
		if err != nil {
			log.Printf("Error al crear índice: %v", err)
		}
	}

	log.Println("Migración de índices completada")
}