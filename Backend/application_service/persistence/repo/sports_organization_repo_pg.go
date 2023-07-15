package repo

import (
	"application_service/core/domain"
	"application_service/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SportsOrganisationRepoPg struct {
	dbClient *gorm.DB
}

func NewSportsOrganisationRepoPg(dbClient *gorm.DB) *SportsOrganisationRepoPg {
	return &SportsOrganisationRepoPg{dbClient: dbClient}
}

func (r *SportsOrganisationRepoPg) Create(organisation *domain.SportsOrganization) (uuid.UUID, error) {
	_, err := r.GetByEmail(organisation.Email)
	if err == nil {
		return uuid.Nil, errors.ErrEmailTaken{}
	}

	if organisation.ID == uuid.Nil {
		organisation.ID, _ = uuid.NewUUID()
	}

	if organisation.Address.ID == uuid.Nil {
		organisation.Address.ID, _ = uuid.NewUUID()
	}

	result := r.dbClient.Create(organisation)
	if result.Error != nil {
		return uuid.UUID{}, result.Error
	}

	return organisation.ID, nil
}

func (r *SportsOrganisationRepoPg) GetByEmail(email string) (*domain.SportsOrganization, error) {
	var spOrg domain.SportsOrganization

	result := r.dbClient.Where("email = ?", email).Preload("Address").First(&spOrg)
	if result.Error != nil {
		return &domain.SportsOrganization{}, result.Error
	}

	if &spOrg == nil {
		return &domain.SportsOrganization{}, errors.ErrNotFound{Message: "Sports organisation with given email not found"}
	}
	return &spOrg, nil
}
