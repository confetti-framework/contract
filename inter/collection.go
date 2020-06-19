package inter

type Collection interface {
	Source() []Value
	First() Value
	Push(item interface{}) Collection
	Reverse() Collection
	Contains(search interface{}) bool
	Len() int
}
