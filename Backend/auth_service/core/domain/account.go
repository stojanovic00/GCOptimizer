package domain

import "github.com/google/uuid"

type Account struct {
	ID       uuid.UUID
	Email    string
	Password string
	RoleID   uuid.UUID
	Role     Role
}
