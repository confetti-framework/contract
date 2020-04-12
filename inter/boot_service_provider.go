package inter

type BootServiceProvider interface {
	Boot(app *Application) *Application
}
