package repo

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"scoring_service/core/domain"
)

type JudgePanelRepoPg struct {
	dbClient *gorm.DB
}

func NewJudgePanelRepoPg(dbClient *gorm.DB) *JudgePanelRepoPg {
	return &JudgePanelRepoPg{dbClient: dbClient}
}

func (r *JudgePanelRepoPg) GetApparatusesWithoutPanel(compId uuid.UUID) ([]domain.Apparatus, error) {
	var schedule domain.Schedule

	result := r.dbClient.Where("competition_id = ?", compId).First(&schedule)
	if result.Error != nil {
		return nil, result.Error
	}

	var appsWithPanel []domain.Apparatus
	result = r.dbClient.Model(&domain.Panel{}).Distinct("apparatus").Where("competition_id = ?", compId).Find(&appsWithPanel)
	if result.Error != nil {
		return nil, result.Error
	}

	return difference(schedule.ApparatusOrder, appsWithPanel), nil
}

// difference returns the elements in `a` that aren't in `b`.
func difference(a, b []domain.Apparatus) []domain.Apparatus {
	var diff []domain.Apparatus

	for _, app1 := range a {
		found := false
		for _, app2 := range b {
			if app1 == app2 {
				found = true
				break
			}
		}
		if !found {
			diff = append(diff, app1)
		}
	}

	return diff
}

func (r *JudgePanelRepoPg) CreateJudgingPanelsForApparatus(apparatus domain.Apparatus, compId uuid.UUID) (uuid.UUID, uuid.UUID, error) {
	var competition domain.Competition

	result := r.dbClient.Where("id = ?", compId).First(&competition)
	if result.Error != nil {
		return uuid.UUID{}, uuid.UUID{}, result.Error
	}

	//D panel
	dPanel := &domain.Panel{
		ID:          uuid.New(),
		Type:        domain.DPanel,
		Apparatus:   apparatus,
		Competition: competition,
	}

	result = r.dbClient.Create(dPanel)
	if result.Error != nil {
		return uuid.UUID{}, uuid.UUID{}, result.Error
	}

	//E panel
	ePanel := &domain.Panel{
		ID:          uuid.New(),
		Type:        domain.EPanel,
		Apparatus:   apparatus,
		Competition: competition,
	}

	result = r.dbClient.Create(ePanel)
	if result.Error != nil {
		return uuid.UUID{}, uuid.UUID{}, result.Error
	}
	return dPanel.ID, ePanel.ID, nil
}
func (r *JudgePanelRepoPg) AssignJudge(judge *domain.Judge, panelId uuid.UUID) (domain.JudgingPanelType, error) {
	var panel domain.Panel

	result := r.dbClient.Where("id = ?", panelId).First(&panel)
	if result.Error != nil {
		return 0, result.Error
	}

	panel.Judges = append(panel.Judges, *judge)

	result = r.dbClient.Save(&panel)
	if result.Error != nil {
		return 0, result.Error
	}

	return panel.Type, nil
}
func (r *JudgePanelRepoPg) GetAssignedJudges(competitionId uuid.UUID) ([]domain.Judge, error) {
	var judges []domain.Judge

	var panels []domain.Panel
	result := r.dbClient.Where("competition_id = ?", competitionId).
		Preload("Judges.SportsOrganization.Address").Find(&panels)
	if result.Error != nil {
		return nil, result.Error
	}

	for _, panel := range panels {
		judges = append(judges, panel.Judges...)
	}

	return judges, nil
}
func (r *JudgePanelRepoPg) AssignScoreCalculationMethod(scoreCalcMethod *domain.ScoreCalculationMethod, panelId uuid.UUID) error {
	var panel domain.Panel

	result := r.dbClient.Where("id = ?", panelId).First(&panel)
	if result.Error != nil {
		return result.Error
	}

	if scoreCalcMethod.ID == nil {
		id := uuid.New()
		scoreCalcMethod.ID = &id
	}

	panel.ScoreCalculationMethod = *scoreCalcMethod

	result = r.dbClient.Save(&panel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
