package inter

type Bootstrap interface {
	Bootstrap(app Container) Container
}
