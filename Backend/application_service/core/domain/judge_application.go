package domain

import "github.com/google/uuid"

type JudgeApplication struct {
	ID            uuid.UUID
	CompetitionID uuid.UUID
	Competition   Competition
	JudgeID       uuid.UUID
	Judge         Judge
}
