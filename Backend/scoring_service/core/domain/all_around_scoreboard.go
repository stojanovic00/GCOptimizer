package domain

import (
	"github.com/google/uuid"
)

type AllAroundScoreboard struct {
	ID            uuid.UUID
	CompetitionID uuid.UUID
	Competition   Competition
	AgeCategory   string
	TieBrake      bool
	Apparatuses   []Apparatus `gorm:"serializer:json"`
	Slots         []AllAroundScoreboardSlot
}
