package utils

import "golang.org/x/mod/modfile"

func HasRequire(f *modfile.File, module string) bool {
	for _, req := range f.Require {
		if req.Mod.Path == module {
			return true
		}
	}

	return false
}

func HasTool(f *modfile.File, module string) bool {
	for _, req := range f.Tool {
		if req.Path == module {
			return true
		}
	}

	return false
}
