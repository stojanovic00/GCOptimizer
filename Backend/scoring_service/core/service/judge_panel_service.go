package service

import (
	auth_pb "common/proto/auth/generated"
	"context"
	"github.com/google/uuid"
	"scoring_service/core/domain"
	"scoring_service/core/repo"
)

type JudgePanelService struct {
	judgePanelRepo repo.JudgePanelRepo
	authClient     auth_pb.AuthServiceClient
}

func NewJudgePanelService(judgePanelRepo repo.JudgePanelRepo, authClient auth_pb.AuthServiceClient) *JudgePanelService {
	return &JudgePanelService{judgePanelRepo: judgePanelRepo, authClient: authClient}
}

func (s *JudgePanelService) GetApparatusesWithoutPanel(compId uuid.UUID) ([]domain.Apparatus, error) {
	return s.judgePanelRepo.GetApparatusesWithoutPanel(compId)
}
func (s *JudgePanelService) CreateJudgingPanelsForApparatus(apparatus domain.Apparatus, compId uuid.UUID) (uuid.UUID, uuid.UUID, error) {
	return s.judgePanelRepo.CreateJudgingPanelsForApparatus(apparatus, compId)
}
func (s *JudgePanelService) AssignJudge(judge *domain.Judge, panelId uuid.UUID) error {
	panelType, err := s.judgePanelRepo.AssignJudge(judge, panelId)
	if err != nil {
		return err
	}

	//Make account for judge
	var role auth_pb.Role

	switch panelType {
	case domain.DPanel:
		role = auth_pb.Role{Name: "d_judge"}
	case domain.EPanel:
		role = auth_pb.Role{Name: "e_judge"}
	}
	//Password is assigned in auth service for judges
	_, err = s.authClient.Create(context.Background(), &auth_pb.Account{
		Email: judge.Email,
		Role:  &role,
	})
	if err != nil {
		return err
	}

	return nil
}
func (s *JudgePanelService) GetAssignedJudges(competitionId uuid.UUID) ([]domain.Judge, error) {
	return s.judgePanelRepo.GetAssignedJudges(competitionId)
}
func (s *JudgePanelService) AssignScoreCalculationMethod(scoreCalcMethod *domain.ScoreCalculationMethod, panelId uuid.UUID) error {
	return s.judgePanelRepo.AssignScoreCalculationMethod(scoreCalcMethod, panelId)
}

func (s *JudgePanelService) DeleteAllJudgeAccounts(competitionId uuid.UUID) error {
	judges, err := s.GetAssignedJudges(competitionId)
	if err != nil {
		return err
	}

	var emails []string
	for _, judge := range judges {
		emails = append(emails, judge.Email)
	}

	_, err = s.authClient.DeleteAccounts(context.Background(), &auth_pb.EmailList{Emails: emails})
	if err != nil {
		return err
	}

	return nil
}
