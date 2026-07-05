package version

import (
	"cmp"
	"fmt"
	"strconv"
	"strings"

	"github.com/80lk/modlike"
)

type Version struct {
	Major, Minor, Patch uint
	PreRelease          []string
	Meta                string
}

func (v *Version) SetPreRelease(pre string) {
	v.PreRelease = strings.Split(pre, ".")
}

func (v *Version) GetPreRelease() string {
	return strings.Join(v.PreRelease, ".")
}

func (v *Version) String() string {
	var str strings.Builder

	str.WriteRune('v')
	str.WriteString(
		v.StringWithoutSuffix(),
	)

	return str.String()
}

func (v *Version) StringWithoutSuffix() string {
	var str strings.Builder
	str.WriteString(strconv.FormatUint(uint64(v.Major), 10))
	str.WriteRune('.')
	str.WriteString(strconv.FormatUint(uint64(v.Minor), 10))
	str.WriteRune('.')
	str.WriteString(strconv.FormatUint(uint64(v.Patch), 10))

	for i, s := range v.PreRelease {
		if i == 0 {
			str.WriteRune('-')
		} else {
			str.WriteRune('.')
		}
		str.WriteString(s)
	}

	if v.Meta != "" {
		str.WriteRune('+')
		str.WriteString(v.Meta)
	}
	return str.String()
}

func (v *Version) Compare(v2 *Version) int {
	compared := cmp.Compare(v.Major, v2.Major)
	if compared != 0 {
		return compared
	}

	compared = cmp.Compare(v.Minor, v2.Minor)
	if compared != 0 {
		return compared
	}

	compared = cmp.Compare(v.Patch, v2.Patch)
	if compared != 0 {
		return compared
	}

	return comparePreRelease(v.PreRelease, v2.PreRelease)
}

func (v *Version) Equal(other *Version) bool {
	return v.Compare(other) == 0
}

func (v *Version) Less(other *Version) bool {
	return v.Compare(other) < 0
}

func (v *Version) Greater(other *Version) bool {
	return v.Compare(other) > 0
}

func (v *Version) EncodeFrom(raw string) error {
	_v, err := Parse(raw)
	if err != nil {
		return err
	}

	v.Major = _v.Major
	v.Minor = _v.Minor
	v.Patch = _v.Patch
	v.Meta = _v.Meta
	v.PreRelease = _v.PreRelease
	return nil
}

func (v *Version) DecodeModlike(val modlike.Value) error {
	kind := val.Kind()
	if kind == modlike.K_MAP {
		return fmt.Errorf("[Version:DecodeModlike]: invalid kind value %s. Available %s and %s", kind, modlike.K_LIST, modlike.K_STRING)
	}

	var raw string
	if kind == modlike.K_LIST {
		l, _ := val.ToList()
		raw, _ = l.GetStringFirst()
	} else {
		str, _ := val.ToString()
		raw = str.Get()
	}

	return v.EncodeFrom(raw)
}

func (v *Version) EncodeModlike(val modlike.Value) error {
	list, err := val.ToList()
	if err != nil {
		return err
	}

	return list.SetStringStart(v.StringWithoutSuffix())
}

func comparePreRelease(a, b []string) int {
	al := len(a)
	bl := len(b)

	if al == 0 && bl == 0 {
		return 0
	}

	if al == 0 {
		return 1
	}
	if bl == 0 {
		return -1
	}

	l := min(len(a), len(b))
	for i := range l {
		av := a[i]
		avi, err := strconv.Atoi(av)
		avIsStr := err != nil

		bv := b[i]
		bvi, err := strconv.Atoi(bv)
		bvIsStr := err != nil

		if avIsStr {
			if bvIsStr {
				compared := cmp.Compare(av, bv)
				if compared == 0 {
					continue
				}
				return compared
			}

			return -1
		}
		if bvIsStr {
			return 1
		}

		compared := cmp.Compare(avi, bvi)
		if compared == 0 {
			continue
		}
		return compared

	}

	return 0
}
