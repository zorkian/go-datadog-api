/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

import (
	"fmt"
)

// GraphDefinitionRequests are the actual elements in a graph definition.
type GraphDefinitionRequests struct {
	Query   string `json:"q"`
	Stacked bool   `json:"stacked"`
}

// GraphDefinition encapsulates configuration information for a graph.
type GraphDefinition struct {
	Viz      string                    `json:"viz"`
	Requests []GraphDefinitionRequests `json:"requests"`
}

// GraphEvents are events that show on a graph.
type GraphEvents struct {
}

// Graph represents a graph that might exist on a dashboard.
type Graph struct {
	Title      string          `json:"title"`
	Events     []GraphEvents   `json:"events"`
	Definition GraphDefinition `json:"definition"`
}

// Dashboard represents a user created dashboard. This is the full dashboard
// struct when we load a dashboard in detail.
type Dashboard struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	Title       string  `json:"title"`
	Graphs      []Graph `json:"graphs"`
}

// DashboardLite represents a user created dashboard. This is the mini
// struct when we load the summaries.
type DashboardLite struct {
	Id          int    `json:"id,string"` // TODO: Remove ',string'.
	Resource    string `json:"resource"`
	Description string `json:"description"`
	Title       string `json:"title"`
}

// ReqGetDashboards from /api/v1/dash
type ReqGetDashboards struct {
	Dashboards []DashboardLite `json:"dashes"`
}

// ReqGetDashboard from /api/v1/dash/:dashboard_id
type ReqGetDashboard struct {
	Resource  string    `json:"resource"`
	Url       string    `json:"url"`
	Dashboard Dashboard `json:"dash"`
}

// GetDashboard returns a single dashboard created on this account.
func (self *Client) GetDashboard(id int) (*Dashboard, error) {
	var out ReqGetDashboard
	err := self.doJsonRequest("GET", fmt.Sprintf("/v1/dash/%d", id), nil, &out)
	if err != nil {
		return nil, err
	}
	return &out.Dashboard, nil
}

// GetDashboards returns a list of all dashboards created on this account.
func (self *Client) GetDashboards() ([]DashboardLite, error) {
	var out ReqGetDashboards
	err := self.doJsonRequest("GET", "/v1/dash", nil, &out)
	if err != nil {
		return nil, err
	}
	return out.Dashboards, nil
}

// DeleteDashboard deletes a dashboard by the identifier.
func (self *Client) DeleteDashboard(id int) error {
	return self.doJsonRequest("DELETE", fmt.Sprintf("/v1/dash/%d", id), nil, nil)
}

// CreateDashboard creates a new dashboard when given a Dashboard struct. Note
// that the Id, Resource, Url and similar elements are not used in creation.
func (self *Client) CreateDashboard(dash *Dashboard) (*Dashboard, error) {
	var out ReqGetDashboard
	err := self.doJsonRequest("POST", "/v1/dash", dash, &out)
	if err != nil {
		return nil, err
	}
	return &out.Dashboard, nil
}

// UpdateDashboard in essence takes a Dashboard struct and persists it back to
// the server. Use this if you've updated your local and need to push it back.
func (self *Client) UpdateDashboard(dash *Dashboard) error {
	err := self.doJsonRequest("PUT",
		fmt.Sprintf("/v1/dash/%d", dash.Id), dash, nil)
	if err != nil {
		return err
	}
	return nil
}
