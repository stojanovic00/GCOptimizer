package repo

import (
	"auth_service/core/domain"
	"github.com/google/uuid"
)

type IAccountRepository interface {
	Create(account *domain.Account) (uuid.UUID, error)
	GetByEmail(email string) (domain.Account, error)
	HasPermission(roleName, permissionName string) (bool, error)
}
