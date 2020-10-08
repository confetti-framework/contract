module github.com/lanvard/contract

go 1.15

require (
	github.com/lanvard/support v0.1.0
	github.com/lanvard/syslog v0.0.0-20201006215111-98d4d91dbaa8
)

replace github.com/lanvard/support v0.1.0 => ../support
replace github.com/lanvard/syslog v0.0.0-20201006215111-98d4d91dbaa8 => ../syslog
