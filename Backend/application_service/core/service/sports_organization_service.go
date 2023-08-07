package service

import (
	"application_service/core/domain"
	"application_service/core/repo"
	"github.com/google/uuid"
)

type SportsOrganisationService struct {
	soRepo repo.SportsOrganisationRepo
}

func NewSportsOrganisationService(soRepo repo.SportsOrganisationRepo) *SportsOrganisationService {
	return &SportsOrganisationService{soRepo: soRepo}
}

func (s *SportsOrganisationService) Create(organisation *domain.SportsOrganization) (uuid.UUID, error) {
	return s.soRepo.Create(organisation)
}
func (s *SportsOrganisationService) GetByEmail(email string) (*domain.SportsOrganization, error) {
	return s.soRepo.GetByEmail(email)
}
func (s *SportsOrganisationService) GetByName(name string) (*domain.SportsOrganization, error) {
	return s.soRepo.GetByName(name)
}
