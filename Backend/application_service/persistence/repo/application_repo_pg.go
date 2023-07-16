package repo

import (
	"application_service/core/domain"
	"application_service/errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ApplicationRepoPg struct {
	dbClient *gorm.DB
}

func NewApplicationRepoPg(dbClient *gorm.DB) *ApplicationRepoPg {
	return &ApplicationRepoPg{dbClient: dbClient}
}

func (r *ApplicationRepoPg) CreateJudgeApplication(app *domain.JudgeApplication) (uuid.UUID, error) {
	if app.ID == uuid.Nil {
		app.ID, _ = uuid.NewUUID()
	}

	var existingApp domain.JudgeApplication
	result := r.dbClient.Where("competition_id = ? and judge_id = ?", app.Competition.ID, app.Judge.ID).First(&existingApp)
	if existingApp.ID != uuid.Nil {
		return uuid.UUID{}, errors.ErrAlreadyExists{Message: "This judge already applied to this competition"}
	}

	result = r.dbClient.Create(app)
	if result.Error != nil {
		return uuid.UUID{}, result.Error
	}

	return app.ID, nil
}

func (r *ApplicationRepoPg) GetAllJudgeApplications(compId uuid.UUID) ([]*domain.JudgeApplication, error) {
	var applications []*domain.JudgeApplication
	err := r.dbClient.Where("competition_id = ?", compId).Preload("Judge").Find(&applications).Error
	if err != nil {
		return nil, nil
	}

	return applications, nil
}
func (r *ApplicationRepoPg) CreateContestantApplication(app *domain.ContestantApplication) (uuid.UUID, error) {
	if app.ID == uuid.Nil {
		app.ID, _ = uuid.NewUUID()
	}

	var existingApp domain.ContestantApplication
	result := r.dbClient.Where("competition_id = ? and contestant_id = ?", app.Competition.ID, app.Contestant.ID).First(&existingApp)
	if existingApp.ID != uuid.Nil {
		return uuid.UUID{}, errors.ErrAlreadyExists{Message: "This contestant already applied to this competition"}
	}

	result = r.dbClient.Create(app)
	if result.Error != nil {
		return uuid.UUID{}, result.Error
	}

	return app.ID, nil
}

func (r *ApplicationRepoPg) GetAllContestantApplications(compId uuid.UUID) ([]*domain.ContestantApplication, error) {
	var applications []*domain.ContestantApplication
	err := r.dbClient.Where("competition_id = ?", compId).
		Preload("Contestant").
		Preload("AgeCategory").
		Preload("ApparatusAnnouncements").
		Find(&applications).Error
	if err != nil {
		return nil, nil
	}

	return applications, nil
}
