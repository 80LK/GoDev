package project

type BuildInfo struct {
	Input      string `modlike:"input"`
	Output     string `modlike:"output"`
	PreScript  string `modlike:"pre-script"`
	PostScript string `modlike:"post-script"`

	Trimpath bool `modlike:"trimpath"`
	Race     bool `modlike:"race"`

	Tags    []string `modlike:"tags"`
	GcFlags []string `modlike:"gcflags"`
	LdFlags []string `modlike:"ldflags"`
}
