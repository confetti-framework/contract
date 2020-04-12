package inter

type Bootstrap interface {
	Bootstrap(app *Application) *Application
}
