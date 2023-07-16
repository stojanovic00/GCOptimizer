package repo

import (
	"application_service/core/domain"
	"github.com/google/uuid"
)

type SportsOrganisationRepo interface {
	Create(organisation *domain.SportsOrganization) (uuid.UUID, error)
	GetByEmail(email string) (*domain.SportsOrganization, error)
}
