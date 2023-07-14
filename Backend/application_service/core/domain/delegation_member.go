package domain

import "github.com/google/uuid"

type DelegationMember struct {
	ID                   uuid.UUID
	FullName             string
	Email                string
	Gender               Gender
	PositionID           uuid.UUID
	Position             DelegationMemberPosition
	Image                string
	SportsOrganisationID uuid.UUID
	SportsOrganisation   SportsOrganisation
}
