package datadog

import (
	"encoding/json"
	"strconv"
)

// StrIntD can unmarshal both number and string JSON values.
type StrIntD string

// UnmarshalJSON is a Custom Unmarshal for StrIntD. The Datadog API can
// return 1 (int), "1" (number, but a string type) or something like "100%" or
// "*" (string).
func (s *StrIntD) UnmarshalJSON(data []byte) error {
	var num json.Number // json.Number will happily accept any string
	err := json.Unmarshal(data, &num)
	if err == nil {
		*s = StrIntD(num.String())
		return nil
	}

	*s = ""
	return err
}

// StrBoolD can unmarshal both boolean and string JSON values.
type StrBoolD string

// UnmarshalJSON is a Custom Unmarshal for StrBoolD. The Datadog API can
// return true (boolean), "true" (boolean, but as a string type) or entirely
// different strings like "test marker".
func (s *StrBoolD) UnmarshalJSON(data []byte) error {
	var b bool
	err := json.Unmarshal(data, &b)
	if err == nil {
		*s = StrBoolD(strconv.FormatBool(b))
		return nil
	}

	var str string
	err = json.Unmarshal(data, &str)
	if err == nil {
		*s = StrBoolD(str)
		return nil
	}

	*s = ""
	return err
}
