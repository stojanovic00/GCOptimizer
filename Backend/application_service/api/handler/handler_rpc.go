package handler

import (
	"application_service/core/domain"
	"application_service/core/service"
	"application_service/errors"
	application_pb "common/proto/application/generated"
	"context"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type HandlerRpc struct {
	application_pb.UnimplementedApplicationServiceServer
	soService *service.SportsOrganisationService
	dmService *service.DelegationMemberService
}

func NewHandlerRpc(soService *service.SportsOrganisationService, dmService *service.DelegationMemberService) *HandlerRpc {
	return &HandlerRpc{soService: soService, dmService: dmService}
}

func (h *HandlerRpc) RegisterSportsOrganisation(ctx context.Context, sOrganisation *application_pb.SportsOrganisation) (*application_pb.IdMessage, error) {
	id, _ := uuid.NewUUID()
	newSOrg := &domain.SportsOrganization{
		ID:                             id,
		Name:                           sOrganisation.Name,
		Email:                          sOrganisation.Email,
		PhoneNumber:                    sOrganisation.PhoneNumber,
		ContactPersonFullName:          sOrganisation.ContactPersonFullName,
		CompetitionOrganisingPrivilege: false,
		Address: domain.Address{
			Country:      sOrganisation.Address.Country,
			City:         sOrganisation.Address.City,
			Street:       sOrganisation.Address.Street,
			StreetNumber: sOrganisation.Address.StreetNumber,
		},
	}

	id, err := h.soService.Create(newSOrg)

	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			return nil, status.Errorf(codes.NotFound, err.Error())
		case errors.ErrEmailTaken:
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		default:
			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}

	return &application_pb.IdMessage{Id: id.String()}, nil
}

func (h *HandlerRpc) GetLoggedSportsOrganisation(ctx context.Context, _ *application_pb.EmptyMessage) (*application_pb.SportsOrganisation, error) {
	userinfo, err := ParseUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	sportsOrganization, err := h.soService.GetByEmail(userinfo.Email)
	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			return nil, status.Errorf(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}
	return sportsOrganizationDomToPb(sportsOrganization), nil

}

func (h *HandlerRpc) RegisterJudge(ctx context.Context, judge *application_pb.Judge) (*application_pb.IdMessage, error) {
	userinfo, err := ParseUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	newJudge := &domain.Judge{
		DelegationMember: domain.DelegationMember{
			FullName: judge.DelegationMember.FullName,
			Email:    judge.DelegationMember.Email,
			Gender:   domain.Gender(judge.DelegationMember.Gender),
			Position: domain.DelegationMemberPosition{
				Name: judge.DelegationMember.Position.Name,
			},
		},
		LicenceType: domain.LicenceType(judge.LicenceType),
		LicenceName: judge.LicenceName,
	}
	id, err := h.dmService.RegisterJudge(newJudge, userinfo.Email)
	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			return nil, status.Errorf(codes.NotFound, err.Error())
		case errors.ErrEmailTaken:
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		default:
			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}

	return &application_pb.IdMessage{Id: id.String()}, nil
}

func (h *HandlerRpc) GetSportOrganisationJudges(ctx context.Context, _ *application_pb.EmptyMessage) (*application_pb.JudgesList, error) {
	userinfo, err := ParseUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	result, err := h.dmService.GetSportsOrganisationJudges(userinfo.Email)
	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			return nil, status.Errorf(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}

	return &application_pb.JudgesList{Judges: judgeListDomToPb(result)}, nil
}
func (h *HandlerRpc) RegisterContestant(ctx context.Context, contestant *application_pb.Contestant) (*application_pb.IdMessage, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterContestant not implemented")
}
func (h *HandlerRpc) GetSportOrganisationContestants(ctx context.Context, _ *application_pb.EmptyMessage) (*application_pb.ContestantList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSportOrganisationContestants not implemented")
}
