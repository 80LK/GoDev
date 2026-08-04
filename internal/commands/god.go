package commands

import (
	"errors"

	"github.com/80LK/godev/internal/meta"
	"github.com/80LK/godev/internal/utils"
	"github.com/80LK/godev/internal/utils/logger"
	"github.com/spf13/cobra"
)

var dryRun bool

var Root = &cobra.Command{
	Use:     "god",
	Short:   "GoDev (god) - Devtool for golang projects",
	Long:    "GoDev (god) - Devtool for golang projects",
	Version: meta.Get().Version,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		if !utils.CommandExists("go") {
			return errors.New("go is not installed or not in PATH")
		}

		if !utils.CommandExists("git") {
			return errors.New("git is not installed or not in PATH")
		}

		return nil
	},
}

func init() {
	flags := Root.PersistentFlags()
	flags.BoolVarP(&dryRun, "dry-run", "n", false, "run command without making any changes")
	flags.BoolVarP(&logger.Enabled, "logger", "l", false, "enable logger")
}
