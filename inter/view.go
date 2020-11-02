package inter

type View interface {
	App() App
	Template() string
}
