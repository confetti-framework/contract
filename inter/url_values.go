package inter

import "net/url"

type UrlValues interface {
	Source() url.Values
	String(key string) string
	StringE(key string) (string, error)
	Strings(key string) []string
	StringsE(key string) ([]string, error)
	Number(key string) int
	NumberE(key string) (int, error)
	Numbers(key string) []int
	NumbersE(key string) ([]int, error)
}
