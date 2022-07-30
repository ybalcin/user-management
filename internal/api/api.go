package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/ybalcin/user-management/internal/application"
	"github.com/ybalcin/user-management/internal/shared/response"
	"github.com/ybalcin/user-management/pkg/err"
)

type api struct {
	userManagementService application.UserManagementService

	app *fiber.App
}

func New(userManagementService application.UserManagementService) *api {
	app := fiber.New()

	return &api{
		userManagementService: userManagementService,
		app:                   app,
	}
}

func (a *api) App() *fiber.App {
	return a.app
}

//CreateNewUser
//@Summary Creates new user
//@Tags Users
//@Accept json
//@Produce json
//@Param application.CreateUserRequest body application.CreateUserRequest true "CreateNewUser body"
//@Success 200 {object} application.UserDTO "Returns UserDTO"
//@Failure 400 {object} err.Error "Returns err.Error"
//@Failure 500 {object} err.Error "Returns err.Error"
//@Router /users [post]
func (a *api) CreateNewUser(c *fiber.Ctx) error {
	req := new(application.CreateUserRequest)

	if e := c.BodyParser(req); e != nil {
		return response.New(c).Error(err.ThrowBadRequestError(e)).JSON()
	}

	userDTO, e := a.userManagementService.CreateNewUser(c.UserContext(), req)
	if e != nil {
		return response.New(c).Error(e).JSON()
	}

	return response.New(c).Data(userDTO).JSON()
}
