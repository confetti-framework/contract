package inter

type RegisterServiceProvider interface {
	Register(app *Application) *Application
}
