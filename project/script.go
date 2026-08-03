package project

type RootScript struct {
	Script

	Commands []*Script `modlike:"commands"`
}

func (r RootScript) AsScript() *Script {
	return &Script{
		Command: r.Command,
		Args:    r.Args,
		WorkDir: r.WorkDir,
		Env:     r.Env,
	}
}

type Script struct {
	Command string   `modlike:"command,omitempty" json:"command,omitempty"`
	Args    []string `modlike:"args" json:"args,omitempty"`
	WorkDir string   `modlike:"workdir,omitempty" json:"workdir,omitempty"`
	Env     []string `modlike:"env" json:"env,omitempty"`
}

type ScriptPhase string

const (
	PhaseBefore ScriptPhase = "pre"
	PhaseAfter  ScriptPhase = "post"
)

func LifecycleName(phase ScriptPhase, target string) string {
	return string(phase) + "-" + target
}
