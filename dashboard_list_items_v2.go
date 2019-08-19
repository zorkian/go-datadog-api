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

// DashboardListV2Item represents a single dashboard in a dashboard list.
type DashboardListV2Item struct {
	ID   *string `json:"id,omitempty"`
	Type *string `json:"type,omitempty"`
}

type reqDashboardListV2Items struct {
	Dashboards []DashboardListV2Item `json:"dashboards,omitempty"`
}

type reqAddedDashboardListV2Items struct {
	Dashboards []DashboardListV2Item `json:"added_dashboards_to_list,omitempty"`
}

type reqDeletedDashboardListV2Items struct {
	Dashboards []DashboardListV2Item `json:"deleted_dashboards_from_list,omitempty"`
}

// GetDashboardListV2Items fetches the dashboard list's dashboard definitions.
func (client *Client) GetDashboardListV2Items(id int) ([]DashboardListV2Item, error) {
	var out reqDashboardListV2Items
	if err := client.doJsonRequest("GET", fmt.Sprintf("/v2/dashboard/lists/manual/%d/dashboards", id), nil, &out); err != nil {
		return nil, err
	}
	return out.Dashboards, nil
}

// AddDashboardListV2Items adds dashboards to an existing dashboard list.
//
// Any items already in the list are ignored (not added twice).
func (client *Client) AddDashboardListV2Items(dashboardListID int, items []DashboardListV2Item) ([]DashboardListV2Item, error) {
	req := reqDashboardListV2Items{items}
	var out reqAddedDashboardListV2Items
	if err := client.doJsonRequest("POST", fmt.Sprintf("/v2/dashboard/lists/manual/%d/dashboards", dashboardListID), req, &out); err != nil {
		return nil, err
	}
	return out.Dashboards, nil
}

// UpdateDashboardListV2Items updates dashboards of an existing dashboard list.
//
// This will set the list of dashboards to contain only the items in items.
func (client *Client) UpdateDashboardListV2Items(dashboardListID int, items []DashboardListV2Item) ([]DashboardListV2Item, error) {
	req := reqDashboardListV2Items{items}
	var out reqDashboardListV2Items
	if err := client.doJsonRequest("PUT", fmt.Sprintf("/v2/dashboard/lists/manual/%d/dashboards", dashboardListID), req, &out); err != nil {
		return nil, err
	}
	return out.Dashboards, nil
}

// DeleteDashboardListV2Items deletes dashboards from an existing dashboard list.
//
// Deletes any dashboards in the list of items from the dashboard list.
func (client *Client) DeleteDashboardListV2Items(dashboardListID int, items []DashboardListV2Item) ([]DashboardListV2Item, error) {
	req := reqDashboardListV2Items{items}
	var out reqDeletedDashboardListV2Items
	if err := client.doJsonRequest("DELETE", fmt.Sprintf("/v2/dashboard/lists/manual/%d/dashboards", dashboardListID), req, &out); err != nil {
		return nil, err
	}
	return out.Dashboards, nil
}
