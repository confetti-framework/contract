package inter

import "io"

type ExitCode int

const (
	Success ExitCode = iota
	Failure
	Help
	Continue
)

type Command interface {
	Name() string
	Description() string
	Handle(app App, writer io.Writer) ExitCode
}
