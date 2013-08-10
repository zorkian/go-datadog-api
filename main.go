/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

// Client is the object that handles talking to the Datadog API. This maintains
// state information for a particular application connection.
type Client struct {
	apiKey, appKey string
}

// NewClient returns a new datadog.Client which can be used to access the API
// methods. The expected argument is the API key.
func NewClient(apiKey, appKey string) *Client {
	return &Client{
		apiKey: apiKey,
		appKey: appKey,
	}
}
