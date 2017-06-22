/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2016 by authors and contributors.
 */

package datadog

import (
	"fmt"
	"net/url"
	"time"
)

// Snapshot creates an image from a graph and returns the URL of the image.
func (client *Client) Snapshot(options map[string]string, start, end time.Time) (string, error) {
	v := url.Values{}
	v.Add("start", fmt.Sprintf("%d", start.Unix()))
	v.Add("end", fmt.Sprintf("%d", end.Unix()))

	for opt, val := range options {
		v.Add(opt, val)
	}

	out := struct {
		SnapshotURL string `json:"snapshot_url,omitempty"`
	}{}
	if err := client.doJsonRequest("GET", "/v1/graph/snapshot?"+v.Encode(), nil, &out); err != nil {
		return "", err
	}
	return out.SnapshotURL, nil
}
