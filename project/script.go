package project

type RootScript struct {
	Command string   `modlike:"command"`
	Args    []string `modlike:"args"`
	WorkDir string   `modlike:"workdir"`
	Env     []string `modlike:"env"`

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
	Command string   `modlike:"command"`
	Args    []string `modlike:"args"`
	WorkDir string   `modlike:"workdir,omitempty"`
	Env     []string `modlike:"env"`
}

type ScriptPhase string

const (
	PhaseBefore ScriptPhase = "pre"
	PhaseAfter  ScriptPhase = "post"
)

func LifecycleName(phase ScriptPhase, target string) string {
	return string(phase) + "-" + target
}
