package cli

import (
	"path"
	"path/filepath"

	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/project"

	fsAct "github.com/80LK/godev/internal/pipeline/actions/fs"
	gitAct "github.com/80LK/godev/internal/pipeline/actions/git"
	goAct "github.com/80LK/godev/internal/pipeline/actions/go"
	projectAct "github.com/80LK/godev/internal/pipeline/actions/project"
)

type InitializeOptions struct {
	IntegrateOptions

	ModuleName string
	Template   string

	Force bool
}

func Initialize(opts InitializeOptions) error {
	ctx := context.New(opts.DryRun)

	moduleName := opts.ModuleName

	//ctx.ProjectDir now CWD
	if moduleName == "." {
		moduleName = filepath.Base(ctx.ProjectDir)
	} else {
		ctx.ProjectDir = filepath.Join(ctx.ProjectDir, path.Base(moduleName))
	}

	pl := pipeline.New()
	if !opts.Force {
		pl.Add(
			fsAct.EnsureEmptyDir{Path: ctx.ProjectDir, Perm: 0777},
			fsAct.CheckNotExists{Path: project.GetGoProjectFile(ctx.ProjectDir)},
			fsAct.CheckNotExists{Path: project.GetJSONProjectFile(ctx.ProjectDir)},
			fsAct.CheckNotExists{Path: project.GetGoModFile(ctx.ProjectDir)},
		)
	}

	pl.Add(
		projectAct.InitProjectContext{},

		projectAct.WriteInProject{
			ModuleName: moduleName,
			Author:     opts.Author,
		},
	)

	if opts.Version != "" {
		pl.Add(projectAct.VersionSet{Value: opts.Version})
	}

	return pl.Add(
		projectAct.CreateFromTemplate{
			Template: opts.Template,
		},

		projectAct.GenerateMeta{},
		gitAct.GitInit{},
		goAct.Tidy{},
		gitAct.GitCommit{Value: "init commit"},
	).Execute(ctx)
}
