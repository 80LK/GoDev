package meta

type Meta struct {
	Name    string
	Version string
	Module  string
	Author  string
}

func Get() Meta {
	return meta
}
