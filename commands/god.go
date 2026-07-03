package commands

import (
	"github.com/spf13/cobra"
)

var dryRun bool

var Root = &cobra.Command{
	Use:   "god",
	Short: "short desc",
	Long:  "long desc",
}

func init() {
	Root.PersistentFlags().BoolVarP(&dryRun, "dry-run", "n", false, "Run command without making any changes")
}
