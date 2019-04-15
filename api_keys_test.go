/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2019 by authors and contributors.
 */

package datadog_test

import (
	"testing"
	"time"

	"encoding/json"

	"github.com/stretchr/testify/assert"
	dd "github.com/zorkian/go-datadog-api"
)

func TestAPIKeySerialization(t *testing.T) {

	raw := `
	{
		"created_by": "john@example.com",
		"name": "myCoolKey",
		"key": "3111111111111111aaaaaaaaaaaaaaaa",
		"created": "2019-04-05 09:47:00"
	}`

	var apikey dd.APIKey
	err := json.Unmarshal([]byte(raw), &apikey)
	assert.Equal(t, err, nil)
	assert.Equal(t, *apikey.Name, "myCoolKey")
	assert.Equal(t, *apikey.CreatedBy, "john@example.com")
	assert.Equal(t, *apikey.Key, "3111111111111111aaaaaaaaaaaaaaaa")
	assert.Equal(t, *apikey.Created, time.Date(2019, 4, 5, 9, 47, 0, 0, time.UTC))

	// make sure that the date is correct after marshaling
	res, _ := json.Marshal(apikey)
	assert.Contains(t, string(res), "\"2019-04-05 09:47:00\"")
}
