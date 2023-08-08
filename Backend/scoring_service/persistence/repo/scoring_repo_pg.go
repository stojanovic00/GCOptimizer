package repo

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"scoring_service/core/domain"
	"scoring_service/core/domain/dto"
	"sort"
)

type ScoringRepoPg struct {
	dbClient *gorm.DB
}

func NewScoringRepoPg(dbClient *gorm.DB) *ScoringRepoPg {
	return &ScoringRepoPg{dbClient: dbClient}
}

// This function is made just to patch things up before presentation, logic of geting judge info should be changed
// This function will also work on long run (one judge can't physicaly be on 2 active competitions)
// but i don't think it is good approach
func (r *ScoringRepoPg) GetPanelWithActiveCompetition(panels []domain.Panel) (*domain.Panel, error) {
	for _, panel := range panels {
		finished, err := r.IsCompetitionFinished(panel.CompetitionID)
		if err != nil {
			return nil, err
		}
		if !finished {
			return &panel, nil
		}
	}
	return nil, errors.New("no panel with active competition found")
}

func (r *ScoringRepoPg) GetJudgeJudgingInfo(email string) (*dto.JudgeJudgingInfo, error) {
	var judge domain.Judge

	result := r.dbClient.Where("email = ?", email).Preload("SportsOrganization.Address").First(&judge)
	if result.Error != nil {
		return nil, result.Error
	}

	var panels []domain.Panel
	result = r.dbClient.Model(domain.Panel{}).
		Joins("left join panel_judges pj on id = pj.panel_id").
		Where("judge_id = ?", judge.ID).
		Preload("ScoreCalculationMethod").
		Find(&panels)
	if result.Error != nil {
		return nil, result.Error
	}
	panel, err := r.GetPanelWithActiveCompetition(panels)
	if err != nil {
		return nil, err
	}

	//If it is D panel it must retrieve E score calculation method so frontend can show deductions on D panel GUI
	var calcMethod domain.ScoreCalculationMethod
	if panel.Type == domain.DPanel {
		//retrieve E panel for this competition and this apparatus
		var ePanel domain.Panel
		result = r.dbClient.
			Where("competition_id = ? and apparatus = ? and type = ?", panel.CompetitionID, panel.Apparatus, domain.EPanel).
			Preload("ScoreCalculationMethod").
			Find(&ePanel)
		if result.Error != nil {
			return nil, result.Error
		}

		calcMethod = ePanel.ScoreCalculationMethod
	} else {
		calcMethod = panel.ScoreCalculationMethod
	}

	return &dto.JudgeJudgingInfo{
		Judge:             judge,
		CompetitionId:     panel.CompetitionID,
		Apparatus:         panel.Apparatus,
		JudgingPanelType:  panel.Type,
		CalculationMethod: calcMethod,
	}, nil

}
func (r *ScoringRepoPg) GetScheduleByCompetitionId(competitionId uuid.UUID) (*domain.Schedule, error) {
	var schedule domain.Schedule

	result := r.dbClient.
		Where("competition_id = ?", competitionId).
		Preload("Competition.TeamComposition").
		First(&schedule)
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
	result := r.dbClient.
		Where("schedule_id = ? and finished = false", schedule.ID).
		Order("number asc").
		First(&session)
	if result.Error != nil {
		return nil, err
	}

	return &session, nil
}

