package inter

type Value interface {
	String() string
	StringE() (string, error)
	Strings() []string
	StringsE() ([]string, error)
	Number() int
	NumberE() (int, error)
	Numbers() []int
	NumbersE() ([]int, error)
	Empty() bool
	Present() bool
	Split(separator string) []string
}
