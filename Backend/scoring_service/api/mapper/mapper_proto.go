package mapper

import (
	application_pb "common/proto/application/generated"
	scheduling_pb "common/proto/scheduling/generated"
	"github.com/google/uuid"
	"scoring_service/core/domain"
	"time"
)

func teamCompositionPbToDom(teamComp *application_pb.TeamComposition) *domain.TeamComposition {
	id, _ := uuid.Parse(teamComp.Id)
	return &domain.TeamComposition{
		ID:                    id,
		BaseContestantNumber:  int(teamComp.BaseContestantNumber),
		BonusContestantNumber: int(teamComp.BonusContestantNumber),
		MultiCategoryTeam:     teamComp.MultiCategoryTeam,
	}
}

func ageCategoryListPbToDom(categories []*application_pb.AgeCategory) []string {
	var catDomList []string
	for _, catPb := range categories {
		catDomList = append(catDomList, catPb.Name)
	}

	return catDomList
}

func SportsOrganizationPbToDom(soPb *application_pb.SportsOrganisation) *domain.SportsOrganization {
	id, _ := uuid.Parse(soPb.Id)
	return &domain.SportsOrganization{
		ID:                             id,
		Name:                           soPb.Name,
		Email:                          soPb.Email,
		PhoneNumber:                    soPb.PhoneNumber,
		ContactPersonFullName:          soPb.ContactPersonFullName,
		CompetitionOrganisingPrivilege: soPb.CompetitionOrganisingPrivilege,
		Address:                        *addressPbToDom(soPb.Address),
	}
}

func addressPbToDom(address *application_pb.Address) *domain.Address {
	var id uuid.UUID
	if address.Id != "" {
		id, _ = uuid.Parse(address.Id)
	} else {
		id = uuid.Nil
	}

	return &domain.Address{
		ID:           id,
		Country:      address.Country,
		City:         address.City,
		Street:       address.Street,
		StreetNumber: address.StreetNumber,
	}
}

func CompetitionPbToDom(compPb *application_pb.Competition) *domain.Competition {
	id, _ := uuid.Parse(compPb.Id)
	return &domain.Competition{
		ID:              id,
		Name:            compPb.Name,
		StartDate:       time.Unix(compPb.StartDate, 0),
		EndDate:         time.Unix(compPb.EndDate, 0),
		Gender:          domain.Gender(compPb.Gender),
		Type:            domain.CompetitionType(compPb.Type),
		Tiebreak:        compPb.Tiebreak,
		Address:         *addressPbToDom(compPb.Address),
		TeamComposition: *teamCompositionPbToDom(compPb.TeamComposition),
		Organizer:       *SportsOrganizationPbToDom(compPb.Organizer),
		AgeCategories:   ageCategoryListPbToDom(compPb.AgeCategories),
	}
}

func ApparatusListPbToDom(apps []scheduling_pb.ApparatusType) []domain.Apparatus {
	var appList []domain.Apparatus
	for _, appPb := range apps {
		appList = append(appList, domain.Apparatus(appPb.Number()))
	}

	return appList
}

func contestantPbToDom(contestant *scheduling_pb.ContestantInfo) *domain.Contestant {
	id, _ := uuid.Parse(contestant.Id)
	return &domain.Contestant{
		ID:                   id,
		CompetingId:          contestant.ContestantCompId,
		FullName:             contestant.Name,
		SportsOrganizationID: uuid.UUID{},
		SportsOrganization: domain.SportsOrganization{ //Fill it later in code!!!
			Name: contestant.Organization,
		},
		CompetingApparatuses: ApparatusListPbToDom(contestant.CompetingApparatuses),
		TeamNumber:           contestant.TeamNumber,
		AgeCategory:          contestant.AgeCategory,
	}
}

func slotPbToDom(slot *scheduling_pb.ScheduleSlot) *domain.ScheduleSlot {
	var contestant domain.Contestant
	if slot.ContestantInfo == nil {
		contestant = domain.Contestant{}
	} else {
		contestant = *contestantPbToDom(slot.ContestantInfo)
	}
	return &domain.ScheduleSlot{
		ID:        uuid.New(), //It is saved right upon receiving in service
		SessionID: uuid.UUID{},
		Session: domain.Session{
			Number: slot.Session,
		}, //Will be assigned later
		StartingApparatus: domain.Apparatus(slot.StartingApparatus.Number()),
		ScoredApparatuses: []domain.Apparatus{},
		ContestantID:      uuid.UUID{},
		Contestant:        contestant,
	}
}
func SlotListPbToDom(slots []*scheduling_pb.ScheduleSlot) []domain.ScheduleSlot {
	var slotList []domain.ScheduleSlot
	for _, slotPb := range slots {
		slotList = append(slotList, *slotPbToDom(slotPb))
	}

	return slotList
}
