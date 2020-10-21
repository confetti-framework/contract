package inter

import (
	"github.com/lanvard/syslog/level"
)

type Error interface {
	error
	Level(inputLevel level.Level) Error
	GetLevel() level.Level
}
