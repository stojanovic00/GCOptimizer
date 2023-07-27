package service

import (
	"auth_service/core/domain"
	"auth_service/core/repo"
	"auth_service/errors"
	"crypto/rand"
	"fmt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"math/big"
	"strings"
)

type AccountService struct {
	accountRepository repo.AccountRepo
	mailService       *MailService
}

func NewAccountService(accountRepository repo.AccountRepo, mailService *MailService) *AccountService {
	return &AccountService{accountRepository: accountRepository, mailService: mailService}
}

func (s *AccountService) Create(account *domain.Account) (uuid.UUID, error) {
	var err error
	//For judges registration is automatic
	plaintTextPassword := account.Password

	isJudge := strings.Contains(strings.ToUpper(account.Role.Name), "JUDGE")
	if isJudge {
		plaintTextPassword, err = GenerateRandomPassword(10)
		if err != nil {
			return uuid.UUID{}, err
		}
	}

	account.Password, err = hashPassword(plaintTextPassword)
	if err != nil {
		return uuid.UUID{}, err
	}

	id, err := s.accountRepository.Create(account)
	if err != nil {
		return uuid.UUID{}, err
	}

	//If judge send email with password

	if isJudge {
		subject := "GCO new registration"
		body := fmt.Sprintf("Your password for private GCO platform is: %s", plaintTextPassword)

		err := s.mailService.sendEmail(account.Email, subject, body)
		if err != nil {
			return uuid.UUID{}, err
		}
	}

	return id, nil
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

func hashPassword(password string) (string, error) {
	// Generate a bcrypt hash of the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Convert the hashed password to a string and return it
	return string(hashedPassword), nil
}

func GenerateRandomPassword(length int) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$%^&*()-_=+,.?/:;{}[]<>"

	randomPassword := make([]byte, length)
	for i := range randomPassword {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		randomPassword[i] = charset[n.Int64()]
	}

	return string(randomPassword), nil
}

func passwordMatches(hashedPassword, password string) bool {
	// Compare the provided password with the stored hashed password
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false
	}
	return true
}
