// Package jsonutil provides a collection of types implementing the json.Unmarshaler and json.Marshaler interface.
package jsonutil

import (
	"encoding/json"
	"time"
)

type Duration struct{ time.Duration }

func (d Duration) MarshalJSON() ([]byte, error) {
	return []byte(`"` + d.String() + `"`), nil
}

func (d *Duration) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	tmp, err := time.ParseDuration(s)
	if err != nil {
		return err
	}

	d.Duration = tmp

	return nil
}
