package commands

import (
	"fmt"

	"github.com/80LK/godev/internal/pipeline"
	projectAct "github.com/80LK/godev/internal/pipeline/actions/project"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/spf13/cobra"
)

var RunCmd = &cobra.Command{
	Use:   "run <script>",
	Short: "Run script",
	Long:  "Run script by name <script>",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.New(dryRun)

		if len(args) != 1 {
			return fmt.Errorf("Expect only name script")
		}

		return pipeline.New().Add(
			projectAct.InitProjectContext{},
			projectAct.RunScript{
				Name: args[0],
			},
		).Execute(ctx)
	},
}

func init() {
	Root.AddCommand(RunCmd)
}
