package api

func (a *api) Listen() error {
	SetRoutes(a)

	if e := a.app.Listen(":8080"); e != nil {
		return e
	}

	return nil
}
