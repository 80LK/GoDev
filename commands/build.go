package commands

import (
	"fmt"

	"github.com/80LK/godev/internal/pipeline"
	goAct "github.com/80LK/godev/internal/pipeline/actions/go"
	projectAct "github.com/80LK/godev/internal/pipeline/actions/project"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/spf13/cobra"
)

var parallel bool
var BuildCmd = &cobra.Command{
	Use:   "build <target>",
	Short: "build target",
	Long:  `Build target sources. If target not set, build all`,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.New(dryRun)
		target := ""
		l := len(args)
		switch l {
		case 0:
		case 1:
			target = args[0]
		default:
			return fmt.Errorf("More args")
		}

		return pipeline.New().Add(
			projectAct.InitProjectContext{},
			goAct.Builds{
				Target:   target,
				Parallel: parallel,
			},
		).Execute(ctx)
	},
}

func init() {
	BuildCmd.Flags().BoolVarP(&parallel, "parallel", "p", false, "paralleling builds")
	Root.AddCommand(BuildCmd)
}
