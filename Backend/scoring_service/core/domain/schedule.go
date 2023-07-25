package domain

import "github.com/google/uuid"

type Schedule struct {
	ID            uuid.UUID
	CompetitionID uuid.UUID
	Competition   Competition
	Sessions      []Session
}
