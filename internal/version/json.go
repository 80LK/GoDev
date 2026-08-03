package version

import "encoding/json"

func (v *Version) UnmarshalJSON(data []byte) error {
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		return err
	}
	return v.EncodeFrom(raw)
}

func (v *Version) MarshalJSON() ([]byte, error) {
	return []byte(v.StringWithoutSuffix()), nil
}
