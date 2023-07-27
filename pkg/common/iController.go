package common

type IController interface {
	ConfigureRoutes(app *FiberApp)
}
