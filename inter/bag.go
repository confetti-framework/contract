package inter

type Bag interface {
	Source() map[string]Collection
	Get(key string) Value
	GetMany(key string) Collection
	Set(key string, value Value)
	Push(key string, value interface{})
	Del(key string)
	Merge(bags ...Bag) Bag
}
