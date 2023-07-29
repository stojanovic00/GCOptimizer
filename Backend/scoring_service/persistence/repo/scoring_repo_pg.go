package repo

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"scoring_service/core/domain"
	"scoring_service/core/domain/dto"
)

type ScoringRepoPg struct {
	dbClient *gorm.DB
}

func NewScoringRepoPg(dbClient *gorm.DB) *ScoringRepoPg {
	return &ScoringRepoPg{dbClient: dbClient}
}

func (r *ScoringRepoPg) GetJudgeJudgingInfo(email string) (*dto.JudgeJudgingInfo, error) {
	var judge domain.Judge

	result := r.dbClient.Where("email = ?", email).Preload("SportsOrganization.Address").First(&judge)
	if result.Error != nil {
		return nil, result.Error
	}

	var panel domain.Panel
	result = r.dbClient.Model(domain.Panel{}).
		Joins("left join panel_judges pj on id = pj.panel_id").
		Where("judge_id = ?", judge.ID).
		Preload("ScoreCalculationMethod").
		Find(&panel)
	if result.Error != nil {
		return nil, result.Error
	}

	return &dto.JudgeJudgingInfo{
		Judge:             judge,
		CompetitionId:     panel.CompetitionID,
		Apparatus:         panel.Apparatus,
		JudgingPanelType:  panel.Type,
		CalculationMethod: panel.ScoreCalculationMethod,
	}, nil

}
func (r *ScoringRepoPg) GetScheduleByCompetitionId(competitionId uuid.UUID) (*domain.Schedule, error) {
	var schedule domain.Schedule

	result := r.dbClient.Where("competition_id = ?", competitionId).First(&schedule)
	if result.Error != nil {
		return nil, result.Error
	}
	return &schedule, nil
}

func (r *ScoringRepoPg) GetCurrentSession(competitionId uuid.UUID) (*domain.Session, error) {
	schedule, err := r.GetScheduleByCompetitionId(competitionId)
	if err != nil {
		return nil, err
	}

	var session domain.Session
	result := r.dbClient.Where("schedule_id = ? and finished = false", schedule.ID).Order("number asc").First(&session)
	if result.Error != nil {
		return nil, err
	}

	return &session, nil
}

func (r *ScoringRepoPg) GetSlotsWithStartingApparatus(competitionId uuid.UUID, sessionNumber int32, apparatus domain.Apparatus) ([]domain.ScheduleSlot, error) {
	schedule, err := r.GetScheduleByCompetitionId(competitionId)
	if err != nil {
		return nil, err
	}

	var session domain.Session
	result := r.dbClient.Where("schedule_id = ? and number = ?", schedule.ID, sessionNumber).Preload("ScheduleSlots.Contestant.SportsOrganization.Address").First(&session)
	if result.Error != nil {
		return nil, err
	}

	//Can't be done via sql because starting apparatuses are stored like json
	var slots = make([]domain.ScheduleSlot, 0)
	for _, slot := range session.ScheduleSlots {
		if slot.StartingApparatus == apparatus {
			slots = append(slots, slot)
		}
	}

	return slots, nil
}

func (r *ScoringRepoPg) SubmitTempScore(tempScore *domain.TempScore) error {
	tempScore.ID = uuid.New()

	result := r.dbClient.Create(tempScore)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ScoringRepoPg) GetContestantsTempScores(competitionId, contestantId uuid.UUID, apparatus domain.Apparatus) ([]domain.TempScore, error) {
	var tempScores []domain.TempScore

	result := r.dbClient.
		Where("competition_id = ? and contestant_id = ? and apparatus = ?", competitionId, contestantId, apparatus).
		Preload("Judge.SportsOrganization.Address").
		Find(&tempScores)
	if result.Error != nil {
		return nil, result.Error
	}

	return tempScores, nil
}
func (r *ScoringRepoPg) SubmitScore(score *domain.Score) error {
	result := r.dbClient.Create(score)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
