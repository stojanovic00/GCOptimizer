package handler

import (
	"application_service/core/domain"
	application_pb "common/proto/application/generated"
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
