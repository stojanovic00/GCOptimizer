package domain

import "github.com/google/uuid"

type Judge struct {
	ID                   uuid.UUID
	FullName             string
	LicenceType          LicenceType
	LicenceName          string
	SportsOrganizationID uuid.UUID
	SportsOrganization   SportsOrganization
	PanelID              uuid.UUID
	Panel                Panel
}
