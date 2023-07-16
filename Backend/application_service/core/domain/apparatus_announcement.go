package domain

import "github.com/google/uuid"

type ApparatusAnnouncement struct {
	ID                      uuid.UUID
	Apparatus               Apparatus
	ContestantApplicationID uuid.UUID
	ContestantApplication   ContestantApplication
}
