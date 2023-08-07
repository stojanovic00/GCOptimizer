package domain

import (
	"github.com/google/uuid"
)

type Role struct {
	ID          uuid.UUID
	Name        string
	Permissions []Permission `gorm:"many2many:role_permission;"`
}
