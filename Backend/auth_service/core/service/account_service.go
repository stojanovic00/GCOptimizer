package service

import (
	"auth_service/core/domain"
	"auth_service/core/repo"
	"auth_service/errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AccountService struct {
	accountRepository repo.IAccountRepository
}

func NewAccountService(accountRepository repo.IAccountRepository) *AccountService {
	return &AccountService{accountRepository: accountRepository}
}

func (s *AccountService) Create(account *domain.Account) (uuid.UUID, error) {
	var err error
	account.Password, err = HashPassword(account.Password)
	if err != nil {
		return uuid.UUID{}, err
	}
	return s.accountRepository.Create(account)
}

func (s *AccountService) Login(email, password string) (domain.Account, error) {
	acc, err := s.accountRepository.GetByEmail(email)
	//There shouldn't be known which one is incorrect because of security reasons
	if err != nil {
		return domain.Account{}, &errors.ErrBadCredentials{}
	}

	if !passwordMatches(acc.Password, password) {
		return domain.Account{}, &errors.ErrBadCredentials{}
	}

	return acc, nil
}

func (s *AccountService) HasPermission(role, permission string) (bool, error) {
	return s.accountRepository.HasPermission(role, permission)
}

func HashPassword(password string) (string, error) {
	// Generate a bcrypt hash of the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Convert the hashed password to a string and return it
	return string(hashedPassword), nil
}

func passwordMatches(hashedPassword, password string) bool {
	// Compare the provided password with the stored hashed password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
