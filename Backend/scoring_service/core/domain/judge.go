package domain

import "github.com/google/uuid"

type Judge struct {
	ID                   uuid.UUID
	FullName             string
	Email                string
	LicenceType          LicenceType
	LicenceName          string
	SportsOrganizationID uuid.UUID
	SportsOrganization   SportsOrganization
}
