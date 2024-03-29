package domain

import "github.com/google/uuid"

type SportsOrganization struct {
	ID                             uuid.UUID
	Name                           string
	Email                          string
	PhoneNumber                    string
	ContactPersonFullName          string
	CompetitionOrganisingPrivilege bool
	AddressID                      uuid.UUID
	Address                        Address
}
