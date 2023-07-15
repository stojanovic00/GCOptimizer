package handler

import (
	"application_service/core/domain"
	application_pb "common/proto/application/generated"
	"github.com/google/uuid"
	"time"
)

func sportsOrganizationDomToPb(sportsOrganization *domain.SportsOrganization) *application_pb.SportsOrganisation {
	return &application_pb.SportsOrganisation{
		Id:                             sportsOrganization.ID.String(),
		Name:                           sportsOrganization.Name,
		Email:                          sportsOrganization.Email,
		PhoneNumber:                    sportsOrganization.PhoneNumber,
		ContactPersonFullName:          sportsOrganization.ContactPersonFullName,
		CompetitionOrganisingPrivilege: sportsOrganization.CompetitionOrganisingPrivilege,
		Address: &application_pb.Address{
			Id:           sportsOrganization.Address.ID.String(),
			Country:      sportsOrganization.Address.Country,
			City:         sportsOrganization.Address.City,
			Street:       sportsOrganization.Address.Street,
			StreetNumber: sportsOrganization.Address.StreetNumber,
		},
	}

}

func delegationMemberDomToPb(judgeDom *domain.DelegationMember) *application_pb.DelegationMember {
	return &application_pb.DelegationMember{
		Id:       judgeDom.ID.String(),
		FullName: judgeDom.FullName,
		Email:    judgeDom.Email,
		Gender:   application_pb.Gender(judgeDom.Gender),
		Position: &application_pb.DelegationMemberPosition{
			Id:   judgeDom.Position.ID.String(),
			Name: judgeDom.Position.Name,
		},
		Image:              judgeDom.Image,
		SportsOrganisation: sportsOrganizationDomToPb(&judgeDom.SportsOrganization),
	}
}
func delegationMemberPbToDom(judgePb *application_pb.DelegationMember) *domain.DelegationMember {
	return &domain.DelegationMember{
		FullName: judgePb.FullName,
		Email:    judgePb.Email,
		Gender:   domain.Gender(judgePb.Gender),
		Position: domain.DelegationMemberPosition{
			Name: judgePb.Position.Name,
		},
	}
}

func judgeDomToPb(judgeDom *domain.Judge) *application_pb.Judge {
	return &application_pb.Judge{
		DelegationMember: delegationMemberDomToPb(&judgeDom.DelegationMember),
		LicenceType:      application_pb.LicenceType(judgeDom.LicenceType),
		LicenceName:      judgeDom.LicenceName,
	}
}

func judgeListDomToPb(judgeDomList []*domain.Judge) []*application_pb.Judge {
	var judgePbList []*application_pb.Judge
	for _, judgeDom := range judgeDomList {
		judgePbList = append(judgePbList, judgeDomToPb(judgeDom))
	}

	return judgePbList
}
func contestantDomToPb(contestantDom *domain.Contestant) *application_pb.Contestant {
	return &application_pb.Contestant{
		DelegationMember: delegationMemberDomToPb(&contestantDom.DelegationMember),
		DateOfBirth:      contestantDom.DateOfBirth.Unix(),
	}
}

func contestantListDomToPb(contestantDomList []*domain.Contestant) []*application_pb.Contestant {
	var contestantPbList []*application_pb.Contestant
	for _, contestantDom := range contestantDomList {
		contestantPbList = append(contestantPbList, contestantDomToPb(contestantDom))
	}

	return contestantPbList
}

func sportsOrganizationPbToDom(soPb *application_pb.SportsOrganisation) *domain.SportsOrganization {
	id, _ := uuid.Parse(soPb.Id)
	return &domain.SportsOrganization{
		ID:                             id,
		Name:                           soPb.Name,
		Email:                          soPb.Email,
		PhoneNumber:                    soPb.PhoneNumber,
		ContactPersonFullName:          soPb.ContactPersonFullName,
		CompetitionOrganisingPrivilege: soPb.CompetitionOrganisingPrivilege,
		Address: domain.Address{
			Country:      soPb.Address.Country,
			City:         soPb.Address.City,
			Street:       soPb.Address.Street,
			StreetNumber: soPb.Address.StreetNumber,
		},
	}
}

