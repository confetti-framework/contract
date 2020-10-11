package inter

import (
	"github.com/lanvard/syslog"
)

type Severity = syslog.Priority
type Facility = syslog.Priority

type Logger interface {
	Log(severity Severity, message string)
	LogWith(severity Severity, message string, data interface{})
	Emergency(message string)
	EmergencyWith(message string, data interface{})
	Alert(message string)
	AlertWith(message string, data interface{})
	Critical(message string)
	CriticalWith(message string, data interface{})
	Error(message string)
	ErrorWith(message string, data interface{})
	Warning(message string)
	WarningWith(message string, data interface{})
	Notice(message string)
	NoticeWith(message string, data interface{})
	Info(message string)
	InfoWith(message string, data interface{})
	Debug(message string)
	DebugWith(message string, data interface{})
}
