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
//@Router /users [put]
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

//UpdateUser
//@Summary Updates new user
//@Tags Users
//@Accept json
//@Produce json
//@Param application.UpdateUserRequest body application.UpdateUserRequest true "UpdateUser body"
//@Param id path string true "id"
//@Success 200 {object} application.UserDTO "Returns UserDTO"
//@Failure 400 {object} err.Error "Returns err.Error"
//@Failure 500 {object} err.Error "Returns err.Error"
//@Router /users/{id} [patch]
func (a *api) UpdateUser(c *fiber.Ctx) error {
	req := new(application.UpdateUserRequest)
	id := c.Params("id")

	if e := c.BodyParser(req); e != nil {
		return response.New(c).Error(err.ThrowBadRequestError(e)).JSON()
	}

	userDTO, e := a.userManagementService.UpdateUser(c.UserContext(), id, req)
	if e != nil {
		return response.New(c).Error(e).JSON()
	}

	return response.New(c).Data(userDTO).JSON()
}

//DeleteUser
//@Summary Deletes user
//@Tags Users
//@Accept json
//@Produce json
//@Param id path string true "id"
//@Failure 400 {object} err.Error "Returns err.Error"
//@Failure 500 {object} err.Error "Returns err.Error"
//@Router /users/{id} [delete]
func (a *api) DeleteUser(c *fiber.Ctx) error {
	id := c.Params("id")

	if e := a.userManagementService.DeleteUser(c.UserContext(), id); e != nil {
		return response.New(c).Error(e).JSON()
	}

	return response.New(c).JSON()
}

//GetUserById
//@Summary Gets user by id
//@Tags Users
//@Accept json
//@Produce json
//@Param id path string true "id"
//@Success 200 {object} application.UserDTO "Returns UserDTO"
//@Failure 400 {object} err.Error "Returns err.Error"
//@Failure 500 {object} err.Error "Returns err.Error"
//@Router /users/{id} [get]
func (a *api) GetUserById(c *fiber.Ctx) error {
	id := c.Params("id")

	userDTO, e := a.userManagementService.GetById(c.UserContext(), id)
	if e != nil {
		return response.New(c).Error(e).JSON()
	}

	return response.New(c).Data(userDTO).JSON()
}

//GetAllUsers
//@Summary Gets all users
//@Tags Users
//@Accept json
//@Produce json
//@Success 200 {object} []application.UserDTO "Returns UserDTO array"
//@Failure 400 {object} err.Error "Returns err.Error"
//@Failure 500 {object} err.Error "Returns err.Error"
//@Router /users [get]
func (a *api) GetAllUsers(c *fiber.Ctx) error {
	userDTOs, e := a.userManagementService.GetAll(c.UserContext())
	if e != nil {
		return response.New(c).Error(e).JSON()
	}

	return response.New(c).Data(userDTOs).JSON()
}
