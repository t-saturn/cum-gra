package dto

type CreateApplicationDTO struct {
	Name         string   `json:"name" validate:"required"`
	ClientID     string   `json:"client_id" validate:"required"`
	ClientSecret string   `json:"client_secret" validate:"required"`
	Domain       string   `json:"domain" validate:"required"`
	Logo         string   `json:"logo"`
	Description  string   `json:"description"`
	CallbackUrls []string `json:"callback_urls"`
	Scopes       []string `json:"scopes"`
	IsFirstParty bool     `json:"is_first_party"`
	Status       string   `json:"status" validate:"required,oneof=active suspended"`
}

type UpdateApplicationDTO struct {
	Name         string   `json:"name" validate:"required"`
	ClientSecret string   `json:"client_secret" validate:"required"`
	Domain       string   `json:"domain" validate:"required"`
	Logo         string   `json:"logo"`
	Description  string   `json:"description"`
	CallbackUrls []string `json:"callback_urls"`
	Scopes       []string `json:"scopes"`
	IsFirstParty bool     `json:"is_first_party"`
	Status       string   `json:"status" validate:"required,oneof=active suspended"`
}
