package inter

import "net/url"

type UrlValues interface {
	Source() url.Values
	GetString(key string) string
	GetStringE(key string) (string, error)
	GetInt(key string) int
	GetIntE(key string) (int, error)
}
