package domain

import "github.com/google/uuid"

type Session struct {
	ID              uuid.UUID
	Number          int32
	CurrentRotation int32
	ScheduleSlots   []ScheduleSlot
	ScheduleID      uuid.UUID
	Schedule        Schedule
}
