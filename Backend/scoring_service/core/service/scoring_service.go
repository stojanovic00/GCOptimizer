package service

import (
	"github.com/google/uuid"
	"scoring_service/core/domain"
	"scoring_service/core/domain/dto"
	"scoring_service/core/repo"
	"sort"
)

type ScoringService struct {
	scRepo repo.ScoringRepo
	jpRepo repo.JudgePanelRepo
}

func NewScoringService(scRepo repo.ScoringRepo, jpRepo repo.JudgePanelRepo) *ScoringService {
	return &ScoringService{scRepo: scRepo, jpRepo: jpRepo}
}

func (s *ScoringService) GetJudgeJudgingInfo(email string) (*dto.JudgeJudgingInfo, error) {
	return s.scRepo.GetJudgeJudgingInfo(email)
}

func (s *ScoringService) GetCurrentApparatusSlots(competitionId uuid.UUID, apparatus domain.Apparatus) ([]domain.ScheduleSlot, error) {
	session, err := s.scRepo.GetCurrentSession(competitionId)
	if err != nil {
		return nil, err
	}

	schedule, err := s.scRepo.GetScheduleByCompetitionId(competitionId)
	if err != nil {
		return nil, err
	}

	var apparatusIndex int
	for idx, app := range schedule.ApparatusOrder {
		if app == apparatus {
			apparatusIndex = idx
			break
		}
	}
	apparatusCount := len(schedule.ApparatusOrder)
	//Calculates depending on rotation number on which apparatus slots started, so they are now on this apparatus
	contestantsStartingApparatusIndex := (apparatusIndex + (apparatusCount - int(session.CurrentRotation))) % apparatusCount

	slots, err := s.scRepo.GetSlotsWithStartingApparatus(competitionId, session.Number, schedule.ApparatusOrder[contestantsStartingApparatusIndex])
	if err != nil {
		return nil, err
	}

	return slots, nil
}

func (s *ScoringService) GetCurrentApparatusContestants(competitionId uuid.UUID, apparatus domain.Apparatus) ([]domain.Contestant, error) {
	slots, err := s.GetCurrentApparatusSlots(competitionId, apparatus)
	if err != nil {
		return nil, err
	}

	var contestants = make([]domain.Contestant, 0)
	for _, slot := range slots {
		contestants = append(contestants, slot.Contestant)
	}

	return contestants, nil
}

func (s *ScoringService) GetNextCurrentApparatusContestant(competitionId uuid.UUID, apparatus domain.Apparatus) (*domain.Contestant, error) {

	slots, err := s.GetCurrentApparatusSlots(competitionId, apparatus)
	if err != nil {
		return nil, err
	}

	//Finds first contestant that  didn't score yet
	var contestant domain.Contestant
	for _, slot := range slots {
		competed := false
		for _, app := range slot.ScoredApparatuses {
			if app == apparatus {
				competed = true
				break
			}
		}
		if !competed && slot.Contestant.CompetesApparatus(apparatus) {
			contestant = slot.Contestant
			break
		}
	}

	return &contestant, nil
}

func (s *ScoringService) SubmitTempScore(tempScore *domain.TempScore) error {
	return s.scRepo.SubmitTempScore(tempScore)
}
func (s *ScoringService) GetContestantsTempScores(competitionId, contestantId uuid.UUID, apparatus domain.Apparatus) ([]domain.TempScore, error) {
	return s.scRepo.GetContestantsTempScores(competitionId, contestantId, apparatus)
}

func (s *ScoringService) CanCalculateScore(competitionId, contestantId uuid.UUID, apparatus domain.Apparatus) (bool, error) {
	tempScores, err := s.GetContestantsTempScores(competitionId, contestantId, apparatus)
	if err != nil {
		return false, err
	}

	panels, err := s.jpRepo.GetApparatusPanels(competitionId, apparatus)
	if err != nil {
		return false, err
	}

	lenSum := 0
	for _, panel := range panels {
		lenSum += len(panel.Judges)
	}

	return len(tempScores) == lenSum, nil
}

func (s *ScoringService) CalculateScore(competitionId, contestantId uuid.UUID, apparatus domain.Apparatus) (*domain.Score, error) {
	tempScores, err := s.GetContestantsTempScores(competitionId, contestantId, apparatus)
	if err != nil {
		return nil, err
	}

	//Group scores
	dScores := make([]domain.TempScore, 0)
	eScores := make([]domain.TempScore, 0)

	for _, tempScore := range tempScores {
		switch tempScore.Type {
		case domain.D:
			dScores = append(dScores, tempScore)
		case domain.E:
			eScores = append(eScores, tempScore)
		}
	}

	//Calculate d score
	var dSum float32 = 0
	for _, dScore := range dScores {
		dSum += dScore.Value
	}
	dAverage := dSum / float32(len(dScores))

	//Calculate e score

	//Get calculation method for e panel
	ePanel, err := s.jpRepo.GetJudgePanelByCompetitionIdAndApparatus(competitionId, apparatus, domain.EPanel)
	if err != nil {
		return nil, err
	}

	deductionNumber := ePanel.ScoreCalculationMethod.ScoreDeductionNum

	// Sort the slice in ascending order based on the Value field
	sort.Slice(eScores, func(i, j int) bool {
		return eScores[i].Value < eScores[j].Value
	})

	//Deduce n highest and lowest scores
	eMiddleScores := eScores[deductionNumber:(len(eScores) - int(deductionNumber))]

	var eSum float32 = 0
	for _, eScore := range eMiddleScores {
		eSum += eScore.Value
	}
	eAverage := eSum / float32(len(eMiddleScores))

	score := &domain.Score{
		ID:            uuid.New(),
		Apparatus:     apparatus,
		DScore:        dAverage,
		EScore:        eAverage,
		TotalScore:    dAverage + eAverage,
		CompetitionID: competitionId,
		Competition:   domain.Competition{}, // Db resolved
		ContestantID:  contestantId,
		Contestant:    domain.Contestant{}, // Db resolved
		Submitted:     false,
	}

	err = s.scRepo.SaveScore(score)
	if err != nil {
		return nil, err
	}

	return score, nil

}

func (s *ScoringService) SubmitScore(score *domain.Score) error {
	return s.scRepo.SubmitScore(score)
}

func (s *ScoringService) GetScore(competitionId, contestantId uuid.UUID, apparatus domain.Apparatus) (*domain.Score, error) {
	return s.scRepo.GetScore(competitionId, contestantId, apparatus)
}

func (s *ScoringService) FinishRotation(competitionId uuid.UUID) error {
	return s.scRepo.FinishRotation(competitionId)
}

func (s *ScoringService) FinishSession(competitionId uuid.UUID) error {
	return s.scRepo.FinishSession(competitionId)
}
func (s *ScoringService) IsRotationFinished(competitionId uuid.UUID) (bool, error) {
	return s.scRepo.IsRotationFinished(competitionId)
}
func (s *ScoringService) IsSessionFinished(competitionId uuid.UUID) (bool, error) {
	return s.scRepo.IsSessionFinished(competitionId)
}
func (s *ScoringService) IsCompetitionFinished(competitionId uuid.UUID) (bool, error) {
	return s.scRepo.IsCompetitionFinished(competitionId)
}
