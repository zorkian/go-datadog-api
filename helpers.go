/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2016 by authors and contributors.
 */

package datadog

import "encoding/json"

// Bool is a helper routine that allocates a new bool value
// to store v and returns a pointer to it.
func Bool(v bool) *bool { return &v }

// GetBool is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetBool(v *bool) (b bool, ok bool) {
	if v != nil {
		return *v, true
	}

	return b, ok
}

// Int is a helper routine that allocates a new int value
// to store v and returns a pointer to it.
func Int(v int) *int { return &v }

// GetInt is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetInt(v *int) (i int, ok bool) {
	if v != nil {
		return *v, true
	}

	return i, ok
}

// String is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func String(v string) *string { return &v }

// GetString is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetString(v *string) (str string, ok bool) {
	if v != nil {
		return *v, true
	}

	return str, ok
}

// JsonNumber is a helper routine that allocates a new string value
// to store v and returns a pointer to it.
func JsonNumber(v json.Number) *json.Number { return &v }

// GetJsonNumber is a helper routine that returns a boolean representing
// if a value was set, and if so, dereferences the pointer to it.
func GetJsonNumber(v *json.Number) (num json.Number, ok bool) {
	if v != nil {
		return *v, true
	}

	return num, ok
}
