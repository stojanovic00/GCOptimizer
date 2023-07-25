package domain

import "github.com/google/uuid"

type ScheduleSlot struct {
	ID                uuid.UUID
	SessionID         uuid.UUID
	Session           Session
	StartingApparatus Apparatus
	ScoredApparatuses []Apparatus `gorm:"type:json;scan:scanApparatuses"`
	ContestantID      uuid.UUID
	Contestant        Contestant
}
