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
	CompetingApparatuses []Apparatus `gorm:"serializer:json"`
	TeamNumber           int32
	AgeCategory          string
}
