package dto

import "time"

type UbigeoDTO struct {
	ID         string    `json:"id"`
	UbigeoCode string    `json:"ubigeo_code"`
	IneiCode   string    `json:"inei_code"`
	Department string    `json:"department"`
	Province   string    `json:"province"`
	District   string    `json:"district"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UbigeosListResponse struct {
	Data     []UbigeoDTO `json:"data"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

type CreateUbigeoRequest struct {
	UbigeoCode string `json:"ubigeo_code" validate:"required,min=6,max=10"`
	IneiCode   string `json:"inei_code" validate:"required,max=10"`
	Department string `json:"department" validate:"required,min=2,max=100"`
	Province   string `json:"province" validate:"required,min=2,max=100"`
	District   string `json:"district" validate:"required,min=2,max=100"`
}

type UpdateUbigeoRequest struct {
	UbigeoCode *string `json:"ubigeo_code" validate:"omitempty,min=6,max=10"`
	IneiCode   *string `json:"inei_code" validate:"omitempty,max=10"`
	Department *string `json:"department" validate:"omitempty,min=2,max=100"`
	Province   *string `json:"province" validate:"omitempty,min=2,max=100"`
	District   *string `json:"district" validate:"omitempty,min=2,max=100"`
}

type UbigeoStatsResponse struct {
	TotalUbigeos     int64 `json:"total_ubigeos"`
	TotalDepartments int64 `json:"total_departments"`
	TotalProvinces   int64 `json:"total_provinces"`
	TotalDistricts   int64 `json:"total_districts"`
}

// Nuevos DTOs para selects
type DepartmentDTO struct {
	Name string `json:"name"`
}

type ProvinceDTO struct {
	Name       string `json:"name"`
	Department string `json:"department"`
}

type DistrictDTO struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Province   string `json:"province"`
	Department string `json:"department"`
	UbigeoCode string `json:"ubigeo_code"`
}