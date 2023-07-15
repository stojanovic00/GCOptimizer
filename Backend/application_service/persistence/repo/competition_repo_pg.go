package repo

import (
	"application_service/core/domain"
	"application_service/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CompetitionRepoPg struct {
	dbClient *gorm.DB
}

func NewCompetitionRepoPg(dbClient *gorm.DB) *CompetitionRepoPg {
	return &CompetitionRepoPg{dbClient: dbClient}
}

func (r *CompetitionRepoPg) Create(competition *domain.Competition) (uuid.UUID, error) {
	if competition.ID == uuid.Nil {
		competition.ID, _ = uuid.NewUUID()
	}

	if competition.Address.ID == uuid.Nil {
		competition.Address.ID, _ = uuid.NewUUID()
	}

	if competition.TeamComposition.ID == uuid.Nil {
		competition.TeamComposition.ID, _ = uuid.NewUUID()
	}
	result := r.dbClient.Create(competition)
	if result.Error != nil {
		return uuid.UUID{}, result.Error
	}

	return competition.ID, nil
}

func (r *CompetitionRepoPg) GetById(id uuid.UUID) (*domain.Competition, error) {
	var comp domain.Competition

	//TODO add more preloads
	result := r.dbClient.Where("id = ?", id).Preload("Address").Preload("TeamComposition").First(&comp)
	if result.Error != nil {
		return &domain.Competition{}, result.Error
	}

	if &comp == nil {
		return &domain.Competition{}, errors.ErrNotFound{Message: "Competition with given id not found"}
	}
	return &comp, nil
}

func (r *CompetitionRepoPg) GetAll() ([]*domain.Competition, error) {
	var competitions []*domain.Competition
	err := r.dbClient.Preload("Address").Preload("TeamComposition").Find(&competitions).Error
	if err != nil {
		return nil, nil
	}

	return competitions, nil
}
