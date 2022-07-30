package domain

import (
	"github.com/ybalcin/user-management/internal/shared/helper"
	"github.com/ybalcin/user-management/pkg/err"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Email        string             `bson:"email"`
	PasswordHash string             `bson:"passwordHash"`
	password     string             `bson:"-"`
}

func NewUser(name, email, password string) (*User, *err.Error) {
	u := &User{
		Id:       primitive.NewObjectID(),
		Name:     name,
		Email:    email,
		password: password,
	}

	if errs := u.validate(); len(errs) > 0 {
		return nil, err.ThrowValidationError(errs...)
	}

	passwordHash, e := helper.HashPassword(u.password)
	if e != nil {
		return nil, e
	}
	u.PasswordHash = passwordHash

	return u, nil
}

func (u *User) validate() []err.ValidationError {
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

	if helper.StrLength(u.password) <= 0 {
		errs = append(errs, UserPasswordCannotBeEmptyError)
	}

	if len(errs) <= 0 {
		return nil
	}

	return errs
}
