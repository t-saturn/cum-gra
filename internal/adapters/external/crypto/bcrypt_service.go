package crypto

import (
	"golang.org/x/crypto/bcrypt"
)

type BcryptService struct{}

func NewBcryptService() *BcryptService {
	return &BcryptService{}
}

func (s *BcryptService) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func (s *BcryptService) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
