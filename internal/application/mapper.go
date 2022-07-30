package application

import "github.com/ybalcin/user-management/internal/domain"

func MapUserToDTO(user *domain.User) *UserDTO {
	return &UserDTO{
		Id:    user.Id.Hex(),
		Name:  user.Name,
		Email: user.Email,
	}
}
