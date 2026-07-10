package project

type BuildInfo struct {
	Input  string `modlike:"input"`
	Output string `modlike:"output"`

	Trimpath bool `modlike:"trimpath"`
	Race     bool `modlike:"race"`

	Tags    []string `modlike:"tags"`
	GcFlags []string `modlike:"gcflags"`
	LdFlags []string `modlike:"ldflags"`
}
