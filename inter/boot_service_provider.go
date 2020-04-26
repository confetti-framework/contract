package inter

type BootServiceProvider interface {
	Boot(app App) App
}
