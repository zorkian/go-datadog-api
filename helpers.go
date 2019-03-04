/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2017 by authors and contributors.
 */

package datadog

import "encoding/json"

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// GetBool is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetBool(v *bool) (bool, bool) {
	if v != nil {
		return *v, true
	}

	return false, false
}

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// Float32 is a helper routine that allocates a new float32 value
// to store v and returns a pointer to it.
func Float32(v float32) *float32 { return &v }

// GetInt is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetIntOk(v *int) (int, bool) {
	if v != nil {
		return *v, true
	}

	return 0, false
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// GetString is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetStringOk(v *string) (string, bool) {
	if v != nil {
		return *v, true
	}

	return "", false
}

// JsonNumber is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func JsonNumber(v json.Number) *json.Number { return &v }

// GetJsonNumber is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetJsonNumberOk(v *json.Number) (json.Number, bool) {
	if v != nil {
		return *v, true
	}

	return "", false
}

// StrInt is a helper routine that allocates a new StrInt value
// to store v and returns a pointer to it.
func StrInt(v StrIntD) *StrIntD { return &v }

// GetStrInt is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetStrInt(v *StrIntD) (StrIntD, bool) {
	if v != nil {
		return *v, true
	}

	return StrIntD(""), false
}

// StrBool is a helper routine that allocates a new StrBool value
// to store v and returns a pointer to it.
func StrBool(v StrBoolD) *StrBoolD { return &v }

// GetStrBool is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetStrBool(v *StrBoolD) (StrBoolD, bool) {
	if v != nil {
		return *v, true
	}

	return StrBoolD(""), false
}
