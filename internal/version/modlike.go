package version

import (
	"fmt"

	"github.com/80LK/modlike"
)

func (v *Version) UnmarshalModlike(val modlike.Value) error {
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

func (v *Version) MarshalModlike(val modlike.Value) error {
	list, err := val.ToList()
	if err != nil {
		return err
	}

	return list.SetStringStart(v.StringWithoutSuffix())
}
