package project

type BuildInfo struct {
	Input      string `modlike:"input"`
	Output     string `modlike:"output"`
	PreScript  string `modlike:"pre-script,omitempty"`
	PostScript string `modlike:"post-script,omitempty"`
	OS         string `modlike:"os,omitempty"`
	Arch       string `modlike:"arch,omitempty"`

	Trimpath bool `modlike:"trimpath,omitempty"`
	Race     bool `modlike:"race,omitempty"`

	Tags    []string `modlike:"tags"`
	GcFlags []string `modlike:"gcflags"`
	LdFlags []string `modlike:"ldflags"`
}
