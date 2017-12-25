/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

import (
	"encoding/json"
	"fmt"
)

// GraphDefinitionRequestStyle represents the graph style attributes
type GraphDefinitionRequestStyle struct {
	Palette *string `json:"palette,omitempty"`
	Width   *string `json:"width,omitempty"`
	Type    *string `json:"type,omitempty"`
}

// GraphDefinitionRequest represents the requests passed into each graph.
type GraphDefinitionRequest struct {
	Query              *string                      `json:"q,omitempty"`
	Stacked            *bool                        `json:"stacked,omitempty"`
	Aggregator         *string                      `json:"aggregator,omitempty"`
	ConditionalFormats []DashboardConditionalFormat `json:"conditional_formats,omitempty"`
	Type               *string                      `json:"type,omitempty"`
	Style              *GraphDefinitionRequestStyle `json:"style,omitempty"`

	// For change type graphs
	ChangeType     *string `json:"change_type,omitempty"`
	OrderDirection *string `json:"order_dir,omitempty"`
	CompareTo      *string `json:"compare_to,omitempty"`
	IncreaseGood   *bool   `json:"increase_good,omitempty"`
	OrderBy        *string `json:"order_by,omitempty"`
	ExtraCol       *string `json:"extra_col,omitempty"`
}

// GraphDefinitionMarker represents the "conditional_formats" field of the GraphDefinitionRequest
type GraphDefinitionMarker struct {
	Type  *string      `json:"type,omitempty"`
	Value *string      `json:"value,omitempty"`
	Label *string      `json:"label,omitempty"`
	Val   *json.Number `json:"val,omitempty"`
	Min   *json.Number `json:"min,omitempty"`
	Max   *json.Number `json:"max,omitempty"`
}

// GraphEvent represents the "events" field of GraphDefinition
type GraphEvent struct {
	Query *string `json:"q,omitempty"`
}

// Yaxis represents the "yaxis" field of GraphDefinition
type Yaxis struct {
	Min   *float64 `json:"min,omitempty"`
	Max   *float64 `json:"max,omitempty"`
	Scale *string  `json:"scale,omitempty"`
}

// Style represents the "style" field of GraphDefinition
type Style struct {
	Palette     *string `json:"palette,omitempty"`
	PaletteFlip *bool   `json:"paletteFlip,omitempty"`
}

// GraphDefinition represents a JSON document that defines a graph
type GraphDefinition struct {
	Viz      *string                  `json:"viz,omitempty"`
	Requests []GraphDefinitionRequest `json:"requests,omitempty"`
	Events   []GraphEvent             `json:"events,omitempty"`
	Markers  []GraphDefinitionMarker  `json:"markers,omitempty"`

	// For timeseries type graphs
	Yaxis Yaxis `json:"yaxis,omitempty"`

	// For query value type graphs
	Autoscale  *bool   `json:"autoscale,omitempty"`
	TextAlign  *string `json:"text_align,omitempty"`
	Precision  *string `json:"precision,omitempty"`
	CustomUnit *string `json:"custom_unit,omitempty"`

	// For hostname type graphs
	Style *Style `json:"Style,omitempty"`

	Groups                []string `json:"group,omitempty"`
	IncludeNoMetricHosts  *bool    `json:"noMetricHosts,omitempty"`
	Scopes                []string `json:"scope,omitempty"`
	IncludeUngroupedHosts *bool    `json:"noGroupHosts,omitempty"`
}

// Graph represents a graph that might exist on a dashboard.
type Graph struct {
	Title      *string          `json:"title,omitempty"`
	Definition *GraphDefinition `json:"definition"`
}

// TemplateVariable represents a template variable that might exist on a dashboard
type TemplateVariable struct {
	Name    *string `json:"name,omitempty"`
	Prefix  *string `json:"prefix,omitempty"`
	Default *string `json:"default,omitempty"`
}

// Dashboard represents a user created dashboard. This is the full dashboard
// struct when we load a dashboard in detail.
type Dashboard struct {
	ID                *int               `json:"id,omitempty"`
	Description       *string            `json:"description,omitempty"`
	Title             *string            `json:"title,omitempty"`
	Graphs            []Graph            `json:"graphs,omitempty"`
	TemplateVariables []TemplateVariable `json:"template_variables,omitempty"`
	ReadOnly          *bool              `json:"read_only,omitempty"`
}

// DashboardLite represents a user created dashboard. This is the mini
// struct when we load the summaries.
type DashboardLite struct {
	ID          *int    `json:"id,string,omitempty"` // TODO: Remove ',string'.
	Resource    *string `json:"resource,omitempty"`
	Description *string `json:"description,omitempty"`
	Title       *string `json:"title,omitempty"`
}

// reqGetDashboards from /api/v1/dash
type reqGetDashboards struct {
	Dashboards []DashboardLite `json:"dashes,omitempty"`
}

// reqGetDashboard from /api/v1/dash/:dashboard_id
type reqGetDashboard struct {
	Resource  *string    `json:"resource,omitempty"`
	URL       *string    `json:"url,omitempty"`
	Dashboard *Dashboard `json:"dash,omitempty"`
}

// DashboardConditionalFormat represents the "conditional_formats" field in GraphDefinitionRequest
type DashboardConditionalFormat struct {
	Palette        *string      `json:"palette,omitempty"`
	Comparator     *string      `json:"comparator,omitempty"`
	CustomBgColor  *string      `json:"custom_bg_color,omitempty"`
	Value          *json.Number `json:"value,omitempty"`
	Inverted       *bool        `json:"invert,omitempty"`
	CustomFgColor  *string      `json:"custom_fg_color,omitempty"`
	CustomImageURL *string      `json:"custom_image,omitempty"`
}

// GetDashboard returns a single dashboard created on this account.
func (client *Client) GetDashboard(id int) (*Dashboard, error) {
	var out reqGetDashboard
	if err := client.doJSONRequest("GET", fmt.Sprintf("/v1/dash/%d", id), nil, &out); err != nil {
		return nil, err
	}
	return out.Dashboard, nil
}

// GetDashboards returns a list of all dashboards created on this account.
func (client *Client) GetDashboards() ([]DashboardLite, error) {
	var out reqGetDashboards
	if err := client.doJSONRequest("GET", "/v1/dash", nil, &out); err != nil {
		return nil, err
	}
	return out.Dashboards, nil
}

// DeleteDashboard deletes a dashboard by the identifier.
func (client *Client) DeleteDashboard(id int) error {
	return client.doJSONRequest("DELETE", fmt.Sprintf("/v1/dash/%d", id), nil, nil)
}

// CreateDashboard creates a new dashboard when given a Dashboard struct. Note
// that the ID, Resource, URL and similar elements are not used in creation.
func (client *Client) CreateDashboard(dash *Dashboard) (*Dashboard, error) {
	var out reqGetDashboard
	if err := client.doJSONRequest("POST", "/v1/dash", dash, &out); err != nil {
		return nil, err
	}
	return out.Dashboard, nil
}

// UpdateDashboard in essence takes a Dashboard struct and persists it back to
// the server. Use this if you've updated your local and need to push it back.
func (client *Client) UpdateDashboard(dash *Dashboard) error {
	return client.doJSONRequest("PUT", fmt.Sprintf("/v1/dash/%d", *dash.ID),
		dash, nil)
}
