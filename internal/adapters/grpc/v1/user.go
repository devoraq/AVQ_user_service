package user

import (
	"context"

	v1 "github.com/DENFNC/awq_user_service/api/gen/go/user/v1"
	"github.com/DENFNC/awq_user_service/internal/adapters/dto"
	"github.com/DENFNC/awq_user_service/internal/utils/mapping"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService interface {
	Create(
		ctx context.Context,
		user *dto.CreateUserDTO,
	) (string, error)
}

type serverAPI struct {
	v1.UnimplementedUserServiceServer
	svc UserService
}

func NewUser(svc UserService) *serverAPI {
	return &serverAPI{
		svc: svc,
	}
}

func (api *serverAPI) Register(grpc *grpc.Server) {
	v1.RegisterUserServiceServer(grpc, api)
}

func (api *serverAPI) CreateUser(
	ctx context.Context,
	req *v1.CreateUserRequest,
) (*v1.CreateUserResponse, error) {
	var dtoUser dto.CreateUserDTO

	if err := mapping.MapStruct(req, &dtoUser); err != nil {
		return nil, status.Error(codes.Internal, "Couldn't create user")
	}

	_, err := api.svc.Create(ctx, &dtoUser)
	if err != nil {
		return nil, status.Error(codes.Internal, "Couldn't create user")
	}

	return &v1.CreateUserResponse{
		UserProfile: &v1.UserProfile{},
	}, nil
}

func (api *serverAPI) DeleteUser(
	ctx context.Context,
	req *v1.DeleteUserRequest,
) (*v1.DeleteUserResponse, error) {
	panic("implement me!")
}

func (api *serverAPI) FetchUser(
	ctx context.Context,
	req *v1.FetchUserRequest,
) (*v1.FetchUserResponse, error) {
	panic("implement me!")
}

func (api *serverAPI) ListUsers(
	ctx context.Context,
	req *v1.ListUsersRequest,
) (*v1.ListUsersResponse, error) {
	panic("implement me!")
}
