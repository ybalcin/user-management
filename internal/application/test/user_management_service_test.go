package test

import (
	"context"
	"errors"
	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	. "github.com/ybalcin/user-management/internal/application"
	"github.com/ybalcin/user-management/internal/domain"
	"github.com/ybalcin/user-management/pkg/err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func setupMockUserRepository(t *testing.T) *MockUserRepository {
	return NewMockUserRepository(gomock.NewController(t))
}

type mockCall func(params ...interface{}) *gomock.Call

func TestUserManagementService_CreateNewUser(t *testing.T) {
	repo := setupMockUserRepository(t)
	service := NewUserManagementService(repo)
	ctx := context.Background()

	name := gofakeit.Name()
	email := gofakeit.Email()
	password := gofakeit.Name()

	cases := []struct {
		tt          string
		request     *CreateUserRequest
		getByEmail  mockCall
		add         mockCall
		expectedErr *err.Error
	}{
		{
			"success",
			&CreateUserRequest{
				Name:     name,
				Email:    email,
				Password: password,
			},
			func(params ...interface{}) *gomock.Call {
				email := params[0]
				return repo.EXPECT().GetByEmail(ctx, email).Return(nil, nil)
			},
			func(params ...interface{}) *gomock.Call {
				return repo.EXPECT().Add(ctx, gomock.Any()).Return(nil)
			},
			nil,
		},
		{
			"should return UserExistWithThatEmailError if user exist with that email",
			&CreateUserRequest{
				Name:     name,
				Email:    email,
				Password: password,
			},
			func(params ...interface{}) *gomock.Call {
				email := params[0]
				return repo.EXPECT().GetByEmail(ctx, email).Return(&domain.User{Email: email.(string)}, nil)
			},
			nil,
			err.ThrowForbiddenError(UserExistWithThatEmailError),
		},
		{
			"should return validation error if request is not valid",
			&CreateUserRequest{},
			nil,
			nil,
			err.ThrowValidationError(domain.UserNameCannotBeEmptyError, domain.UserEmailCannotBeEmptyError, domain.UserPasswordCannotBeEmptyError),
		},
	}

	for _, c := range cases {
		t.Run(c.tt, func(t *testing.T) {
			if c.getByEmail != nil {
				c.getByEmail(c.request.Email)
			}
			if c.add != nil {
				c.add()
			}

			dto, e := service.CreateNewUser(ctx, c.request)

			if c.expectedErr != nil {
				assert.Equal(t, c.expectedErr, e)
				assert.Nil(t, dto)
			} else {
				assert.Nil(t, e)
				assert.NotNil(t, dto)
			}
		})
	}
}

func TestUserManagementService_UpdateUser(t *testing.T) {
	repo := setupMockUserRepository(t)
	service := NewUserManagementService(repo)
	ctx := context.Background()

	name := gofakeit.Name()
	email := gofakeit.Email()
	password := gofakeit.Name()

	cases := []struct {
		tt          string
		request     *UpdateUserRequest
		id          string
		getById     mockCall
		update      mockCall
		expectedErr *err.Error
	}{
		{
			"success",
			&UpdateUserRequest{
				Name:     name,
				Email:    email,
				Password: password,
			},
			primitive.NewObjectID().Hex(),
			func(params ...interface{}) *gomock.Call {
				id := params[0]
				return repo.EXPECT().GetById(ctx, id).Return(&domain.User{}, nil)
			},
			func(params ...interface{}) *gomock.Call {
				id := params[0]
				return repo.EXPECT().Update(ctx, id, gomock.Any()).Return(nil)
			},
			nil,
		},
		{
			"should return IdCannotBeEmptyError id is empty",
			&UpdateUserRequest{
				Name:     name,
				Email:    email,
				Password: password,
			},
			"",
			nil,
			nil,
			err.ThrowBadRequestError(IdCannotBeEmptyError),
		},
		{
			"should return UserIsNotExistWithThatIdError if user not exist with that id",
			&UpdateUserRequest{
				Name:     name,
				Email:    email,
				Password: password,
			},
			primitive.NewObjectID().Hex(),
			func(params ...interface{}) *gomock.Call {
				id := params[0]
				return repo.EXPECT().GetById(ctx, id).Return(nil, nil)
			},
			nil,
			err.ThrowNotFoundError(UserIsNotExistWithThatIdError),
		},
		{
			"should return error if update returns error",
			&UpdateUserRequest{
				Name:     name,
				Email:    email,
				Password: password,
			},
			primitive.NewObjectID().Hex(),
			func(params ...interface{}) *gomock.Call {
				id := params[0]
				return repo.EXPECT().GetById(ctx, id).Return(&domain.User{}, nil)
			},
			func(params ...interface{}) *gomock.Call {
				id := params[0]
				return repo.EXPECT().Update(ctx, id, gomock.Any()).Return(err.ThrowInternalServerError(errors.New("")))
			},
			err.ThrowInternalServerError(errors.New("")),
		},
	}

	for _, c := range cases {
		t.Run(c.tt, func(t *testing.T) {
			if c.getById != nil {
				c.getById(c.id)
			}
			if c.update != nil {
				c.update(c.id)
			}

			dto, e := service.UpdateUser(ctx, c.id, c.request)

			if c.expectedErr != nil {
				assert.Equal(t, c.expectedErr, e)
				assert.Nil(t, dto)
			} else {
				assert.Nil(t, e)
				assert.NotNil(t, dto)
			}
		})
	}
}

func TestUserManagementService_DeleteUser(t *testing.T) {
	repo := setupMockUserRepository(t)
	service := NewUserManagementService(repo)
	ctx := context.Background()

	name := gofakeit.Name()
	email := gofakeit.Email()
	password := gofakeit.Name()

	id := primitive.NewObjectID()

	model := &domain.User{
		Id:           id,
		Name:         name,
		Email:        email,
		PasswordHash: password,
	}

	cases := []struct {
		tt       string
		id       string
		getById  mockCall
		update   mockCall
		expected *err.Error
	}{
		{
			"success",
			id.Hex(),
			func(params ...interface{}) *gomock.Call {
				return repo.EXPECT().GetById(ctx, id.Hex()).Return(model, nil)
			},
			func(params ...interface{}) *gomock.Call {
				model.Delete()
				return repo.EXPECT().Update(ctx, id.Hex(), model).Return(nil)
			},
			nil,
		},
		{
			"should return error if getById returns error",
			id.Hex(),
			func(params ...interface{}) *gomock.Call {
				return repo.EXPECT().GetById(ctx, id.Hex()).Return(nil, err.ThrowInternalServerError(errors.New("")))
			},
			nil,
			err.ThrowInternalServerError(errors.New("")),
		},
		{
			"should return error if update returns error",
			id.Hex(),
			func(params ...interface{}) *gomock.Call {
				return repo.EXPECT().GetById(ctx, id.Hex()).Return(model, nil)
			},
			func(params ...interface{}) *gomock.Call {
				model.Delete()
				return repo.EXPECT().Update(ctx, id.Hex(), model).Return(err.ThrowInternalServerError(errors.New("")))
			},
			err.ThrowInternalServerError(errors.New("")),
		},
	}

	for _, c := range cases {
		t.Run(c.tt, func(t *testing.T) {
			if c.getById != nil {
				c.getById(c.id)
			}
			if c.update != nil {
				c.update(c.id)
			}

			e := service.DeleteUser(ctx, c.id)
			if c.expected != nil {
				assert.Equal(t, c.expected, e)
			} else {
				assert.Nil(t, e)
			}
		})
	}
}
