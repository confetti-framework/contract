module github.com/confetti-framework/contract

go 1.15

require (
	github.com/confetti-framework/support v0.1.0
	github.com/confetti-framework/syslog v0.0.0-20201116213126-66df7a6162c3
)

replace github.com/confetti-framework/support v0.1.0 => ../support
replace github.com/confetti-framework/syslog v0.0.0-20201116213126-66df7a6162c3 => ../syslog
