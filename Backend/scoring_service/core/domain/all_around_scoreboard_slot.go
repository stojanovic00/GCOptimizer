package domain

import "github.com/google/uuid"

type AllAroundScoreboardSlot struct {
	ID                    uuid.UUID
	Place                 int
	ContestantID          uuid.UUID
	Contestant            Contestant
	Scores                []Score `gorm:"-"`
	TotalEScore           float32 `gorm:"-"`
	TotalDScore           float32 `gorm:"-"`
	TotalScore            float32 `gorm:"-"`
	AllAroundScoreboardID uuid.UUID
	AllAroundScoreboard   AllAroundScoreboard
}
