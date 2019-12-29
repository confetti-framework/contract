package application

type App interface {
	Make(abstract interface{}) interface{}
}
