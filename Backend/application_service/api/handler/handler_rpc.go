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
}

func NewHandlerRpc(soService *service.SportsOrganisationService) *HandlerRpc {
	return &HandlerRpc{soService: soService}
}

func (h *HandlerRpc) RegisterSportsOrganisation(ctx context.Context, sOrganisation *application_pb.SportsOrganisation) (*application_pb.IdResponse, error) {
	id, _ := uuid.NewUUID()
	newSOrg := &domain.SportsOrganisation{
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

	return &application_pb.IdResponse{Id: id.String()}, nil
}
