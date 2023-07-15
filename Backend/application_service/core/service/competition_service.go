package service

import (
	"application_service/core/domain"
	"application_service/core/repo"
	"github.com/google/uuid"
)

type CompetitionService struct {
	compRepo repo.CompetitionRepo
	soRepo   repo.SportsOrganisationRepo
}

func NewCompetitionService(compRepo repo.CompetitionRepo, soRepo repo.SportsOrganisationRepo) *CompetitionService {
	return &CompetitionService{compRepo: compRepo, soRepo: soRepo}
}

func (s *CompetitionService) Create(competition *domain.Competition, soEmail string) (uuid.UUID, error) {
	sportsOrg, err := s.soRepo.GetByEmail(soEmail)
	if err != nil {
		return uuid.UUID{}, err
	}

	competition.Organizer = *sportsOrg
	return s.compRepo.Create(competition)
}

func (s *CompetitionService) GetById(id uuid.UUID) (*domain.Competition, error) {
	return s.compRepo.GetById(id)
}

func (s *CompetitionService) GetAll() ([]*domain.Competition, error) {
	return s.compRepo.GetAll()
}
