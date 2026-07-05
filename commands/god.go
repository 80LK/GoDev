package commands

import (
	"github.com/spf13/cobra"
)

var dryRun bool

var Root = &cobra.Command{
	Use:   "god",
	Short: "GoDev (god) - Devtool for golang projects",
	Long:  "GoDev (god) - Devtool for golang projects",
}

func init() {
	Root.PersistentFlags().BoolVarP(&dryRun, "dry-run", "n", false, "run command without making any changes")
}
