package repo

import (
	"gorm.io/gorm"
	"scoring_service/core/domain"
)

type ScheduleRepoPg struct {
	dbClient *gorm.DB
}

func NewScheduleRepoPg(dbClient *gorm.DB) *ScheduleRepoPg {
	return &ScheduleRepoPg{dbClient: dbClient}
}

func (r *ScheduleRepoPg) Save(schedule *domain.Schedule) error {
	result := r.dbClient.Create(schedule)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
