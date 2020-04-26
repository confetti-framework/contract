package inter

type RegisterServiceProvider interface {
	Register(app App) App
}
