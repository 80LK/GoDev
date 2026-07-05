module github.com/80LK/godev

go 1.26.1

require (
	github.com/80lk/modlike v0.1.0
	github.com/pmezard/go-difflib v1.0.0
	github.com/spf13/cobra v1.10.2
	golang.org/x/mod v0.37.0
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.9 // indirect
)

tool github.com/80LK/godev/cmd/god

replace github.com/80LK/godev/cmd/god v0.0.0 => ./
