package domain

import "github.com/google/uuid"

type ContestantApplication struct {
	ID                     uuid.UUID
	TeamNumber             int
	CompetitionID          uuid.UUID
	Competition            Competition
	ContestantID           uuid.UUID
	Contestant             Contestant
	AgeCategoryID          uuid.UUID
	AgeCategory            AgeCategory
	ApparatusAnnouncements []ApparatusAnnouncement
}
