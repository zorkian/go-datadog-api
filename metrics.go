/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2015 by authors and contributors.
 */

package datadog

import "fmt"

type Data struct {
	Metric string     `json:"metric"`
	Points [][]string `json:"points"`
	Type   string     `json:"type"`
	Host   string     `json:"host"`
	Tags   []string   `json:"tags"`
}

type Series []Data

type SeriesData struct {
	Series []Data `json:"series"`
}

// Post data to datadog metric
func (self *Client) PostData(seriesData *SeriesData) error {
	err := self.doJsonRequest("POST", fmt.Sprintf("/v1/series"), &seriesData, nil)
	if err != nil {
		return err
	}
	return nil
}
