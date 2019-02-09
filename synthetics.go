/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

type SyntheticsCheck struct {
	Status          *string                       `json:"status,omitempty"`
	PublicId        *string                       `json:"public_id,omitempty"`
	Tags            []string                      `json:"tags,omitempty"`
	Locations       []string                      `json:"locations,omitempty"`
	Notifications   []SyntheticsCheckNotification `json:"notifications,omitempty"`
	CheckStatus     *string                       `json:"check_status,omitempty"`
	Message         *string                       `json:"message,omitempty"`
	Id              *int                          `json:"id,omitempty"`
	LastTriggeredTs *int                          `json:"last_triggered_ts,omitempty"`
	Name            *string                       `json:"string,omitempty"`
	MonitorId       *int                          `json:"monitor_id,omitempty"`
	Type            *string                       `json:"api,omitempty"`
	CreatedAt       *string                       `json:"created_at,omitempty"`
	ModifiedAt      *string                       `json:"modified_at,omitempty"`
	Config          []SyntheticsCheckConfig       `json:"config,omitempty"`
	Options         *SyntheticsCheckOptions       `json:"options,omitempty"`
}

type SyntheticsCheckNotification struct {
	Handle  *string `json:"handle,omitempty"`
	Name    *string `json:"name,omitempty"`
	Service *string `json:"service,omitempty"`
	Icon    *string `json:"icon,omitempty"`
}

type SyntheticsCheckConfig struct {
	Request    *SyntheticsCheckRequest    `json:"request,omitempty"`
	Assertions []SyntheticsCheckAssertion `json:"assertions,omitempty"`
	Locations  *string                    `json:"locations,omitempty"`
}

type SyntheticsCheckRequest struct {
	Url     *string `json:"url,omitempty"`
	Method  *string `json:"method,omitempty"`
	Timeout *int    `json:"timeout,omitempty"`
}

type SyntheticsCheckAssertion struct {
	Operator *string `json:"operator,omitempty"`
	Property *string `json:"property,omitempty"`
	Type     *string `json:"type,omitempty"`
	Target   *string `json:"target,omitempty"`
}

type SyntheticsCheckOptions struct {
	TickEvery *int `json:"tick_every,omitempty"`
}

// /api/v0/synthetics/checks/search
type reqSearchChecks struct {
	Checks []SyntheticsCheck `json:"screenboards,omitempty"`
}

func (client *Client) CreateSyntheticsChecks(text string) ([]SyntheticsCheck, error) {
	var out reqSearchChecks
	if err := client.doJsonRequest("GET", "/v0/synthetics/checks/search?text="+text, nil, &out); err != nil {
		return nil, err
	}
	return out.Checks, nil
}

func (client *Client) CreateSyntheticsCheck(check *SyntheticsCheck) (*SyntheticsCheck, error) {
	var out SyntheticsCheck
	if err := client.doJsonRequest("POST", "/v0/synthetics/checks", check, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
