package handler

import (
	"auth_service/core/domain"
	"auth_service/core/service"
	"auth_service/errors"
	auth_pb "common/proto/auth/generated"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AccountHandler struct {
	auth_pb.UnimplementedAuthServiceServer
	accountService *service.AccountService
}

func NewAccountHandler(accountService *service.AccountService) *AccountHandler {
	return &AccountHandler{accountService: accountService}
}

func (h *AccountHandler) Create(ctx context.Context, accountPb *auth_pb.Account) (*auth_pb.IdResponse, error) {
	newAcc := domain.Account{
		Email:    accountPb.GetEmail(),
		Password: accountPb.GetPassword(),
		Role: domain.Role{
			Name: accountPb.GetRole().GetName(),
		},
	}

	id, err := h.accountService.Create(&newAcc)
	if err != nil {
		switch err.(type) {
		case errors.ErrEmailTaken:
			return nil, status.Errorf(codes.AlreadyExists, err.Error())
		default:
			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}

	return &auth_pb.IdResponse{Id: id.String()}, nil
}

func (h *AccountHandler) Login(ctx context.Context, accountPb *auth_pb.Account) (*auth_pb.AccessToken, error) {

	account, err := h.accountService.Login(accountPb.Email, accountPb.Password)
	if err != nil {
		switch err.(type) {
		case errors.ErrNotFound:
			return nil, status.Errorf(codes.NotFound, err.Error())
		case errors.ErrBadCredentials:
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		default:
			return nil, status.Errorf(codes.Unknown, err.Error())
		}
	}

	accessToken, err := service.GenerateToken(&account)
	if err != nil {
		return nil, status.Errorf(codes.Unknown, err.Error())
	}

	return &auth_pb.AccessToken{Token: accessToken}, nil
}
func (h *AccountHandler) HasPermission(ctx context.Context, request *auth_pb.HasPermissionRequest) (*auth_pb.BoolMessage, error) {
	userinfo, err := ParseUserInfo(ctx)
	if err != nil {
		return nil, err
	}

	hasPermission, err := h.accountService.HasPermission(userinfo.Role, request.Permission)
	if err != nil {
		return nil, err
	}

	return &auth_pb.BoolMessage{Value: hasPermission}, nil
}

func (h *AccountHandler) DeleteAccounts(ctx context.Context, request *auth_pb.EmailList) (*auth_pb.EmptyMessage, error) {
	err := h.accountService.DeleteAccounts(request.Emails)
	if err != nil {
		return nil, err
	}
	return &auth_pb.EmptyMessage{}, nil
}
