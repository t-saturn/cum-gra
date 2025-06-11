package entities

import "github.com/google/uuid"

type User struct {
	ID        uuid.UUID
	Name      string
	LastName  string
	UserName  string
	Email     string
	Password  string
}
