package infrastructure

import (
	"context"
	"github.com/ybalcin/user-management/internal/domain"
	"github.com/ybalcin/user-management/pkg/err"
	"go.mongodb.org/mongo-driver/mongo"
)

const userCollection = "users"

type userMongoRepository struct {
	coll *mongo.Collection
}

func NewUserMongoRepository(db *mongo.Database) *userMongoRepository {
	return &userMongoRepository{coll: db.Collection(userCollection)}
}

func (r *userMongoRepository) Add(ctx context.Context, user *domain.User) *err.Error {
	if _, e := r.coll.InsertOne(ctx, user); e != nil {
		return err.ThrowInternalServerError(e)
	}

	return nil
}
