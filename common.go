package datadog

import "encoding/json"

// StrintD ...
type StrintD string

// UnmarshalJSON is a Custom Unmarshal for StrintD. The Datadog API can
// return 1 (int), "1" (number, but a string type) or something like "100%" or
// "*" (string).
func (s *StrintD) UnmarshalJSON(data []byte) error {
	var err error
	var strintNumber json.Number
	if err = json.Unmarshal(data, &strintNumber); err == nil {
		*s = StrintD(strintNumber)
		return nil
	}

	var strintStr StrintD
	if err = json.Unmarshal(data, &strintStr); err == nil {
		*s = StrintD(strintStr)
		return nil
	}

	var s0 StrintD
	*s = s0

	return err
}
