package repo

import (
	"auth_service/core/domain"
	"auth_service/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"strings"
)

type AccountRepoPg struct {
	dbClient *gorm.DB
}

func NewAccountRepositoryPg(dbClient *gorm.DB) *AccountRepoPg {
	return &AccountRepoPg{dbClient: dbClient}
}

func (r *AccountRepoPg) Create(account *domain.Account) (uuid.UUID, error) {
	_, err := r.GetByEmail(account.Email)
	if err == nil {
		return uuid.Nil, errors.ErrEmailTaken{}
	}

	if account.ID == uuid.Nil {
		account.ID, _ = uuid.NewUUID()
	}

	//Get role
	var role domain.Role
	result := r.dbClient.Where("name = ?", strings.ToUpper(account.Role.Name)).First(&role)
	if result.Error != nil {
		return uuid.UUID{}, result.Error
	}

	account.Role = role
	result = r.dbClient.Create(account)
	if result.Error != nil {
		return uuid.UUID{}, result.Error
	}

	return account.ID, nil
}

func (r *AccountRepoPg) GetByEmail(email string) (domain.Account, error) {
	var account domain.Account

	result := r.dbClient.Where("email = ?", email).Preload("Role").First(&account)
	if result.Error != nil {
		return domain.Account{}, result.Error
	}
	return account, nil
}

func (r *AccountRepoPg) HasPermission(roleName, permissionName string) (bool, error) {
	//Get role
	var role domain.Role
	result := r.dbClient.Where("name = ?", strings.ToUpper(roleName)).First(&role)
	if result.Error != nil {
		return false, result.Error
	}

	//Get permission
	var permission domain.Permission
	result = r.dbClient.Where("name = ?", permissionName).First(&permission)
	if result.Error != nil {
		return false, result.Error
	}

	var count int64 = 0
	r.dbClient.Table("role_permission").Where("role_id = ? AND permission_id = ?", role.ID, permission.ID).Count(&count)

	if count > 0 {
		return true, nil
	} else {
		return false, nil
	}
}
