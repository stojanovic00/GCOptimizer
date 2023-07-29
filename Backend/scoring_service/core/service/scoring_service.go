package service

import (
	"github.com/google/uuid"
	"scoring_service/core/domain"
	"scoring_service/core/domain/dto"
	"scoring_service/core/repo"
)

type ScoringService struct {
	scRepo repo.ScoringRepo
}

func NewScoringService(scRepo repo.ScoringRepo) *ScoringService {
	return &ScoringService{scRepo: scRepo}
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
