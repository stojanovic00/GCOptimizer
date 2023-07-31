package repo

import (
	"github.com/google/uuid"
	"scoring_service/core/domain"
	"scoring_service/core/domain/dto"
)

type ScoringRepo interface {
	GetJudgeJudgingInfo(email string) (*dto.JudgeJudgingInfo, error)
	GetScheduleByCompetitionId(competitionId uuid.UUID) (*domain.Schedule, error)
	GetCurrentSession(competitionId uuid.UUID) (*domain.Session, error)
	GetSlotsWithStartingApparatus(competitionId uuid.UUID, sessionNumber int32, apparatus domain.Apparatus) ([]domain.ScheduleSlot, error)
	SubmitTempScore(tempScore *domain.TempScore) error
	GetContestantsTempScores(competitionId, contestantId uuid.UUID, apparatus domain.Apparatus) ([]domain.TempScore, error)
	SaveScore(score *domain.Score) error
	SubmitScore(score *domain.Score) error
	FinishRotation(competitionId uuid.UUID) error
	FinishSession(competitionId uuid.UUID) error
	IsRotationFinished(competitionId uuid.UUID) (bool, error)
	IsSessionFinished(competitionId uuid.UUID) (bool, error)
	IsCompetitionFinished(competitionId uuid.UUID) (bool, error)
	GetScore(competitionId, contestantId uuid.UUID, apparatus domain.Apparatus) (*domain.Score, error)
}
