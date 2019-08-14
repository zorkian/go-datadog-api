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

func TestIPRangesSerialization(t *testing.T) {

	var mapResponse map[string][]string
	mapResponse = make(map[string][]string)

	mapResponse["prefixes_ipv4"] = []string{"10.10.10.10/32", "10.10.10.10/32"}
	mapResponse["prefixes_ipv6"] = []string{"2000:1900:0:100c::/128", "2000:1900:0:c100::/128"}

	raw := `
	{
		"agents": {
			"prefixes_ipv4": [
				"10.10.10.10/32",
				"10.10.10.10/32"
			],
			"prefixes_ipv6": [
				"2000:1900:0:100c::/128",
				"2000:1900:0:c100::/128"
			]
		}
	}`

	var ipranges dd.IPRangesResp
	err := json.Unmarshal([]byte(raw), &ipranges)
	assert.Equal(t, err, nil)
	assert.Equal(t, ipranges.Agents["prefixes_ipv4"], mapResponse["prefixes_ipv4"])
	assert.Equal(t, ipranges.Agents["prefixes_ipv6"], mapResponse["prefixes_ipv6"])
}
