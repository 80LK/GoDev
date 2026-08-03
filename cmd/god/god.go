package main

import (
	"os"

	"github.com/80LK/godev/internal/commands"
)

func main() {
	if err := commands.Root.Execute(); err != nil {
		os.Exit(1)
	}
}
