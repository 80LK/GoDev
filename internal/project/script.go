package project

type RootScript struct {
	*Script

	Commands []*Script `modlike:"commands"`
}

type Script struct {
	Command *string  `modlike:"command"`
	Args    []string `modlike:"args"`
	WorkDir *string  `modlike:"workdir"`
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
