package dto

// ListSessionsParams define filtros y paginación para listar sesiones.
type ListSessionsParamsDTO struct {
	// Estado de la sesión (por ejemplo, "active", "revoked").
	Status *string `json:"status,omitempty" query:"status"`

	// Indica si la sesión está activa.
	IsActive *bool `json:"isActive,omitempty" query:"isActive"`
}
