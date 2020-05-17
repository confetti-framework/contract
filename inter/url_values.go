package inter

type UrlValues interface {
	Get(key string) Value
	GetMulti(key string) []Value
	Set(key string, value Value)
	Add(key string, value Value)
	Del(key string)
}