func competitionPbToDom(compPb *application_pb.Competition) *domain.Competition {
	id, _ := uuid.Parse(compPb.Id)
	return &domain.Competition{
		ID:        id,
		Name:      compPb.Name,
		StartDate: time.Unix(compPb.StartDate, 0),
		EndDate:   time.Unix(compPb.EndDate, 0),
		Gender:    domain.Gender(compPb.Gender),
		Type:      domain.CompetitionType(compPb.Type),
		Tiebreak:  compPb.Tiebreak,
		Address: domain.Address{
			Country:      compPb.Address.Country,
			City:         compPb.Address.City,
			Street:       compPb.Address.Street,
			StreetNumber: compPb.Address.StreetNumber,
		},
		TeamComposition: *teamCompositionPbToDom(compPb.TeamComposition),
		//Organizer: *sportsOrganizationPbToDom(compPb.Organizer),
	}
}

func teamCompositionPbToDom(teamComp *application_pb.TeamComposition) *domain.TeamComposition {
	return &domain.TeamComposition{
		BaseContestantNumber:  int(teamComp.BaseContestantNumber),
		BonusContestantNumber: int(teamComp.BonusContestantNumber),
		MultiCategoryTeam:     teamComp.MultiCategoryTeam,
	}
}

func addressDomToPb(addr *domain.Address) *application_pb.Address {
	return &application_pb.Address{
		Id:           addr.ID.String(),
		Country:      addr.Country,
		City:         addr.City,
		Street:       addr.Street,
		StreetNumber: addr.StreetNumber,
	}
}

func delegationMemberPositionDomToPb(pos *domain.DelegationMemberPosition) *application_pb.DelegationMemberPosition {
	return &application_pb.DelegationMemberPosition{
		Id:   pos.ID.String(),
		Name: pos.Name,
	}
}

func delegationMemberPropositionDomToPb(prop *domain.DelegationMemberProposition) *application_pb.DelegationMemberProposition {
	return &application_pb.DelegationMemberProposition{
		Id:        prop.ID.String(),
		Position:  delegationMemberPositionDomToPb(&prop.Position),
		MinNumber: int32(prop.MinNumber),
		MaxNumber: int32(prop.MaxNumber),
	}
}

func delegationMemberPropositionListDomToPb(propList []domain.DelegationMemberProposition) []*application_pb.DelegationMemberProposition {
	var propPbList []*application_pb.DelegationMemberProposition
	for _, propDom := range propList {
		propPbList = append(propPbList, delegationMemberPropositionDomToPb(&propDom))
	}

	return propPbList
}

func teamCompositionDomToPb(comp *domain.TeamComposition) *application_pb.TeamComposition {
	return &application_pb.TeamComposition{
		Id:                    comp.ID.String(),
		BaseContestantNumber:  int32(comp.BaseContestantNumber),
		BonusContestantNumber: int32(comp.BonusContestantNumber),
		MultiCategoryTeam:     comp.MultiCategoryTeam,
	}
}

func ageCategoryDomToPb(cat *domain.AgeCategory) *application_pb.AgeCategory {
	return &application_pb.AgeCategory{
		Id:          cat.ID.String(),
		Name:        cat.Name,
		MinAge:      int32(cat.MinAge),
		MaxAge:      int32(cat.MaxAge),
		Competition: nil,
	}
}

func ageCategoryListDomToPb(catList []domain.AgeCategory) []*application_pb.AgeCategory {
	var ageCatPbList []*application_pb.AgeCategory
	for _, propDom := range catList {
		ageCatPbList = append(ageCatPbList, ageCategoryDomToPb(&propDom))
	}

	return ageCatPbList
}

func competitionDomToPb(comp *domain.Competition) *application_pb.Competition {
	return &application_pb.Competition{
		Id:                           comp.ID.String(),
		Name:                         comp.Name,
		StartDate:                    comp.StartDate.Unix(),
		EndDate:                      comp.EndDate.Unix(),
		Gender:                       application_pb.Gender(comp.Gender),
		Type:                         application_pb.CompetitionType(comp.Type),
		Tiebreak:                     comp.Tiebreak,
		Address:                      addressDomToPb(&comp.Address),
		Organizer:                    sportsOrganizationDomToPb(&comp.Organizer),
		DelegationMemberPropositions: delegationMemberPropositionListDomToPb(comp.DelegationMemberPropositions),
		TeamComposition:              teamCompositionDomToPb(&comp.TeamComposition),
		AgeCategories:                ageCategoryListDomToPb(comp.AgeCategories),
	}
}

func competitionListDomToPb(compList []*domain.Competition) []*application_pb.Competition {
	var compPbList []*application_pb.Competition
	for _, compDom := range compList {
		compPbList = append(compPbList, competitionDomToPb(compDom))
	}

	return compPbList
}
