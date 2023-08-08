package domain

import (
	"github.com/google/uuid"
)

type Contestant struct {
	// TODO Migrate competition specific fields
	ID                   uuid.UUID
	CompetingId          int32
	FullName             string
	SportsOrganizationID uuid.UUID
	SportsOrganization   SportsOrganization
	TeamNumber           int32
	AgeCategory          string
}
