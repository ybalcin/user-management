package repositories

import (
	"context"
	"github.com/ybalcin/user-management/internal/domain"
	"github.com/ybalcin/user-management/pkg/err"
)

type UserRepository interface {
	Add(ctx context.Context, user *domain.User) *err.Error
}
