package service

import (
	application_pb "common/proto/application/generated"
	scheduling_pb "common/proto/scheduling/generated"
	"context"
	"github.com/google/uuid"
	"scoring_service/api/mapper"
	"scoring_service/core/domain"
	"scoring_service/core/repo"
)

type ScheduleService struct {
	appClient application_pb.ApplicationServiceClient
	schClient scheduling_pb.SchedulingServiceClient
	schRepo   repo.ScheduleRepo
}

func NewScheduleService(appClient application_pb.ApplicationServiceClient, schClient scheduling_pb.SchedulingServiceClient, schRepo repo.ScheduleRepo) *ScheduleService {
	return &ScheduleService{appClient: appClient, schClient: schClient, schRepo: schRepo}
}

func (s *ScheduleService) StartCompetition(competitionId uuid.UUID) error {
	competitionPb, err := s.appClient.GetCompetitionById(context.Background(), &application_pb.IdMessage{Id: competitionId.String()})
	if err != nil {
		return err
	}

	schedulePb, err := s.schClient.GetByCompetitionId(context.Background(), &scheduling_pb.IdMessage{Id: competitionId.String()})
	if err != nil {
		return err
	}

	//Process slots to sessions
	slots := mapper.SlotListPbToDom(schedulePb.Slots)
	sessions, err := s.generateSessions(slots)
	if err != nil {
		return err
	}

	schedule := &domain.Schedule{
		ID:             uuid.New(),
		CompetitionID:  uuid.UUID{},
		Competition:    *mapper.CompetitionPbToDom(competitionPb),
		ApparatusOrder: mapper.ApparatusListPbToDom(schedulePb.GetApparatusOrder()),
		Sessions:       sessions,
	}

	for idx, _ := range sessions {
		sessions[idx].Schedule = *schedule //Because of gorm
	}

	//TODO SAVE TO DB GODSPEED
	return s.schRepo.Save(schedule)
}

func (s *ScheduleService) generateSessions(slots []domain.ScheduleSlot) ([]domain.Session, error) {
	filteredSlots := make([]domain.ScheduleSlot, 0)

	//Filter empty slots
	for _, slot := range slots {
		if slot.Contestant.ID != uuid.Nil {
			filteredSlots = append(filteredSlots, slot)
		}
	}

	//Assign sports organization to contestant... (Only name present in scheduling bounded context)
	for idx := range filteredSlots {
		sportsOrgName := filteredSlots[idx].Contestant.SportsOrganization.Name
		sportsOrg, err := s.appClient.GetSportsOrganisationByName(context.Background(), &application_pb.GetSportsOrganisationByNameRequest{Name: sportsOrgName})
		if err != nil {
			return nil, err
		}

		filteredSlots[idx].Contestant.SportsOrganization = *mapper.SportsOrganizationPbToDomApp(sportsOrg)
	}

	// Group by session number
	groupedBySession := make(map[int32][]domain.ScheduleSlot)

	for _, slot := range filteredSlots {
		groupedBySession[slot.Session.Number] = append(groupedBySession[slot.Session.Number], slot)
	}

	//Create sessions
	sessions := make([]domain.Session, 0)
	for _, sessionSlots := range groupedBySession {
		session := domain.Session{
			ID:              uuid.New(),
			Number:          sessionSlots[0].Session.Number,
			CurrentRotation: 0,
			Finished:        false,
			ScheduleSlots:   sessionSlots,
		}
		for idx, _ := range sessionSlots {
			sessionSlots[idx].Session = session //Updating (until now only number was defined)
		}

		sessions = append(sessions, session)
	}
	return sessions, nil
}
