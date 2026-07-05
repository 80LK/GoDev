package version

import (
	"fmt"
	"strconv"
	"strings"
)

func stripVersionPrefix(s string) string {
	s = strings.TrimSpace(s)
	if v, ok := strings.CutPrefix(s, "v"); ok {
		return v
	}
	if v, ok := strings.CutPrefix(s, "go"); ok {
		return v
	}
	return s
}

func parseUint(s string) (uint, error) {
	v, err := strconv.ParseUint(s, 10, 0)
	return uint(v), err
}

func Parse(s string) (*Version, error) {
	s = stripVersionPrefix(s)

	var v *Version = &Version{}

	main, meta, hasMeta := strings.Cut(s, "+")
	if hasMeta {
		v.Meta = meta
	}

	core, pre, hasPre := strings.Cut(main, "-")
	if hasPre {
		v.PreRelease = pre
	}

	parts := strings.Split(core, ".")
	lParts := len(parts)
	if lParts > 3 || lParts < 1 {
		return v, fmt.Errorf("invalid version core: %q", core)
	}

	var err error
	if v.Major, err = parseUint(parts[0]); err != nil {
		return v, err
	}
	if lParts > 1 {
		if v.Minor, err = parseUint(parts[1]); err != nil {
			return v, err
		}
		if lParts > 2 {
			if v.Patch, err = parseUint(parts[2]); err != nil {
				return v, err
			}
		}
	}

	return v, nil
}

func IsValid(s string) bool {
	_, err := Parse(s)
	return err == nil
}
