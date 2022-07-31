package domain

import (
	"github.com/ybalcin/user-management/internal/shared/helper"
	"github.com/ybalcin/user-management/pkg/err"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"passwordHash"`
	Password     string             `bson:"-"`
	CreatedAt    time.Time          `bson:"createdAt"`
	UpdatedAt    time.Time          `bson:"updatedAt"`
	IsDeleted    bool               `bson:"isDeleted"`
}

func NewUser(name, email, password string) (*User, *err.Error) {
	u := &User{
		Id:        primitive.NewObjectID(),
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if errs := u.Validate(); len(errs) > 0 {
		return nil, err.ThrowValidationError(errs...)
	}

	return u, nil
}

func (u *User) Validate() []err.ValidationError {
	var errs []err.ValidationError

	if helper.StrLength(u.Name) <= 0 {
		errs = append(errs, UserNameCannotBeEmptyError)
	}

	if helper.StrLength(u.Email) <= 0 {
		errs = append(errs, UserEmailCannotBeEmptyError)
	} else {
		if !helper.IsEmailValid(u.Email) {
			errs = append(errs, UserEmailInvalidError)
		}
	}

	if helper.StrLength(u.Password) <= 0 {
		errs = append(errs, UserPasswordCannotBeEmptyError)
	}

	if len(errs) <= 0 {
		return nil
	}

	return errs
}

func (u *User) Delete() {
	u.IsDeleted = true
	u.UpdatedAt = time.Now().UTC()
}
