/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

// GraphDefinitionRequests are the actual elements in a graph definition.
type GraphDefinitionRequests struct {
	Query   string `json:"q"`
	Stacked bool
}

// GraphDefinition encapsulates configuration information for a graph.
type GraphDefinition struct {
	Viz      string
	Requests []GraphDefinitionRequests
}

// GraphEvents are events that show on a graph.
type GraphEvents struct {
}

// Graph represents a graph that might exist on a dashboard.
type Graph struct {
	Title      string
	Events     []GraphEvents
	Definition GraphDefinition
}

// Dashboard represents a user created dashboard. This is the full dashboard
// struct when we load a dashboard in detail.
type Dashboard struct {
	Id          int
	Description string
	Title       string
	Graphs      []Graph
}

// DashboardLite represents a user created dashboard. This is the mini
// struct when we load the summaries.
type DashboardLite struct {
	Id          int `json:",string"` // FIXME: Remove if DD fixes the API.
	Resource    string
	Description string
	Title       string
}

// ReqGetDashboards from /api/v1/dash
type ReqGetDashboards struct {
	Dashboards []DashboardLite `json:"dashes"`
}

// ReqGetDashboard from /api/v1/dash/:dashboard_id
type ReqGetDashboard struct {
	Resource  string
	Url       string
	Dashboard Dashboard `json:"dash"`
}
