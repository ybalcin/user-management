package api

func SetRoutes(api *api) {
	group := api.app.Group("/users")

	group.Put("", api.CreateNewUser)
	group.Patch("/:id", api.UpdateUser)
	group.Delete("/:id", api.DeleteUser)
	group.Get("/:id", api.GetUserById)
	group.Get("", api.GetAllUsers)
}
