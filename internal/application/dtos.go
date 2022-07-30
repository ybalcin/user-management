package application

type (
	CreateUserRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UpdateUserRequest struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UserDTO struct {
		Id    string `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}
)
