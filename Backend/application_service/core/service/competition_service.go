package service

import (
	"application_service/core/domain"
	"application_service/core/repo"
	"github.com/google/uuid"
)

type CompetitionService struct {
	compRepo repo.CompetitionRepo
	soRepo   repo.SportsOrganisationRepo
	dmRepo   repo.DelegationMemberRepo
}

func NewCompetitionService(compRepo repo.CompetitionRepo, soRepo repo.SportsOrganisationRepo, dmRepo repo.DelegationMemberRepo) *CompetitionService {
	return &CompetitionService{compRepo: compRepo, soRepo: soRepo, dmRepo: dmRepo}
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

func (s *CompetitionService) AddAgeCategory(ageCat *domain.AgeCategory, compId uuid.UUID) (uuid.UUID, error) {
	comp, err := s.compRepo.GetById(compId)
	if err != nil {
		return uuid.UUID{}, err
	}

	ageCat.Competition = *comp
	return s.compRepo.AddAgeCategory(ageCat)
}
func (s *CompetitionService) AddDelegationMemberProposition(prop *domain.DelegationMemberProposition, compId uuid.UUID) (uuid.UUID, error) {
	comp, err := s.compRepo.GetById(compId)
	if err != nil {
		return uuid.UUID{}, err
	}
	prop.Competition = *comp

	position, err := s.dmRepo.GetPositionByName(prop.Position.Name)
	if err != nil {
		return uuid.UUID{}, err
	}
	prop.Position = *position

	return s.compRepo.AddDelegationMemberProposition(prop)
}
