package main

import (
	"fmt"
	"os"

	"github.com/80LK/godev/commands"
)

func main() {
	if err := commands.Root.Execute(); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}
