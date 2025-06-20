package services

type HashService interface {
	HashPassword(password string) (string, error)
	CheckPasswordHash(password, hash string) bool
}
