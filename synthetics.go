/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

type Check struct {
	Status          *string         `json:"status,omitempty"`
	PublicId        *string         `json:"public_id,omitempty"`
	Tags            []*string       `json:"tags,omitempty"`
	Locations       []*string       `json:"locations,omitempty"`
	Notifications   []*Notification `json:"notifications,omitempty"`
	CheckStatus     *string         `json:"check_status,omitempty"`
	Message         *string         `json:"message,omitempty"`
	Id              *int            `json:"id,omitempty"`
	LastTriggeredTs *int            `json:"last_triggered_ts,omitempty"`
	Name            *string         `json:"string,omitempty"`
	MonitorId       *int            `json:"monitor_id,omitempty"`
	Type            *string         `json:"api,omitempty"`
	CreatedAt       *string         `json:"created_at,omitempty"`
	ModifiedAt      *string         `json:"modified_at,omitempty"`
	Config          []*Config       `json:"config,omitempty"`
	Options         *Options        `json:"options,omitempty"`
}

type Notification struct {
	Handle  *string `json:"handle,omitempty"`
	Name    *string `json:"name,omitempty"`
	Service *string `json:"service,omitempty"`
	Icon    *string `json:"icon,omitempty"`
}

type Config struct {
	Request    *Request     `json:"request,omitempty"`
	Assertions []*Assertion `json:"assertions,omitempty"`
}

type Request struct {
	Url     *string `json:"url,omitempty"`
	Method  *string `json:"method,omitempty"`
	Timeout *int    `json:"timeout,omitempty"`
}

type Assertion struct {
	Operator *string `json:"operator,omitempty"`
	Property *string `json:"property,omitempty"`
	Type     *string `json:"type,omitempty"`
	Target   *string `json:"target,omitempty"`
}

type Options struct {
	TickEvery *int `json:"tick_every,omitempty"`
}

// /api/v0/synthetics/checks/search
type reqSearchChecks struct {
	Checks []*Check `json:"screenboards,omitempty"`
}

func (client *Client) SearchChecks(text string) ([]Check, error) {
	var out reqSearch
	if err := client.doJsonRequest("GET", "/v0/synthetics/checks/search?text="+text, nil, &out); err != nil {
		return nil, err
	}
	return out.Checks, nil
}
