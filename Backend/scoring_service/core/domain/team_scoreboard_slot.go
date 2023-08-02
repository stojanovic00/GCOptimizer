package domain

import "github.com/google/uuid"

type TeamScoreboardSlot struct {
	ID                   uuid.UUID
	Place                int
	SportsOrganizationID uuid.UUID
	SportsOrganization   SportsOrganization
	TeamNumber           int
	Scores               []Score               `gorm:"-"`
	ApparatusTotalScores map[Apparatus]float32 `gorm:"serializer:json"`
	TotalScore           float32
	TeamScoreboardID     uuid.UUID
	TeamScoreboard       TeamScoreboard
}
