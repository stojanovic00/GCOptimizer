package service

import (
	"application_service/core/domain"
	"application_service/core/repo"
	"github.com/google/uuid"
)

type ApplicationService struct {
	appRepo  repo.ApplicationRepo
	compRepo repo.CompetitionRepo
	dmRepo   repo.DelegationMemberRepo
}

func NewApplicationService(appRepo repo.ApplicationRepo, compRepo repo.CompetitionRepo, dmRepo repo.DelegationMemberRepo) *ApplicationService {
	return &ApplicationService{appRepo: appRepo, compRepo: compRepo, dmRepo: dmRepo}
}
func (s *ApplicationService) CreateJudgeApplication(app *domain.JudgeApplication, compId uuid.UUID) (uuid.UUID, error) {
	comp, err := s.compRepo.GetById(compId)
	if err != nil {
		return uuid.UUID{}, err
	}
	app.Competition = *comp

	judge, err := s.dmRepo.GetJudgeById(app.JudgeID)
	if err != nil {
		return uuid.UUID{}, err
	}
	app.Judge = *judge

	return s.appRepo.CreateJudgeApplication(app)
}

func (s *ApplicationService) GetAllJudgeApplications(compId uuid.UUID) ([]*domain.JudgeApplication, error) {
	return s.appRepo.GetAllJudgeApplications(compId)
}
func (s *ApplicationService) CreateContestantApplication(app *domain.ContestantApplication) (uuid.UUID, error) {
	comp, err := s.compRepo.GetById(app.CompetitionID)
	if err != nil {
		return uuid.UUID{}, err
	}
	app.Competition = *comp

	contestant, err := s.dmRepo.GetContestantById(app.ContestantID)
	if err != nil {
		return uuid.UUID{}, err
	}
	app.Contestant = *contestant

	ageCategory, err := s.compRepo.GetAgeCategoryById(app.AgeCategoryID)
	if err != nil {
		return uuid.UUID{}, err
	}
	app.AgeCategory = *ageCategory

	for idx := range app.ApparatusAnnouncements {
		app.ApparatusAnnouncements[idx].ID, _ = uuid.NewUUID()
	}

	return s.appRepo.CreateContestantApplication(app)
}
func (s *ApplicationService) GetAllContestantApplications(compId uuid.UUID) ([]*domain.ContestantApplication, error) {
	return s.appRepo.GetAllContestantApplications(compId)
}
