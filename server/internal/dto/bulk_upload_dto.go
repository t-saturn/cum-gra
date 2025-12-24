package dto

// Respuesta genérica para carga masiva
type BulkUploadResponse struct {
	Created  int               `json:"created"`
	Skipped  int               `json:"skipped"`
	Failed   int               `json:"failed"`
	Errors   []BulkUploadError `json:"errors,omitempty"`
	Details  []map[string]any  `json:"details,omitempty"`
}

type BulkUploadError struct {
	Row     int    `json:"row"`
	Field   string `json:"field,omitempty"`
	Message string `json:"message"`
}

// DTOs para carga masiva de Ubigeos
type BulkUbigeoItem struct {
	UbigeoCode string `json:"ubigeo_code" validate:"required"`
	IneiCode   string `json:"inei_code"`
	Department string `json:"department" validate:"required"`
	Province   string `json:"province" validate:"required"`
	District   string `json:"district" validate:"required"`
}

type BulkUbigeosRequest struct {
	Ubigeos []BulkUbigeoItem `json:"ubigeos" validate:"required,min=1,dive"`
}

// DTOs para carga masiva de Unidades Orgánicas
type BulkOrganicUnitItem struct {
	Name        string  `json:"name" validate:"required"`
	Acronym     string  `json:"acronym" validate:"required"`
	Brand       *string `json:"brand,omitempty"`
	Description *string `json:"description,omitempty"`
	ParentID    *uint   `json:"parent_id,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
	CodDepSGD   *string `json:"cod_dep_sgd,omitempty"`
}

type BulkOrganicUnitsRequest struct {
	OrganicUnits []BulkOrganicUnitItem `json:"organic_units" validate:"required,min=1,dive"`
}

// DTOs para carga masiva de Posiciones Estructurales
type BulkPositionItem struct {
	Name        string  `json:"name" validate:"required"`
	Code        string  `json:"code" validate:"required"`
	Level       *int    `json:"level,omitempty"`
	Description *string `json:"description,omitempty"`
	IsActive    *bool   `json:"is_active,omitempty"`
	CodCarSGD   *string `json:"cod_car_sgd,omitempty"`
}

type BulkPositionsRequest struct {
	Positions []BulkPositionItem `json:"positions" validate:"required,min=1,dive"`
}