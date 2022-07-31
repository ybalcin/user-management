package test

import (
	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	. "github.com/ybalcin/user-management/internal/domain"
	"github.com/ybalcin/user-management/pkg/err"
	"testing"
)

func TestUser_Validate(t *testing.T) {
	cases := []struct {
		tt       string
		user     *User
		expected []err.ValidationError
	}{
		{
			"should return UserEmailInvalidError if email is invalid",
			&User{
				Name:     gofakeit.Name(),
				Email:    gofakeit.Name(),
				Password: gofakeit.Name(),
			},
			[]err.ValidationError{
				UserEmailInvalidError,
			},
		},
		{
			"should return UserNameCannotBeEmptyError if name is empty or whitespace",
			&User{
				Email:    gofakeit.Email(),
				Password: gofakeit.Name(),
			},
			[]err.ValidationError{
				UserNameCannotBeEmptyError,
			},
		},
		{
			"should return UserPasswordCannotBeEmptyError if password is empty or whitespace",
			&User{
				Name:  gofakeit.Name(),
				Email: gofakeit.Email(),
			},
			[]err.ValidationError{
				UserPasswordCannotBeEmptyError,
			},
		},
		{
			"success",
			&User{
				Name:     gofakeit.Name(),
				Email:    gofakeit.Email(),
				Password: gofakeit.Name(),
			},
			nil,
		},
	}

	for _, c := range cases {
		t.Run(c.tt, func(t *testing.T) {
			errs := c.user.Validate()
			assert.Equal(t, c.expected, errs)
		})
	}
}

func TestUser_Delete(t *testing.T) {
	u := &User{}
	u.Delete()

	assert.Equal(t, true, u.IsDeleted)
}
