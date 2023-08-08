package domain

import "github.com/google/uuid"

type ScheduleSlot struct {
	ID                   uuid.UUID
	SessionID            uuid.UUID
	Session              Session
	StartingApparatus    Apparatus
	CompetingApparatuses []Apparatus `gorm:"serializer:json"`
	ScoredApparatuses    []Apparatus `gorm:"serializer:json"`
	ContestantID         uuid.UUID
	Contestant           Contestant
	Position             int
}

func (s *ScheduleSlot) CompetesApparatus(apparatus Apparatus) bool {
	found := false

	for _, app := range s.CompetingApparatuses {
		if app == apparatus {
			found = true
			break
		}
	}

	return found
}
