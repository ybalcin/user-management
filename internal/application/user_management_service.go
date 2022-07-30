package application

import (
	"context"
	"errors"
	"github.com/ybalcin/user-management/internal/domain"
	"github.com/ybalcin/user-management/internal/domain/repositories"
	"github.com/ybalcin/user-management/internal/shared/helper"
	"github.com/ybalcin/user-management/pkg/err"
)

type UserManagementService interface {
	CreateNewUser(ctx context.Context, request *CreateUserRequest) (*UserDTO, *err.Error)
	UpdateUser(ctx context.Context, id string, request *UpdateUserRequest) (*UserDTO, *err.Error)
	DeleteUser(ctx context.Context, id string) *err.Error
	GetById(ctx context.Context, id string) (*UserDTO, *err.Error)
	GetAll(ctx context.Context) ([]UserDTO, *err.Error)
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

	// check for email exist already
	existUser, e := s.repo.GetByEmail(ctx, user.Email)
	if e != nil {
		return nil, e
	}

	if existUser != nil {
		return nil, err.ThrowForbiddenError(errors.New("user with that email already exists"))
	}

	// hash password
	passwordHash, e := helper.HashPassword(user.Password)
	if e != nil {
		return nil, e
	}
	user.PasswordHash = passwordHash

	if e = s.repo.Add(ctx, user); e != nil {
		return nil, e
	}

	return MapUserToDTO(user), nil
}

func (s *userManagementService) UpdateUser(ctx context.Context, id string, request *UpdateUserRequest) (*UserDTO, *err.Error) {
	user, e := s.getById(ctx, id)
	if e != nil {
		return nil, e
	}

	if helper.StrLength(request.Email) <= 0 {
		request.Email = user.Email
	}
	if helper.StrLength(request.Name) <= 0 {
		request.Name = user.Name
	}
	if helper.StrLength(request.Password) <= 0 {
		request.Password = user.PasswordHash
	} else {
		// hash password
		passwordHash, e := helper.HashPassword(user.Password)
		if e != nil {
			return nil, e
		}
		request.Password = passwordHash
	}

	newUser, e := domain.NewUser(request.Name, request.Email, request.Password)
	if e != nil {
		return nil, e
	}
	newUser.CreatedAt = user.CreatedAt
	newUser.PasswordHash = request.Password

	if e = s.repo.Update(ctx, id, newUser); e != nil {
		return nil, e
	}

	return MapUserToDTO(newUser), nil
}

func (s *userManagementService) DeleteUser(ctx context.Context, id string) *err.Error {
	user, e := s.getById(ctx, id)
	if e != nil {
		return e
	}

	user.Delete()

	if e = s.repo.Update(ctx, id, user); e != nil {
		return e
	}

	return nil
}

func (s *userManagementService) GetById(ctx context.Context, id string) (*UserDTO, *err.Error) {
	user, e := s.getById(ctx, id)
	if e != nil {
		return nil, e
	}

	return MapUserToDTO(user), nil
}

func (s *userManagementService) getById(ctx context.Context, id string) (*domain.User, *err.Error) {
	if helper.StrLength(id) <= 0 {
		return nil, err.ThrowBadRequestError(errors.New("id cannot be empty"))
	}

	user, e := s.repo.GetById(ctx, id)
	if e != nil {
		return nil, e
	}

	if user == nil {
		return nil, err.ThrowNotFoundError(errors.New("user with that id does not exist"))
	}

	return user, nil
}

func (s *userManagementService) GetAll(ctx context.Context) ([]UserDTO, *err.Error) {
	users, e := s.repo.GetAll(ctx)
	if e != nil {
		return nil, e
	}

	userDTOs := make([]UserDTO, len(users))
	for i, u := range users {
		userDTOs[i] = *MapUserToDTO(&u)
	}

	return userDTOs, nil
}
