package repo

import (
	"application_service/core/domain"
	"github.com/google/uuid"
)

type ApplicationRepo interface {
	CreateJudgeApplication(app *domain.JudgeApplication) (uuid.UUID, error)
	GetAllJudgeApplications(compId uuid.UUID) ([]*domain.JudgeApplication, error)
	CreateContestantApplication(app *domain.ContestantApplication) (uuid.UUID, error)
	GetAllContestantApplications(compId uuid.UUID) ([]*domain.ContestantApplication, error)
}
