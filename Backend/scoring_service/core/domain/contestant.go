package domain

import (
	"github.com/google/uuid"
)

type Contestant struct {
	ID                   uuid.UUID
	CompetingId          int32
	FullName             string
	SportsOrganizationID uuid.UUID
	SportsOrganization   SportsOrganization
	CompetingApparatuses []Apparatus `gorm:"type:json;scan:scanApparatuses"`
	TeamNumber           int32
	AgeCategoryID        uuid.UUID
	AgeCategory          AgeCategory
}
