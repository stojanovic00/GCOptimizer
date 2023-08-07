package repo

import (
	"application_service/core/domain"
	"github.com/google/uuid"
)

type DelegationMemberRepo interface {
	RegisterJudge(judge *domain.Judge) (uuid.UUID, error)
	GetSportsOrganisationJudges(soID uuid.UUID) ([]*domain.Judge, error)
	RegisterContestant(contestant *domain.Contestant) (uuid.UUID, error)
	GetSportsOrganisationContestants(soID uuid.UUID) ([]*domain.Contestant, error)
	GetPositionByName(name string) (*domain.DelegationMemberPosition, error)
	GetJudgeByEmail(email string) (*domain.Judge, error)
	GetJudgeById(id uuid.UUID) (*domain.Judge, error)
	GetContestantById(id uuid.UUID) (*domain.Contestant, error)
}
