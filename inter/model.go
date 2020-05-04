package inter

type Model interface {
	ToMap
	Id() int
}
