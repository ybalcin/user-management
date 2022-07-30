package api

func SetRoutes(api *api) {
	api.app.Post("/users", api.CreateNewUser)
}
