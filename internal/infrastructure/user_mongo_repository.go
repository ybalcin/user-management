package infrastructure

import (
	"context"
	"github.com/ybalcin/user-management/internal/domain"
	"github.com/ybalcin/user-management/pkg/err"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (r *userMongoRepository) GetByEmail(ctx context.Context, email string) (*domain.User, *err.Error) {
	res := r.coll.FindOne(ctx, bson.M{
		"$and": []bson.M{
			{
				"email": email,
			},
			{
				"isDeleted": false,
			},
		},
	})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err.ThrowInternalServerError(res.Err())
	}

	user := new(domain.User)
	if e := res.Decode(user); e != nil {
		return nil, err.ThrowInternalServerError(e)
	}

	return user, nil
}

func (r *userMongoRepository) Update(ctx context.Context, id string, user *domain.User) *err.Error {
	objId, e := primitive.ObjectIDFromHex(id)
	if e != nil {
		return err.ThrowInternalServerError(e)
	}

	_, e = r.coll.UpdateByID(ctx, objId, bson.M{
		"$set": bson.M{
			"name":         user.Name,
			"email":        user.Email,
			"passwordHash": user.PasswordHash,
			"updatedAt":    user.UpdatedAt,
			"isDeleted":    user.IsDeleted,
		},
	})
	if e != nil {
		return err.ThrowInternalServerError(e)
	}

	return nil
}

func (r *userMongoRepository) GetById(ctx context.Context, id string) (*domain.User, *err.Error) {
	objId, e := primitive.ObjectIDFromHex(id)
	if e != nil {
		return nil, err.ThrowInternalServerError(e)
	}

	res := r.coll.FindOne(ctx, bson.M{
		"$and": []bson.M{
			{
				"_id": objId,
			},
			{
				"isDeleted": false,
			},
		},
	})
	if res.Err() != nil {
		if res.Err() == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err.ThrowInternalServerError(res.Err())
	}

	user := new(domain.User)
	if e := res.Decode(user); e != nil {
		return nil, err.ThrowInternalServerError(e)
	}

	return user, nil
}

func (r *userMongoRepository) GetAll(ctx context.Context) ([]domain.User, *err.Error) {
	crs, e := r.coll.Find(ctx, bson.M{"isDeleted": false})
	if e != nil {
		return nil, err.ThrowInternalServerError(e)
	}

	var users []domain.User
	if e = crs.All(ctx, &users); e != nil {
		return nil, err.ThrowInternalServerError(e)
	}

	if users == nil {
		return []domain.User{}, nil
	}

	return users, nil
}
