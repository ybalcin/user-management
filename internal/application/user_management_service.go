package application

import (
	"context"
	"github.com/ybalcin/user-management/internal/domain"
	"github.com/ybalcin/user-management/internal/domain/repositories"
	"github.com/ybalcin/user-management/pkg/err"
)

type UserManagementService interface {
	CreateNewUser(ctx context.Context, request *CreateUserRequest) (*UserDTO, *err.Error)
}

type userManagementService struct {
	repo repositories.UserRepository
}

func NewUserManagementService(repo repositories.UserRepository) *userManagementService {
	return &userManagementService{repo: repo}
}

func (s *userManagementService) CreateNewUser(ctx context.Context, request *CreateUserRequest) (*UserDTO, *err.Error) {
	user, e := domain.NewUser(request.Name, request.Email, request.Password)
	if e != nil {
		return nil, e
	}
	// check for email exist before

	if e = s.repo.Add(ctx, user); e != nil {
		return nil, e
	}

	return &UserDTO{
		Id:    user.Id.Hex(),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
