package datadog

import "encoding/json"

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