func (r *ScoringRepoPg) GetCurrentSessionWithSlots(competitionId uuid.UUID) (*domain.Session, error) {
	schedule, err := r.GetScheduleByCompetitionId(competitionId)
	if err != nil {
		return nil, err
	}

	var session domain.Session
	result := r.dbClient.
		Where("schedule_id = ? and finished = false", schedule.ID).
		Order("number asc").
		Preload("ScheduleSlots.Session").
		Preload("ScheduleSlots.Contestant").
		First(&session)
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

	//!!!SORT BY POSITION ASCENDING!!!
	sort.Slice(slots, func(i, j int) bool {
		return slots[i].Position < slots[j].Position
	})

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
func (r *ScoringRepoPg) SaveScore(score *domain.Score) error {
	result := r.dbClient.Create(score)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ScoringRepoPg) SubmitScore(score *domain.Score) error {
	//Update score to submitted
	scoreDb, err := r.GetScore(score.CompetitionID, score.ContestantID, score.Apparatus)
	if err != nil {
		return err
	}

	scoreDb.Submitted = true
	result := r.dbClient.Save(scoreDb)
	if result.Error != nil {
		return result.Error
	}

	session, err := r.GetCurrentSession(score.CompetitionID)
	if err != nil {
		return err
	}

	//Update schedule slots scoredApparatuses
	var slot domain.ScheduleSlot
	result = r.dbClient.
		Where("contestant_id = ? and session_id = ?", score.ContestantID, session.ID).
		First(&slot)
	if result.Error != nil {
		return result.Error
	}
	slot.ScoredApparatuses = append(slot.ScoredApparatuses, score.Apparatus)

	result = r.dbClient.Save(slot)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (r *ScoringRepoPg) GetScore(competitionId, contestantId uuid.UUID, apparatus domain.Apparatus) (*domain.Score, error) {
	var score domain.Score

	result := r.dbClient.
		Where("competition_id = ? and contestant_id = ? and apparatus = ?", competitionId, contestantId, apparatus).
		First(&score)
	if result.Error != nil {
		return nil, result.Error
	}

	return &score, nil
}

func (r *ScoringRepoPg) FinishRotation(competitionId uuid.UUID) error {
	session, err := r.GetCurrentSessionWithSlots(competitionId)
	if err != nil {
		return err
	}

	// Start a transaction
	tx := r.dbClient.Begin()

	session.CurrentRotation++

	//MOVE EVERYONE ONE POSITION UP INSIDE APPARATUS GROUP

	//This is map of counters for each apparatus separately
	apparatusCount := make(map[domain.Apparatus]int)
	for _, slot := range session.ScheduleSlots {
		apparatusCount[slot.StartingApparatus]++
	}

	for idx := range session.ScheduleSlots {
		slot := &session.ScheduleSlots[idx]
		slot.Position = (slot.Position + 1) % apparatusCount[slot.StartingApparatus]
		//Gorm will not update slots when saving session, so we are doing it one by one
		//That's why I started transaction to not make n calls to db
		result := tx.Save(&slot)
		if result.Error != nil {
			tx.Rollback()
			return result.Error
		}
	}

	result := tx.Save(&session)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	// Commit the transaction
	err = tx.Commit().Error
	if err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
func (r *ScoringRepoPg) FinishSession(competitionId uuid.UUID) error {
	session, err := r.GetCurrentSession(competitionId)
	if err != nil {
		return err
	}
	if session == nil {
		return errors.New("no more sessions to finish")
	}
	session.Finished = true

	result := r.dbClient.Save(session)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
func (r *ScoringRepoPg) IsRotationFinished(competitionId uuid.UUID) (bool, error) {
	schedule, err := r.GetScheduleByCompetitionId(competitionId)
	if err != nil {
		return false, err
	}

	session, err := r.GetCurrentSession(competitionId)
	if err != nil {
		return false, err
	}
	if session == nil {
		return false, errors.New("no more active sessions")
	}

	//Logic start
	apparatusOrder := schedule.ApparatusOrder
	currentRotation := session.CurrentRotation
	rotationCount := len(apparatusOrder)
	//Check if all contestants scored inside current rotation
	for _, apparatus := range apparatusOrder {
		//Contestants that started on apparatus, on which apparatus are they in current rotation?(currentApparatus)
		var apparatusIndex int
		for idx, app := range apparatusOrder {
			if app == apparatus {
				apparatusIndex = idx
				break
			}
		}

		//They are on this apparatus(current apparatus)
		currentApparatusIndex := (apparatusIndex + int(currentRotation)) % rotationCount
		currentApparatus := apparatusOrder[currentApparatusIndex]

		//Get all slots for apparatus and check if all contestants that compete currentApparatus
		//scored on it (currentApparatus present in scoredApparatuses)
		slots, err := r.GetSlotsWithStartingApparatus(competitionId, session.Number, apparatus)
		if err != nil {
			return false, err
		}

		for _, slot := range slots {
			if slot.CompetesApparatus(currentApparatus) {
				//Check if he scored on this apparatus
				scored := false
				for _, app := range slot.ScoredApparatuses {
					if app == currentApparatus {
						scored = true
						break
					}
				}
				if !scored {
					return false, nil //Some contestant didn't score yet inside current rotation, so rotation isn't finished
				}
			}
		}

	}
	//Logic end

	//If it didn't return false inside all of those for loops then rotation is finished
	return true, err
}

func (r *ScoringRepoPg) IsSessionFinished(competitionId uuid.UUID) (bool, error) {
	session, err := r.GetCurrentSessionWithSlots(competitionId)
	if err != nil {
		return false, err
	}
	if session == nil {
		return false, errors.New("no more active sessions")
	}

	allCompeted := true
	for _, slot := range session.ScheduleSlots {

		if len(slot.ScoredApparatuses) != len(slot.CompetingApparatuses) {
			allCompeted = false
			break
		}
	}

	return allCompeted, nil
}
func (r *ScoringRepoPg) IsCompetitionFinished(competitionId uuid.UUID) (bool, error) {
	schedule, err := r.GetScheduleByCompetitionId(competitionId)
	if err != nil {
		return false, err
	}

	var count int64
	result := r.dbClient.
		Model(domain.Session{}).
		Where("schedule_id = ? and finished = false", schedule.ID).
		Count(&count)

	if result.Error != nil {
		return false, result.Error
	}

	if count > 0 {
		return false, nil
	}
	return true, nil
}

func (r *ScoringRepoPg) GetCompetition(competitionId uuid.UUID) (*domain.Competition, error) {
	var competition domain.Competition
	result := r.dbClient.
		Where("id = ?", competitionId).
		Preload("TeamComposition").
		First(&competition)
	if result.Error != nil {
		return nil, result.Error
	}

	return &competition, nil
}

func (r *ScoringRepoPg) GetScores(competitionId uuid.UUID) ([]domain.Score, error) {
	var scores []domain.Score
	result := r.dbClient.
		Where("competition_id = ?", competitionId).
		Preload("Contestant").
		Find(&scores)
	if result.Error != nil {
		return nil, result.Error
	}
	return scores, nil
}

func (r *ScoringRepoPg) GetScoresByContestantId(contestantId uuid.UUID) ([]domain.Score, error) {
	var scores []domain.Score
	result := r.dbClient.
		Where("contestant_id = ?", contestantId).
		Find(&scores)
	if result.Error != nil {
		return nil, result.Error
	}
	return scores, nil
}

func (r *ScoringRepoPg) SaveAllAroundScoreBoard(scoreBoard *domain.AllAroundScoreboard) error {
	result := r.dbClient.Create(scoreBoard)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *ScoringRepoPg) GetAllAroundScoreBoards(competitionId uuid.UUID) ([]domain.AllAroundScoreboard, error) {
	var scoreBoards []domain.AllAroundScoreboard
	result := r.dbClient.
		Where("competition_id = ?", competitionId).
		Preload("Slots.Contestant.SportsOrganization.Address").
		Find(&scoreBoards)
	if result.Error != nil {
		return nil, result.Error
	}

	//Calculate info that is not stored
	for _, scoreBoard := range scoreBoards {
		for idx, slot := range scoreBoard.Slots {
			scores, err := r.GetScoresByContestantId(slot.ContestantID)
			if err != nil {
				return nil, err
			}
			var totalE float32 = 0
			var totalD float32 = 0
			for _, score := range scores {
				totalE += score.EScore
				totalD += score.DScore
			}

			scoreBoard.Slots[idx].Scores = scores
			scoreBoard.Slots[idx].TotalEScore = totalE
			scoreBoard.Slots[idx].TotalDScore = totalD
			scoreBoard.Slots[idx].TotalScore = totalE + totalD
		}
	}

	return scoreBoards, nil
}

func (r *ScoringRepoPg) SaveTeamScoreBoard(scoreBoard *domain.TeamScoreboard) error {
	result := r.dbClient.Create(scoreBoard)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ScoringRepoPg) GetTeamScoreBoards(competitionId uuid.UUID) ([]domain.TeamScoreboard, error) {
	var scoreBoards []domain.TeamScoreboard
	result := r.dbClient.
		Where("competition_id = ?", competitionId).
		Preload("Slots.SportsOrganization.Address").
		Find(&scoreBoards)
	if result.Error != nil {
		return nil, result.Error
	}
	//Calculate info that is not stored
	//Not needed for now :)
	return scoreBoards, nil
}
