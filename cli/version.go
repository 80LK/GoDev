package cli

import (
	"github.com/80LK/godev/internal/pipeline"
	"github.com/80LK/godev/internal/pipeline/context"
	"github.com/80LK/godev/project"

	fsAct "github.com/80LK/godev/internal/pipeline/actions/fs"
	gitAct "github.com/80LK/godev/internal/pipeline/actions/git"
	projectAct "github.com/80LK/godev/internal/pipeline/actions/project"
)

type baseVersionOptions struct {
	DryRunOptions

	Release  bool
	NoCommit bool
}

type VersionBumpOptions struct {
	VersionPreReleaseOptions

	Bump projectAct.Bump
}

type VersionPreReleaseOptions struct {
	baseVersionOptions

	Pre string
}

type VersionSetOptions struct {
	baseVersionOptions

	Version string
}

func VersionBump(opts VersionBumpOptions) error {
	ctx := context.New(opts.DryRun)
	pl := pipeline.New().Add(
		fsAct.CheckExistsFile{Path: project.GetGoProjectFile(ctx.ProjectDir)},
	)
	if !opts.NoCommit {
		pl.Add(gitAct.CheckClearGit{})
	}
	pl.Add(
		projectAct.InitProjectContext{},
		projectAct.RunScript{IgnoreNotFound: true, Name: project.LifecycleName(project.PhaseBefore, "version")},
		projectAct.VersionBump{OldVersionKey: "old_version", NewVersionKey: "new_version", Value: opts.Bump},
	)
	if opts.Pre != "" {
		pl.Add(projectAct.VersionSetPre{NewVersionKey: "new_version", Value: opts.Pre})
	}

	pl.Add(
		projectAct.EncodeGoProject{},
		projectAct.PatchSources{OldVersionKey: "old_version"},
		projectAct.GenerateMeta{},
	)

	if !opts.NoCommit {
		pl.Add(gitAct.GitCommit{InputKey: "new_version"})
	}

	if !opts.NoCommit && opts.Release {
		pl.Add(gitAct.GitTagVersion{})
	}
	pl.Add(projectAct.RunScript{IgnoreNotFound: true, Name: project.LifecycleName(project.PhaseAfter, "version")})

	return pl.Execute(ctx)
}
func VersionPreRelease(opts VersionPreReleaseOptions) error {
	ctx := context.New(opts.DryRun)

	pl := pipeline.New().Add(
		fsAct.CheckExistsFile{Path: project.GetGoProjectFile(ctx.ProjectDir)},
	)

	if !opts.NoCommit {
		pl.Add(gitAct.CheckClearGit{})
	}

	pl.Add(
		projectAct.InitProjectContext{},
		projectAct.VersionSetPre{OldVersionKey: "old_version", NewVersionKey: "new_version", Value: opts.Pre},
		projectAct.EncodeGoProject{},
		projectAct.GenerateMeta{},
	)
	if !opts.NoCommit {
		pl.Add(gitAct.GitCommit{InputKey: "new_version"})
	}

	if !opts.NoCommit && opts.Release {
		pl.Add(gitAct.GitTagVersion{})
	}

	return pl.Execute(ctx)
}
func VersionSet(opts VersionSetOptions) error {
	ctx := context.New(opts.DryRun)

	pl := pipeline.New().Add(
		fsAct.CheckExistsFile{Path: project.GetGoProjectFile(ctx.ProjectDir)},
	)

	if !opts.NoCommit {
		pl.Add(gitAct.CheckClearGit{})
	}

	pl.Add(
		projectAct.InitProjectContext{},
		projectAct.VersionSetPre{OldVersionKey: "old_version", NewVersionKey: "new_version", Value: opts.Version},
		projectAct.EncodeGoProject{},
		projectAct.GenerateMeta{},
	)
	if !opts.NoCommit {
		pl.Add(gitAct.GitCommit{InputKey: "new_version"})
	}

	if !opts.NoCommit && opts.Release {
		pl.Add(gitAct.GitTagVersion{})
	}

	return pl.Execute(ctx)
}
