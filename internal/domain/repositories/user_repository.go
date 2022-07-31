package repositories

import (
	"context"
	"github.com/ybalcin/user-management/internal/domain"
	"github.com/ybalcin/user-management/pkg/err"
)

// cd internal/domain/repositories
// mockgen --source=user_repository.go --destination=../../application/test/user_repository_mock.go --package=test

type UserRepository interface {
	Add(ctx context.Context, user *domain.User) *err.Error
	GetByEmail(ctx context.Context, email string) (*domain.User, *err.Error)
	Update(ctx context.Context, id string, user *domain.User) *err.Error
	GetById(ctx context.Context, id string) (*domain.User, *err.Error)
	GetAll(ctx context.Context) ([]domain.User, *err.Error)
}
