package main

import (
	"os"

	"github.com/80LK/godev/commands"
)

func main() {
	if err := commands.Root.Execute(); err != nil {
		os.Exit(1)
	}
}
