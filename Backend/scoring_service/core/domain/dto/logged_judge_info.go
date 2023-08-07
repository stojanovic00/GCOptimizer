package dto

import (
	"github.com/google/uuid"
	"scoring_service/core/domain"
)

type JudgeJudgingInfo struct {
	Judge             domain.Judge
	CompetitionId     uuid.UUID
	Apparatus         domain.Apparatus
	JudgingPanelType  domain.JudgingPanelType
	CalculationMethod domain.ScoreCalculationMethod
}
