package repo

import (
	"github.com/google/uuid"
	"scoring_service/core/domain"
)

type ScheduleRepo interface {
	Save(schedule *domain.Schedule) error
	CompetitionExists(compId uuid.UUID) (bool, error)
}
