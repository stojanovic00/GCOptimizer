package repo

import (
	"github.com/google/uuid"
	"scoring_service/core/domain"
)

type JudgePanelRepo interface {
	GetApparatusesWithoutPanel(compId uuid.UUID) ([]domain.Apparatus, error)
	CreateJudgingPanelsForApparatus(apparatus domain.Apparatus, compId uuid.UUID) (uuid.UUID, uuid.UUID, error)
	AssignJudge(judge *domain.Judge, panelId uuid.UUID) (domain.JudgingPanelType, error)
	GetAssignedJudges(competitionId uuid.UUID) ([]domain.Judge, error)
	AssignScoreCalculationMethod(scoreCalcMethod *domain.ScoreCalculationMethod, panelId uuid.UUID) error
	GetJudgePanelByCompetitionIdAndApparatus(competitionId uuid.UUID, apparatus domain.Apparatus, panelType domain.JudgingPanelType) (*domain.Panel, error)
}
