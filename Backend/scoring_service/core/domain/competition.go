package domain

import (
	"github.com/google/uuid"
	"time"
)

type Competition struct {
	ID                uuid.UUID
	Name              string
	StartDate         time.Time
	EndDate           time.Time
	Gender            Gender
	Type              CompetitionType
	Tiebreak          bool
	AddressId         uuid.UUID
	Address           Address
	OrganizerID       uuid.UUID
	Organizer         SportsOrganization
	TeamCompositionID uuid.UUID
	TeamComposition   TeamComposition
	AgeCategories     []string `gorm:"serializer:json"`
}
