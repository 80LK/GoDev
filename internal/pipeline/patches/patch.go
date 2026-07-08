package patches

type Patch interface {
	Apply() error
	Diff() (string, error)
	Summary() string
}
