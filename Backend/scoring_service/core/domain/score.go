package domain

import "github.com/google/uuid"

type Score struct {
	ID            uuid.UUID
	Apparatus     Apparatus
	DScore        float32
	EScore        float32
	TotalScore    float32
	CompetitionID uuid.UUID
	Competition   Competition
	ContestantID  uuid.UUID
	Contestant    Contestant
	Submitted     bool
}
