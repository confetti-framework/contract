module github.com/lanvard/contract

go 1.15

require (
	github.com/lanvard/support v0.1.0
	github.com/lanvard/syslog v0.0.0-20201116213126-66df7a6162c3
)

replace github.com/lanvard/support v0.1.0 => ../support
replace github.com/lanvard/syslog v0.0.0-20201116213126-66df7a6162c3 => ../syslog
