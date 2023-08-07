package domain

import (
	"github.com/google/uuid"
)

type AgeCategory struct {
	ID            uuid.UUID
	Name          string
	MinAge        int
	MaxAge        int
	CompetitionID uuid.UUID
	Competition   Competition
}
