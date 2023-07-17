package repo

import (
	"application_service/core/domain"
	"application_service/errors"
	errors2 "errors"
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

	result := r.dbClient.
		Where("id = ?", id).
		Preload("Address").
		Preload("Organizer.Address").
		Preload("TeamComposition").
		Preload("DelegationMemberPropositions.Position").
		Preload("AgeCategories").
		First(&comp)
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
	err := r.dbClient.
		Preload("Address").
		Preload("Organizer").
		Find(&competitions).Error
	if err != nil {
		return nil, nil
	}

	return competitions, nil
}

func (r *CompetitionRepoPg) AddAgeCategory(ageCat *domain.AgeCategory) (uuid.UUID, error) {
	if ageCat.ID == uuid.Nil {
		ageCat.ID, _ = uuid.NewUUID()
	}

	result := r.dbClient.Create(ageCat)
	if result.Error != nil {
		return uuid.UUID{}, result.Error
	}

	return ageCat.ID, nil
}

func (r *CompetitionRepoPg) GetAgeCategoryById(id uuid.UUID) (*domain.AgeCategory, error) {
	var cat domain.AgeCategory

	result := r.dbClient.Where("id = ?", id).First(&cat)
	if result.Error != nil {
		return &domain.AgeCategory{}, result.Error
	}

	if &cat == nil {
		return &domain.AgeCategory{}, errors.ErrNotFound{Message: "Age category with given id not found"}
	}
	return &cat, nil
}
func (r *CompetitionRepoPg) DelegationMemberPropositionPositionWithSameName(name string, compId uuid.UUID) (bool, error) {
	var count int64

	result := r.dbClient.Model(&domain.DelegationMemberProposition{}).
		Joins("LEFT JOIN delegation_member_positions"+
			" ON delegation_member_propositions.position_id = delegation_member_positions.id").
		Where("delegation_member_positions.name = ? and delegation_member_propositions.competition_id = ?", name, compId).
		Count(&count)

	if result.Error != nil {
		return false, result.Error
	}

	return count > 0, nil
}

func (r *CompetitionRepoPg) AddDelegationMemberProposition(prop *domain.DelegationMemberProposition) (uuid.UUID, error) {
	if prop.ID == uuid.Nil {
		prop.ID, _ = uuid.NewUUID()
	}

	sameNamePosition, err := r.DelegationMemberPropositionPositionWithSameName(prop.Position.Name, prop.Competition.ID)
	if err != nil {
		return uuid.UUID{}, err
	}
	if sameNamePosition {
		return uuid.UUID{}, errors2.New("Already given proposition for given position")
	}

	result := r.dbClient.Create(prop)
	if result.Error != nil {
		return uuid.UUID{}, result.Error
	}

	return prop.ID, nil
}
