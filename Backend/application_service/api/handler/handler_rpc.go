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
	"time"
)

type HandlerRpc struct {
	application_pb.UnimplementedApplicationServiceServer
	soService   *service.SportsOrganisationService
	dmService   *service.DelegationMemberService
	compService *service.CompetitionService
	appService  *service.ApplicationService
}

func NewHandlerRpc(soService *service.SportsOrganisationService, dmService *service.DelegationMemberService, compService *service.CompetitionService, appService *service.ApplicationService) *HandlerRpc {
	return &HandlerRpc{soService: soService, dmService: dmService, compService: compService, appService: appService}
}

func (h *HandlerRpc) RegisterSportsOrganisation(ctx context.Context, sOrganisation *application_pb.SportsOrganisation) (*application_pb.IdMessage, error) {
	id, _ := uuid.NewUUID()
	newSOrg := sportsOrganizationPbToDom(sOrganisation)

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
		DelegationMember: *delegationMemberPbToDom(judge.DelegationMember),
		LicenceType:      domain.LicenceType(judge.LicenceType),
		LicenceName:      judge.LicenceName,
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
	userinfo, err := ParseUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	newContestant := &domain.Contestant{
		DelegationMember: *delegationMemberPbToDom(contestant.DelegationMember),
		DateOfBirth:      time.Unix(contestant.DateOfBirth, 0),
	}

	id, err := h.dmService.RegisterContestant(newContestant, userinfo.Email)
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

func (h *HandlerRpc) GetSportOrganisationContestants(ctx context.Context, _ *application_pb.EmptyMessage) (*application_pb.ContestantList, error) {
	userinfo, err := ParseUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	result, err := h.dmService.GetSportsOrganisationContestants(userinfo.Email)
	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			return nil, status.Errorf(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}

	return &application_pb.ContestantList{Contestants: contestantListDomToPb(result)}, nil
}

func (h *HandlerRpc) CreateCompetition(ctx context.Context, competition *application_pb.Competition) (*application_pb.IdMessage, error) {
	userinfo, err := ParseUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	newComp := competitionPbToDom(competition)

	id, err := h.compService.Create(newComp, userinfo.Email)
	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			return nil, status.Errorf(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}

	return &application_pb.IdMessage{Id: id.String()}, nil
}
func (h *HandlerRpc) GetAllCompetitions(ctx context.Context, _ *application_pb.EmptyMessage) (*application_pb.CompetitionList, error) {
	comps, err := h.compService.GetAll()
	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			return nil, status.Errorf(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}
	return &application_pb.CompetitionList{Competitions: competitionListDomToPb(comps)}, nil
}
func (h *HandlerRpc) GetCompetitionById(ctx context.Context, compId *application_pb.IdMessage) (*application_pb.Competition, error) {
	id, _ := uuid.Parse(compId.Id)
	comp, err := h.compService.GetById(id)
	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			return nil, status.Errorf(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}
	return competitionDomToPb(comp), nil
}
func (h *HandlerRpc) AddAgeCategory(ctx context.Context, request *application_pb.AddAgeCategoryRequest) (*application_pb.IdMessage, error) {
	compId, _ := uuid.Parse(request.CompetitionId)

	id, err := h.compService.AddAgeCategory(ageCategoryPbToDom(request.AgeCategory), compId)
	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			return nil, status.Errorf(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}
	return &application_pb.IdMessage{Id: id.String()}, nil
}
func (h *HandlerRpc) AddDelegationMemberProposition(ctx context.Context, request *application_pb.AddDelegationMemberPropositionRequest) (*application_pb.IdMessage, error) {
	compId, _ := uuid.Parse(request.CompetitionId)

	id, err := h.compService.AddDelegationMemberProposition(delegationMemberPropositionPbToDom(request.DelegationMemberProposition), compId)
	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			return nil, status.Errorf(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}
	return &application_pb.IdMessage{Id: id.String()}, nil
}
func (h *HandlerRpc) CreateJudgeApplication(ctx context.Context, request *application_pb.CreateJudgeApplicationRequest) (*application_pb.IdMessage, error) {
	compId, _ := uuid.Parse(request.CompetitionId)
	newApp := judgeApplicationRequestPbToDom(request)

	id, err := h.appService.CreateJudgeApplication(newApp, compId)
	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			return nil, status.Errorf(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}
	return &application_pb.IdMessage{Id: id.String()}, nil
}

func (h *HandlerRpc) GetAllJudgeApplications(ctx context.Context, compId *application_pb.IdMessage) (*application_pb.JudgeApplicationList, error) {
	cmpId, _ := uuid.Parse(compId.Id)

	applications, err := h.appService.GetAllJudgeApplications(cmpId)
	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			return nil, status.Errorf(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}
	return &application_pb.JudgeApplicationList{JudgeApplications: judgeApplicationListDomToPb(applications)}, nil
}

func (h *HandlerRpc) CreateContestantApplication(ctx context.Context, request *application_pb.CreateContestantApplicationRequest) (*application_pb.IdMessage, error) {
	compId, _ := uuid.Parse(request.CompetitionId)
	newApp := contestantApplicationRequestPbToDom(request)
	newApp.CompetitionID = compId

	id, err := h.appService.CreateContestantApplication(newApp)
	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			return nil, status.Errorf(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}
	return &application_pb.IdMessage{Id: id.String()}, nil
}
func (h *HandlerRpc) GetAllContestantApplications(ctx context.Context, compId *application_pb.IdMessage) (*application_pb.ContestantApplicationList, error) {
	cmpId, _ := uuid.Parse(compId.Id)

	applications, err := h.appService.GetAllContestantApplications(cmpId)
	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			return nil, status.Errorf(codes.NotFound, err.Error())
		default:
			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}
	return &application_pb.ContestantApplicationList{ContestantApplications: contestantApplicationListDomToPb(applications)}, nil
}

func (h *HandlerRpc) GetSportsOrganisationByName(ctx context.Context, request *application_pb.GetSportsOrganisationByNameRequest) (*application_pb.SportsOrganisation, error) {

	sportsOrganization, err := h.soService.GetByName(request.Name)
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
