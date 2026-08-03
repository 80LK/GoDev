package cli

import (
	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/project"

	fsAct "github.com/80LK/godev/internal/pipeline/actions/fs"
	gitAct "github.com/80LK/godev/internal/pipeline/actions/git"
	goAct "github.com/80LK/godev/internal/pipeline/actions/go"
	projectAct "github.com/80LK/godev/internal/pipeline/actions/project"
	"github.com/80LK/godev/internal/pipeline/context"
)

type IntegrateOptions struct {
	DryRunOptions

	Author  string
	Version string
}

func Integrate(opts IntegrateOptions) error {
	ctx := context.New(opts.DryRun)

	pl := pipeline.New().Add(
		fsAct.CheckNotExists{Path: project.GetGoProjectFile(ctx.ProjectDir)},
		fsAct.CheckExistsFile{Path: project.GetGoModFile(ctx.ProjectDir)},
		gitAct.CheckClearGit{},

		projectAct.InitProjectContext{},

		projectAct.IntMod{},
		projectAct.IntProject{
			Author: opts.Author,
		},
	)

	if opts.Version != "" {
		pl.Add(projectAct.VersionSet{Value: opts.Version})
	} else {
		pl.Add(projectAct.VersionLoadFromGit{})
	}

	return pl.Add(
		projectAct.EncodeGoMod{},
		projectAct.EncodeGoProject{},

		gitAct.GitInit{},
		goAct.Tidy{},
		gitAct.GitCommit{Value: "integrate tool god"},
	).Execute(ctx)
}
