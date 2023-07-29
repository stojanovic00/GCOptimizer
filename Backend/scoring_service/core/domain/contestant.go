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

func (c *Contestant) CompetesApparatus(apparatus Apparatus) bool {
	found := false

	for _, app := range c.CompetingApparatuses {
		if app == apparatus {
			found = true
			break
		}
	}

	return found
}
