package project

type BuildInfo struct {
	Input      string `modlike:"input" json:"input"`
	Output     string `modlike:"output" json:"output"`
	PreScript  string `modlike:"pre-script,omitempty" json:"pre-script,omitempty"`
	PostScript string `modlike:"post-script,omitempty" json:"post-script,omitempty"`
	OS         string `modlike:"os,omitempty" json:"os,omitempty"`
	Arch       string `modlike:"arch,omitempty" json:"arch,omitempty"`

	Trimpath bool `modlike:"trimpath,omitempty" json:"trimpath,omitempty"`
	Race     bool `modlike:"race,omitempty" json:"race,omitempty"`

	Tags    []string `modlike:"tags" json:"tags,omitempty"`
	GcFlags []string `modlike:"gcflags" json:"gcflags,omitempty"`
	LdFlags []string `modlike:"ldflags" json:"ldflags,omitempty"`
}
