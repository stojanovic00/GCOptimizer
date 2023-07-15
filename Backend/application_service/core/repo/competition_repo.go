package repo

import (
	"application_service/core/domain"
	"github.com/google/uuid"
)

type CompetitionRepo interface {
	Create(competition *domain.Competition) (uuid.UUID, error)
	GetById(id uuid.UUID) (*domain.Competition, error)
	GetAll() ([]*domain.Competition, error)
}
