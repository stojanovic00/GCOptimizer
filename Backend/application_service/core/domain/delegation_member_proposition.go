package domain

import "github.com/google/uuid"

type DelegationMemberProposition struct {
	ID         uuid.UUID
	PositionID uuid.UUID
	Position   DelegationMemberPosition
	MinNumber  string
	MaxNumber  string
}
