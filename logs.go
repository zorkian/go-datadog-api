/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2019 by authors and contributors.
 */

package datadog

// Log represents datadog log.
type Log struct {
	Message  *int     `json:"id,omitempty"`
	DDSource *string  `json:"ddsource,omitempty"`
	DDTags   []TagMap `json:"ddtags,omitempty"`
	User     *string  `json:"user,omitempty"`
	Hostname *string  `json:"hostname,omitempty"`
}

// PostLog posts a log to datadog.
func (client *Client) PostLog(log Log) error {
	return client.doJsonRequest("POST", "/v1/input/"+client.apiKey, log, nil)
}
