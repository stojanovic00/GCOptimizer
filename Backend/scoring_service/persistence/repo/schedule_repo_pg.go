package repo

import (
	"github.com/google/uuid"
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

func (r *ScheduleRepoPg) CompetitionExists(compId uuid.UUID) (bool, error) {
	var count int64
	result := r.dbClient.
		Model(domain.Competition{}).
		Where("id = ?", compId).
		Count(&count)
	if result.Error != nil {
		return false, result.Error
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
