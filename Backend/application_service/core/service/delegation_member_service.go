package service

import (
	"application_service/core/domain"
	"application_service/core/repo"
	"github.com/google/uuid"
)

type DelegationMemberService struct {
	dmRepo repo.DelegationMemberRepo
	soRepo repo.SportsOrganisationRepo
}

func NewDelegationMemberService(dmRepo repo.DelegationMemberRepo, soRepo repo.SportsOrganisationRepo) *DelegationMemberService {
	return &DelegationMemberService{dmRepo: dmRepo, soRepo: soRepo}
}

func (s *DelegationMemberService) RegisterJudge(judge *domain.Judge, soEmail string) (uuid.UUID, error) {
	sportsOrg, err := s.soRepo.GetByEmail(soEmail)
	if err != nil {
		return uuid.UUID{}, err
	}
	judge.SportsOrganization = *sportsOrg

	return s.dmRepo.RegisterJudge(judge)
}

func (s *DelegationMemberService) GetSportsOrganisationJudges(soEmail string) ([]*domain.Judge, error) {
	sportsOrg, err := s.soRepo.GetByEmail(soEmail)
	if err != nil {
		return nil, err
	}

	return s.dmRepo.GetSportsOrganisationJudges(sportsOrg.ID)
}
func (s *DelegationMemberService) RegisterContestant(contestant domain.Contestant) (uuid.UUID, error) {
	return uuid.UUID{}, nil
}
func (s *DelegationMemberService) GetSportsOrganisationContestants(soEmail string) ([]domain.Contestant, error) {
	return nil, nil
}
