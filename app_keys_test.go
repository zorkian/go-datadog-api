/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2019 by authors and contributors.
 */

package datadog_test

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
	dd "github.com/zorkian/go-datadog-api"
)

func TestAPPKeySerialization(t *testing.T) {

	raw := `
	{
		"owner": "john@example.com",
		"name": "myCoolKey",
		"hash": "31111111111111111111aaaaaaaaaaaaaaaaaaaa"
	}`

	var appkey dd.APPKey
	err := json.Unmarshal([]byte(raw), &appkey)
	assert.Equal(t, err, nil)
	assert.Equal(t, *appkey.Name, "myCoolKey")
	assert.Equal(t, *appkey.Owner, "john@example.com")
	assert.Equal(t, *appkey.Hash, "31111111111111111111aaaaaaaaaaaaaaaaaaaa")
}
