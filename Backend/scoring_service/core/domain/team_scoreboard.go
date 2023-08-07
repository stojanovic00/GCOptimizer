package domain

import "github.com/google/uuid"

type TeamScoreboard struct {
	ID            uuid.UUID
	CompetitionID uuid.UUID
	Competition   Competition
	AgeCategory   string
	Apparatuses   []Apparatus `gorm:"serializer:json"`
	Slots         []TeamScoreboardSlot
}
