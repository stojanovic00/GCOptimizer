package handler

import (
	"api_gateway/dto"
	"auth_service/api/middleware"
	application_pb "common/proto/application/generated"
	auth_pb "common/proto/auth/generated"
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
)

type ApplicationHandler struct {
	appClient  application_pb.ApplicationServiceClient
	authClient auth_pb.AuthServiceClient
}

func NewApplicationHandler(appClient application_pb.ApplicationServiceClient, authClient auth_pb.AuthServiceClient) *ApplicationHandler {
	return &ApplicationHandler{appClient: appClient, authClient: authClient}
}

func (h *ApplicationHandler) RegisterSportsOrganisation(ctx *gin.Context) {
	var registrationDto dto.SportsOrganisationRegistration

	err := ctx.ShouldBindJSON(&registrationDto)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	//Just in case frontend messed up
	registrationDto.Account.Email = registrationDto.SportsOrganisation.Email

	accId, err := h.authClient.Create(context.TODO(), &registrationDto.Account)
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.AlreadyExists:
				ctx.JSON(http.StatusConflict, grpcError.Message())
				return

			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	soId, err := h.appClient.RegisterSportsOrganisation(context.TODO(), &registrationDto.SportsOrganisation)
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.AlreadyExists:
				ctx.JSON(http.StatusConflict, grpcError.Message())
				return
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, grpcError.Message())
				return

			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, dto.SportsOrganisationRegistrationResponse{AccountID: accId.Id, SportsOrganisationID: soId.Id})
}

func (h *ApplicationHandler) GetLoggedSportsOrganisation(ctx *gin.Context) {
	ctxWithUserInfo, err := middleware.GetGrpcContextWithUserInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		ctx.Abort()
		return
	}

	sportsOrg, err := h.appClient.GetLoggedSportsOrganisation(ctxWithUserInfo, &application_pb.EmptyMessage{})
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, grpcError.Message())
				return

			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, sportsOrg)
}

func (h *ApplicationHandler) RegisterJudge(ctx *gin.Context) {
	var newJudge application_pb.Judge

	err := ctx.ShouldBindJSON(&newJudge)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctxWithUserInfo, err := middleware.GetGrpcContextWithUserInfo(ctx)
	id, err := h.appClient.RegisterJudge(ctxWithUserInfo, &newJudge)

	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, grpcError.Message())
				return
			case codes.AlreadyExists:
				ctx.JSON(http.StatusConflict, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, id)
}
func (h *ApplicationHandler) GetSportOrganisationJudges(ctx *gin.Context) {
	ctxWithUserInfo, err := middleware.GetGrpcContextWithUserInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		ctx.Abort()
		return
	}

	judges, err := h.appClient.GetSportOrganisationJudges(ctxWithUserInfo, &application_pb.EmptyMessage{})

	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, judges.Judges)
}
func (h *ApplicationHandler) RegisterContestant(ctx *gin.Context) {
	var newContestant application_pb.Contestant

	err := ctx.ShouldBindJSON(&newContestant)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctxWithUserInfo, err := middleware.GetGrpcContextWithUserInfo(ctx)
	id, err := h.appClient.RegisterContestant(ctxWithUserInfo, &newContestant)

	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, grpcError.Message())
				return
			case codes.AlreadyExists:
				ctx.JSON(http.StatusConflict, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, id)
}
func (h *ApplicationHandler) GetSportOrganisationContestants(ctx *gin.Context) {
	ctxWithUserInfo, err := middleware.GetGrpcContextWithUserInfo(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		ctx.Abort()
		return
	}

	contestants, err := h.appClient.GetSportOrganisationContestants(ctxWithUserInfo, &application_pb.EmptyMessage{})

	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, contestants.Contestants)
}
func (h *ApplicationHandler) CreateCompetition(ctx *gin.Context) {
	var newCompetition application_pb.Competition

	err := ctx.ShouldBindJSON(&newCompetition)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctxWithUserInfo, err := middleware.GetGrpcContextWithUserInfo(ctx)
	id, err := h.appClient.CreateCompetition(ctxWithUserInfo, &newCompetition)

	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, id)
}
func (h *ApplicationHandler) GetAllCompetitions(ctx *gin.Context) {
	comps, err := h.appClient.GetAllCompetitions(context.Background(), &application_pb.EmptyMessage{})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, comps.Competitions)
}
func (h *ApplicationHandler) GetCompetitionById(ctx *gin.Context) {
	id := ctx.Param("id")

	comp, err := h.appClient.GetCompetitionById(context.Background(), &application_pb.IdMessage{Id: id})
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, comp)
}
func (h *ApplicationHandler) AddAgeCategory(ctx *gin.Context) {
	compId := ctx.Param("id")

	var newCat application_pb.AgeCategory

	err := ctx.ShouldBindJSON(&newCat)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.appClient.AddAgeCategory(context.Background(), &application_pb.AddAgeCategoryRequest{
		AgeCategory:   &newCat,
		CompetitionId: compId,
	})

	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, id)

}
func (h *ApplicationHandler) AddDelegationMemberProposition(ctx *gin.Context) {
	compId := ctx.Param("id")

	var newProp application_pb.DelegationMemberProposition

	err := ctx.ShouldBindJSON(&newProp)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.appClient.AddDelegationMemberProposition(context.Background(), &application_pb.AddDelegationMemberPropositionRequest{
		DelegationMemberProposition: &newProp,
		CompetitionId:               compId,
	})

	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, id)
}
func (h *ApplicationHandler) CreateJudgeApplication(ctx *gin.Context) {
	compId := ctx.Param("id")

	var newAppReq application_pb.CreateJudgeApplicationRequest
	newAppReq.CompetitionId = compId

	err := ctx.ShouldBindJSON(&newAppReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.appClient.CreateJudgeApplication(context.Background(), &newAppReq)

	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, id)
}

func (h *ApplicationHandler) GetAllJudgeApplications(ctx *gin.Context) {
	compId := ctx.Param("id")

	applications, err := h.appClient.GetAllJudgeApplications(context.Background(), &application_pb.IdMessage{Id: compId})
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, applications.JudgeApplications)
}
func (h *ApplicationHandler) CreateContestantApplication(ctx *gin.Context) {
	compId := ctx.Param("id")

	var newAppReq application_pb.CreateContestantApplicationRequest
	newAppReq.CompetitionId = compId

	err := ctx.ShouldBindJSON(&newAppReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.appClient.CreateContestantApplication(context.Background(), &newAppReq)

	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusCreated, id)

}
func (h *ApplicationHandler) GetAllContestantApplications(ctx *gin.Context) {
	compId := ctx.Param("id")

	applications, err := h.appClient.GetAllContestantApplications(context.Background(), &application_pb.IdMessage{Id: compId})
	if err != nil {
		grpcError, ok := status.FromError(err)
		if ok {
			switch grpcError.Code() {
			case codes.NotFound:
				ctx.JSON(http.StatusNotFound, grpcError.Message())
				return
			}
		}
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, applications.ContestantApplications)
}
