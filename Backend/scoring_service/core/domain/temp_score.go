package domain

import "github.com/google/uuid"

type TempScore struct {
	ID            uuid.UUID
	Type          ScoreType
	Apparatus     Apparatus
	Value         float32
	ContestantID  uuid.UUID
	Contestant    Contestant
	CompetitionID uuid.UUID
	Competition   Competition
	JudgeID       uuid.UUID
	Judge         Judge
}
