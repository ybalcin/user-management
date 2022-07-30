package helper

import (
	"github.com/ybalcin/user-management/pkg/err"
	"golang.org/x/crypto/bcrypt"
	"regexp"
)

// StrLength returns string length by rune
func StrLength(str string) int {
	return len([]rune(str))
}

// IsEmailValid checks for email is valid
func IsEmailValid(email string) bool {
	match, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, email)
	return match
}

func HashPassword(password string) (string, *err.Error) {
	bytes, e := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if e != nil {
		return "", err.ThrowBadRequestError(e)
	}
	return string(bytes), nil
}
