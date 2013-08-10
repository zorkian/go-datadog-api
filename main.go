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

// GetDashboard returns a single dashboard created on this account.
func (self *Client) GetDashboard(id int) (*Dashboard, error) {
	var out ReqGetDashboard
	err := self.doJsonRequest("GET", fmt.Sprintf("/v1/dash/%d", id), &out)
	if err != nil {
		return nil, err
	}
	return &out.Dashboard, nil
}

// GetDashboards returns a list of all dashboards created on this account.
func (self *Client) GetDashboards() ([]DashboardLite, error) {
	var out ReqGetDashboards
	err := self.doJsonRequest("GET", "/v1/dash", &out)
	if err != nil {
		return nil, err
	}
	return out.Dashboards, nil
}

// DeleteDashboard deletes a dashboard by the identifier.
func (self *Client) DeleteDashboard(id int) error {
	_, err := self.doSimpleRequest("DELETE", fmt.Sprintf("/v1/dash/%d", id))
	if err != nil {
		return err
	}
	return nil
}
