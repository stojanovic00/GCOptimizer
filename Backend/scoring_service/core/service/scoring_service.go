package service

import (
	"errors"
	"fmt"
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
func (s *ScoringService) GetCurrentSessionInfo(competitionId uuid.UUID) (*dto.CurrentSessionInfo, error) {
	session, err := s.scRepo.GetCurrentSession(competitionId)
	if err != nil {
		return nil, err
	}

	//If there is no more active session, competition is finished
	if session == nil {
		return &dto.CurrentSessionInfo{
			CurrentRotation:     0,
			CurrentSession:      0,
			RotationFinished:    false,
			SessionFinished:     false,
			CompetitionFinished: true,
		}, nil
	}

	isRotationFinished, err := s.IsRotationFinished(competitionId)
	if err != nil {
		return nil, err
	}

	isSessionFinished, err := s.IsSessionFinished(competitionId)
	if err != nil {
		return nil, err
	}

	isCompetitionFinished, err := s.IsCompetitionFinished(competitionId)
	if err != nil {
		return nil, err
	}

	//For frontend purposes (so user can't increment rotations no more)
	if isSessionFinished {
		isRotationFinished = false
	}
	return &dto.CurrentSessionInfo{
		CurrentRotation:     session.CurrentRotation,
		CurrentSession:      session.Number,
		RotationFinished:    isRotationFinished,
		SessionFinished:     isSessionFinished,
		CompetitionFinished: isCompetitionFinished,
	}, nil

}

func (s *ScoringService) GenerateScoreboards(compId uuid.UUID) error {
	schedule, err := s.scRepo.GetScheduleByCompetitionId(compId)
	if err != nil {
		return err
	}

	switch schedule.Competition.Type {
	case domain.Qualifications:
		err := s.GenerateAllAroundScoreboards(schedule)
		if err != nil {
			return err
		}
		err = s.GenerateTeamScoreboards(schedule)
		if err != nil {
			return err
		}
	default:
		return errors.New("no such competition type")
	}

	return nil
}

type AllAroundScoreboardSlots []domain.AllAroundScoreboardSlot

func (slots AllAroundScoreboardSlots) CompareWithEPrivilege(i, j int) bool {
	if slots[i].TotalScore != slots[j].TotalScore {
		// Sort by TotalScore in descending order
		return slots[i].TotalScore > slots[j].TotalScore
	} else if slots[i].TotalEScore != slots[j].TotalEScore {
		// If TotalScore is the same, sort by TotalEScore in descending order
		return slots[i].TotalEScore > slots[j].TotalEScore
	} else {
		// If both TotalScore and EScore are the same, sort by DScore in descending order
		return slots[i].TotalDScore > slots[j].TotalDScore
	}
}

func (s *ScoringService) GenerateAllAroundScoreboards(schedule *domain.Schedule) error {
	scores, err := s.scRepo.GetScores(schedule.Competition.ID)
	if err != nil {
		return err
	}

	//Group by contestant (ID)
	scoresByContestant := make(map[uuid.UUID][]domain.Score)
	for _, score := range scores {
		scoresByContestant[score.ContestantID] = append(scoresByContestant[score.ContestantID], score)
	}

	//Create scoreboard slots
	var slots AllAroundScoreboardSlots
	for _, contestantsScores := range scoresByContestant {

		var totalE float32 = 0
		var totalD float32 = 0
		var scores []domain.Score
		for _, score := range contestantsScores {
			totalE += score.EScore
			totalD += score.DScore
			scores = append(scores, score)
		}

		slot := domain.AllAroundScoreboardSlot{
			ID:           uuid.New(),
			ContestantID: scores[0].ContestantID,
			Contestant:   scores[0].Contestant,
			Scores:       scores,
			TotalEScore:  totalE,
			TotalDScore:  totalD,
			TotalScore:   totalE + totalD,
		}
		slots = append(slots, slot)
	}

	//Group by age category
	slotsByAgeCat := make(map[string]AllAroundScoreboardSlots)
	for _, slot := range slots {
		slotsByAgeCat[slot.Contestant.AgeCategory] = append(slotsByAgeCat[slot.Contestant.AgeCategory], slot)
	}

	//Assign places
	for _, slotsByACat := range slotsByAgeCat {
		//Sort
		sort.Slice(slotsByACat, slotsByACat.CompareWithEPrivilege)

		//Assign places depending on having same score or not
		//First one is always on first place
		placeCounter := 1
		slotsByACat[0].Place = placeCounter
		//Starting from second member
		for idx := 1; idx < len(slotsByACat); idx++ {
			if slotsByACat[idx].TotalScore != slotsByACat[idx-1].TotalScore {
				placeCounter++
			} else if schedule.Competition.Tiebreak && (slotsByACat[idx].TotalDScore != slotsByACat[idx-1].TotalDScore) {
				placeCounter++
			}
			slotsByACat[idx].Place = placeCounter
		}
	}

	//Create and save scoreboards
	for _, slotsByACat := range slotsByAgeCat {
		scoreBoard := &domain.AllAroundScoreboard{
			ID:            uuid.New(),
			CompetitionID: schedule.CompetitionID,
			AgeCategory:   slotsByACat[0].Contestant.AgeCategory,
			TieBrake:      schedule.Competition.Tiebreak,
			Apparatuses:   schedule.ApparatusOrder,
			Slots:         slotsByACat,
		}
		err = s.scRepo.SaveAllAroundScoreBoard(scoreBoard)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *ScoringService) GenerateTeamScoreboards(schedule *domain.Schedule) error {
	allAroundScoreboards, err := s.GetAllAroundScoreBoards(schedule.CompetitionID)
	if err != nil {
		return err
	}

	//Because I am using all around scoreboard, slots are already grouped by age category
	for _, scoreBoard := range allAroundScoreboards {

		slotsByOrgTeam := make(map[string]AllAroundScoreboardSlots)
		for _, slot := range scoreBoard.Slots {
			key := fmt.Sprintf("%s-%d", slot.Contestant.SportsOrganizationID, slot.Contestant.TeamNumber)
			slotsByOrgTeam[key] = append(slotsByOrgTeam[key], slot)
		}

		//Create teams slots
		var teamScoreboardSlots []domain.TeamScoreboardSlot
		for _, teamSlots := range slotsByOrgTeam {
			teamBaseNum := schedule.Competition.TeamComposition.BaseContestantNumber

			//Calculate best scores for each apparatus and sum them to get total E and D
			scoresByApparatus := make(map[domain.Apparatus][]domain.Score)
			for _, slot := range teamSlots {
				for _, score := range slot.Scores {
					scoresByApparatus[score.Apparatus] = append(scoresByApparatus[score.Apparatus], score)
				}
			}

			//Check if there is enough contestants to make a team
			if len(scoresByApparatus[schedule.ApparatusOrder[0]]) < teamBaseNum {
				continue
			}

			//Sum best n scores on each apparatus and add it to total score
			var totalScore float32 = 0
			apparatusTotalScores := make(map[domain.Apparatus]float32)
			var allScores []domain.Score
			for apparatus, scores := range scoresByApparatus {
				sort.Slice(scores, func(i, j int) bool {
					return scores[i].TotalScore > scores[j].TotalScore
				})

				//For possible out of bounds error
				var limit int
				if teamBaseNum <= len(scores) {
					limit = teamBaseNum
				} else {
					limit = len(scores)
				}

				best := scores[:limit]
				var bestSum float32 = 0
				for _, score := range best {
					bestSum += score.TotalScore
				}
				totalScore += bestSum
				apparatusTotalScores[apparatus] = bestSum
				allScores = append(allScores, scores...)
			}

			teamSlot := domain.TeamScoreboardSlot{
				ID:                   uuid.New(),
				SportsOrganizationID: teamSlots[0].Contestant.SportsOrganizationID,
				TeamNumber:           int(teamSlots[0].Contestant.TeamNumber),
				Scores:               allScores,
				ApparatusTotalScores: apparatusTotalScores,
				TotalScore:           totalScore,
			}
			teamScoreboardSlots = append(teamScoreboardSlots, teamSlot)
		}

		//Assign positions
		sort.Slice(teamScoreboardSlots, func(i, j int) bool {
			return teamScoreboardSlots[i].TotalScore > teamScoreboardSlots[j].TotalScore
		})

		placeCounter := 1
		teamScoreboardSlots[0].Place = placeCounter
		for idx := 1; idx < len(teamScoreboardSlots); idx++ {
			if teamScoreboardSlots[idx].TotalScore != teamScoreboardSlots[idx-1].TotalScore {
				placeCounter++
			}
			teamScoreboardSlots[idx].Place = placeCounter
		}

		//Save scoreboard
		teamScoreBoard := &domain.TeamScoreboard{
			ID:            uuid.New(),
			CompetitionID: schedule.CompetitionID,
			AgeCategory:   scoreBoard.AgeCategory,
			Apparatuses:   scoreBoard.Apparatuses,
			Slots:         teamScoreboardSlots,
		}

		err = s.scRepo.SaveTeamScoreBoard(teamScoreBoard)
		if err != nil {
			return err
		}

	}
	return nil
}

func (s *ScoringService) GetAllAroundScoreBoards(competitionId uuid.UUID) ([]domain.AllAroundScoreboard, error) {
	return s.scRepo.GetAllAroundScoreBoards(competitionId)
}
func (s *ScoringService) GetTeamScoreBoards(competitionId uuid.UUID) ([]domain.TeamScoreboard, error) {
	return s.scRepo.GetTeamScoreBoards(competitionId)
}
