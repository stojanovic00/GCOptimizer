package repo

import "scoring_service/core/domain"

type ScheduleRepo interface {
	Save(schedule *domain.Schedule) error
}
