package domain

import "github.com/google/uuid"

type DelegationMemberApplication struct {
	CompetitionID      uuid.UUID
	Competition        Competition
	DelegationMemberID uuid.UUID
	DelegationMember   DelegationMember
}
