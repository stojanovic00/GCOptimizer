package mapper

import (
	application_pb "common/proto/application/generated"
	scheduling_pb "common/proto/scheduling/generated"
	scoring_pb "common/proto/scoring/generated"
	"github.com/google/uuid"
	"scoring_service/core/domain"
	"scoring_service/core/domain/dto"
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

func SportsOrganizationPbToDom(soPb *scoring_pb.SportsOrganization) *domain.SportsOrganization {
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

func SportsOrganizationPbToDomApp(soPb *application_pb.SportsOrganisation) *domain.SportsOrganization {
	id, _ := uuid.Parse(soPb.Id)
	return &domain.SportsOrganization{
		ID:                             id,
		Name:                           soPb.Name,
		Email:                          soPb.Email,
		PhoneNumber:                    soPb.PhoneNumber,
		ContactPersonFullName:          soPb.ContactPersonFullName,
		CompetitionOrganisingPrivilege: soPb.CompetitionOrganisingPrivilege,
		Address:                        *addressPbToDomApp(soPb.Address),
	}
}
func addressPbToDom(address *scoring_pb.Address) *domain.Address {
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

func addressPbToDomApp(address *application_pb.Address) *domain.Address {
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
		Address:         *addressPbToDomApp(compPb.Address),
		TeamComposition: *teamCompositionPbToDom(compPb.TeamComposition),
		Organizer:       *SportsOrganizationPbToDomApp(compPb.Organizer),
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

func ApparatusListDomToPb(appList []domain.Apparatus) []scoring_pb.Apparatus {
	var appListPb []scoring_pb.Apparatus
	for _, app := range appList {
		appListPb = append(appListPb, scoring_pb.Apparatus(app))
	}

	return appListPb
}

func JudgePbToDom(judge *scoring_pb.Judge) *domain.Judge {
	id, _ := uuid.Parse(judge.Id)
	return &domain.Judge{
		ID:                   id,
		FullName:             judge.FullName,
		Email:                judge.Email,
		LicenceType:          domain.LicenceType(judge.LicenceType),
		LicenceName:          judge.LicenceName,
		SportsOrganizationID: uuid.UUID{},
		SportsOrganization:   *SportsOrganizationPbToDom(judge.SportsOrganization),
	}
}
func JudgeDomToPb(judge *domain.Judge) *scoring_pb.Judge {
	return &scoring_pb.Judge{
		Id:                 judge.ID.String(),
		FullName:           judge.FullName,
		Email:              judge.Email,
		LicenceType:        scoring_pb.LicenceType(judge.LicenceType),
		LicenceName:        judge.LicenceName,
		SportsOrganization: sportsOrganizationDomToPb(&judge.SportsOrganization),
	}
}

func JudgeListDomToPb(judges []domain.Judge) []*scoring_pb.Judge {
	var judgePbList []*scoring_pb.Judge
	for _, judgeDom := range judges {
		judgePbList = append(judgePbList, JudgeDomToPb(&judgeDom))
	}

	return judgePbList
}

func addressDomToPb(address *domain.Address) *scoring_pb.Address {
	return &scoring_pb.Address{
		Id:           address.ID.String(),
		Country:      address.Country,
		City:         address.City,
		Street:       address.Street,
		StreetNumber: address.StreetNumber,
	}
}
func sportsOrganizationDomToPb(sportsOrganization *domain.SportsOrganization) *scoring_pb.SportsOrganization {
	return &scoring_pb.SportsOrganization{
		Id:                             sportsOrganization.ID.String(),
		Name:                           sportsOrganization.Name,
		Email:                          sportsOrganization.Email,
		PhoneNumber:                    sportsOrganization.PhoneNumber,
		ContactPersonFullName:          sportsOrganization.ContactPersonFullName,
		CompetitionOrganisingPrivilege: sportsOrganization.CompetitionOrganisingPrivilege,
		Address:                        addressDomToPb(&sportsOrganization.Address),
	}

}

func ScoreCalcMethodPbToDom(method *scoring_pb.ScoreCalculationMethod) *domain.ScoreCalculationMethod {
	var id *uuid.UUID
	if method.Id == "" {
		id = nil
	} else {
		*id, _ = uuid.Parse(method.Id)
	}
	return &domain.ScoreCalculationMethod{
		ID:                id,
		ScoreDeductionNum: method.ScoreDeductionNum,
	}
}

func ScoreCalcMethodDomToPb(method *domain.ScoreCalculationMethod) *scoring_pb.ScoreCalculationMethod {
	return &scoring_pb.ScoreCalculationMethod{
		Id:                method.ID.String(),
		ScoreDeductionNum: method.ScoreDeductionNum,
	}
}

func JudgeJudgingInfoDomToPb(info *dto.JudgeJudgingInfo) *scoring_pb.JudgeJudgingInfo {
	return &scoring_pb.JudgeJudgingInfo{
		Judge:             JudgeDomToPb(&info.Judge),
		CompetitionId:     info.CompetitionId.String(),
		Apparatus:         scoring_pb.Apparatus(info.Apparatus),
		JudgingPanelType:  scoring_pb.JudgingPanelType(info.JudgingPanelType),
		CalculationMethod: ScoreCalcMethodDomToPb(&info.CalculationMethod),
	}
}

func ContestantDomToPb(contestant *domain.Contestant) *scoring_pb.Contestant {
	return &scoring_pb.Contestant{
		Id:                   contestant.ID.String(),
		CompetingId:          contestant.CompetingId,
		FullName:             contestant.FullName,
		SportsOrganization:   sportsOrganizationDomToPb(&contestant.SportsOrganization),
		CompetingApparatuses: ApparatusListDomToPb(contestant.CompetingApparatuses),
		TeamNumber:           contestant.TeamNumber,
		AgeCategory:          contestant.AgeCategory,
	}
}
func ContestantListDomToPb(contestants []domain.Contestant) []*scoring_pb.Contestant {
	var contListPb []*scoring_pb.Contestant
	for _, contestant := range contestants {
		contListPb = append(contListPb, ContestantDomToPb(&contestant))
	}

	return contListPb
}

func ContestantCompetingDomToPb(contestant *domain.Contestant, apparatus domain.Apparatus) *scoring_pb.ContestantCompeting {
	return &scoring_pb.ContestantCompeting{
		Contestant: ContestantDomToPb(contestant),
		Competes:   contestant.CompetesApparatus(apparatus),
	}
}

func ContestantCompetingListDomToPbSorted(contestants []domain.Contestant, apparatus domain.Apparatus) []*scoring_pb.ContestantCompeting {
	var contListPb []*scoring_pb.ContestantCompeting
	for _, contestant := range contestants {
		contListPb = append(contListPb, ContestantCompetingDomToPb(&contestant, apparatus))
	}

	//Filter those who compete on this apparatus and those who don't
	competing := make([]*scoring_pb.ContestantCompeting, 0)
	notCompeting := make([]*scoring_pb.ContestantCompeting, 0)

	for _, contestant := range contListPb {
		if contestant.Competes {
			competing = append(competing, contestant)
		} else {
			notCompeting = append(notCompeting, contestant)
		}
	}
	//First goes all who competes then those who don't
	return append(competing, notCompeting...)
}

func TempScorePbToDom(ts *scoring_pb.TempScore) *domain.TempScore {
	contestantId, _ := uuid.Parse(ts.Contestant.Id)
	competitionId, _ := uuid.Parse(ts.CompetitionId)
	judgeId, _ := uuid.Parse(ts.Judge.Id)
	return &domain.TempScore{
		ID:            uuid.Nil,
		Type:          domain.ScoreType(ts.Type),
		Apparatus:     domain.Apparatus(ts.Apparatus),
		Value:         ts.Value,
		ContestantID:  contestantId,
		Contestant:    domain.Contestant{}, //Will be retrieved from db
		CompetitionID: competitionId,
		Competition:   domain.Competition{}, //Will be retrieved from db
		JudgeID:       judgeId,
		Judge:         domain.Judge{}, //Will be retrieved from db
	}
}

func TempScoreDomToPb(ts *domain.TempScore) *scoring_pb.TempScore {
	return &scoring_pb.TempScore{
		Id:            ts.ID.String(),
		Type:          scoring_pb.ScoreType(ts.Type),
		Apparatus:     scoring_pb.Apparatus(ts.Apparatus),
		Value:         ts.Value,
		Contestant:    &scoring_pb.Contestant{Id: ts.ContestantID.String()},
		CompetitionId: ts.CompetitionID.String(),
		Judge:         JudgeDomToPb(&ts.Judge),
	}
}

func TempScoreListDomToPb(tScores []domain.TempScore) []*scoring_pb.TempScore {
	var tempScoreListPb []*scoring_pb.TempScore
	for _, tempScore := range tScores {
		tempScoreListPb = append(tempScoreListPb, TempScoreDomToPb(&tempScore))
	}

	return tempScoreListPb
}

func ScorePbToDom(s *scoring_pb.Score) *domain.Score {
	var id uuid.UUID
	if s.Id == "" {
		id = uuid.Nil
	} else {
		id, _ = uuid.Parse(s.Id)
	}
	competitionId, _ := uuid.Parse(s.CompetitionId)
	contestantId, _ := uuid.Parse(s.Contestant.Id)
	return &domain.Score{
		ID:            id,
		Apparatus:     domain.Apparatus(s.Apparatus),
		DScore:        s.DScore,
		EScore:        s.EScore,
		TotalScore:    s.TotalScore,
		CompetitionID: competitionId,
		Competition:   domain.Competition{}, //  Will be retrieved from db
		ContestantID:  contestantId,
		Contestant:    domain.Contestant{}, //  Will be retrieved from db
		Submitted:     s.Submitted,
	}
}

func ScoreDomToPb(s *domain.Score) *scoring_pb.Score {
	return &scoring_pb.Score{
		Id:            s.ID.String(),
		Apparatus:     scoring_pb.Apparatus(s.Apparatus),
		DScore:        s.DScore,
		EScore:        s.EScore,
		TotalScore:    s.TotalScore,
		CompetitionId: s.CompetitionID.String(),
		Contestant:    &scoring_pb.Contestant{Id: s.ContestantID.String()},
		Submitted:     s.Submitted,
	}
}
