package inter

type Bootstrap interface {
	Bootstrap(app App) App
}
