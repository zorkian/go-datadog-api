/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2019 by authors and contributors.
 */

package datadog

import (
	"fmt"
)

const (
	logsPipelinesPath = "/v1/logs/config/pipelines"
)

type LogsPipeline struct {
	Id         *string              `json:"id,omitempty"`
	Type       *string              `json:"type,omitempty"`
	Name       *string              `json:"name"`
	IsEnabled  *bool                `json:"is_enabled,omitempty"`
	IsReadOnly *bool                `json:"is_read_only,omitempty"`
	Filter     *FilterConfiguration `json:"filter"`
	Processors []LogsProcessor      `json:"processors,omitempty"`
}

type FilterConfiguration struct {
	Query *string `json:"query"`
}

func (client *Client) GetLogsPipeline(id string) (*LogsPipeline, error) {
	var pipeline LogsPipeline
	if err := client.doJsonRequest("GET", fmt.Sprintf(logsPipelinesPath+"/%s", id), nil, &pipeline); err != nil {
		return nil, err
	}
	return &pipeline, nil
}

func (client *Client) CreateLogsPipeline(pipeline *LogsPipeline) (*LogsPipeline, error) {
	var createdPipeline = &LogsPipeline{}
	if err := client.doJsonRequest("POST", logsPipelinesPath, pipeline, createdPipeline); err != nil {
		return nil, err
	}
	return createdPipeline, nil
}

func (client *Client) UpdateLogsPipeline(id string, pipeline *LogsPipeline) (*LogsPipeline, error) {
	var updatedPipeline = &LogsPipeline{}
	if err := client.doJsonRequest("PUT", fmt.Sprintf(logsPipelinesPath+"/%s", id), pipeline, updatedPipeline); err != nil {
		return nil, err
	}
	return updatedPipeline, nil
}

// DeleteLogsPipeline returns 200 OK when operation succeed
func (client *Client) DeleteLogsPipeline(id string) error {
	return client.doJsonRequest("DELETE", fmt.Sprintf(logsPipelinesPath+"/%s", id), nil, nil)
}
