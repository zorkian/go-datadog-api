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

// WidthS ...
type WidthS string

// UnmarshalJSON is a Custom Unmarshal for WidthS. The Datadog API can
// return 1 (int), "1" (number, but a string type) or something like "100%" (string).
func (w *WidthS) UnmarshalJSON(data []byte) error {
	var err error
	var widthNum json.Number
	if err = json.Unmarshal(data, &widthNum); err == nil {
		*w = WidthS(widthNum)
		return nil
	}
	var widthStr string
	if err = json.Unmarshal(data, &widthStr); err == nil {
		*w = WidthS(widthStr)
		return nil
	}
	var w0 WidthS
	*w = w0
	return err
}

// HeightS ...
type HeightS string

// UnmarshalJSON is a Custom Unmarshal for HeightS. The Datadog API can
// return 1 (int), "1" (number, but a string type) or something like "100%" (string).
func (h *HeightS) UnmarshalJSON(data []byte) error {
	var err error
	var heightNum json.Number
	if err = json.Unmarshal(data, &heightNum); err == nil {
		*h = HeightS(heightNum)
		return nil
	}
	var heightStr string
	if err = json.Unmarshal(data, &heightStr); err == nil {
		*h = HeightS(heightStr)
		return nil
	}
	var h0 HeightS
	*h = h0
	return err
}

// Screenboard represents a user created screenboard. This is the full screenboard
// struct when we load a screenboard in detail.
type Screenboard struct {
	Id                *int               `json:"id,omitempty"`
	Title             *string            `json:"board_title,omitempty"`
	Height            *HeightS           `json:"height,omitempty"`
	Width             *WidthS            `json:"width,omitempty"`
	Shared            *bool              `json:"shared,omitempty"`
	TemplateVariables []TemplateVariable `json:"template_variables,omitempty"`
	Widgets           []Widget           `json:"widgets"`
	ReadOnly          *bool              `json:"read_only,omitempty"`
}

// ScreenboardLite represents a user created screenboard. This is the mini
// struct when we load the summaries.
type ScreenboardLite struct {
	Id       *int    `json:"id,omitempty"`
	Resource *string `json:"resource,omitempty"`
	Title    *string `json:"title,omitempty"`
}

// reqGetScreenboards from /api/v1/screen
type reqGetScreenboards struct {
	Screenboards []*ScreenboardLite `json:"screenboards,omitempty"`
}

// GetScreenboard returns a single screenboard created on this account.
func (client *Client) GetScreenboard(id int) (*Screenboard, error) {
	out := &Screenboard{}
	if err := client.doJsonRequest("GET", fmt.Sprintf("/v1/screen/%d", id), nil, out); err != nil {
		return nil, err
	}
	return out, nil
}

// GetScreenboards returns a list of all screenboards created on this account.
func (client *Client) GetScreenboards() ([]*ScreenboardLite, error) {
	var out reqGetScreenboards
	if err := client.doJsonRequest("GET", "/v1/screen", nil, &out); err != nil {
		return nil, err
	}
	return out.Screenboards, nil
}

// DeleteScreenboard deletes a screenboard by the identifier.
func (client *Client) DeleteScreenboard(id int) error {
	return client.doJsonRequest("DELETE", fmt.Sprintf("/v1/screen/%d", id), nil, nil)
}

// CreateScreenboard creates a new screenboard when given a Screenboard struct. Note
// that the Id, Resource, Url and similar elements are not used in creation.
func (client *Client) CreateScreenboard(board *Screenboard) (*Screenboard, error) {
	out := &Screenboard{}
	if err := client.doJsonRequest("POST", "/v1/screen", board, out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateScreenboard in essence takes a Screenboard struct and persists it back to
// the server. Use this if you've updated your local and need to push it back.
func (client *Client) UpdateScreenboard(board *Screenboard) error {
	return client.doJsonRequest("PUT", fmt.Sprintf("/v1/screen/%d", *board.Id), board, nil)
}

type ScreenShareResponse struct {
	BoardId   int    `json:"board_id"`
	PublicUrl string `json:"public_url"`
}

// ShareScreenboard shares an existing screenboard, it takes and updates ScreenShareResponse
func (client *Client) ShareScreenboard(id int, response *ScreenShareResponse) error {
	return client.doJsonRequest("POST", fmt.Sprintf("/v1/screen/share/%d", id), nil, response)
}

// RevokeScreenboard revokes a currently shared screenboard
func (client *Client) RevokeScreenboard(id int) error {
	return client.doJsonRequest("DELETE", fmt.Sprintf("/v1/screen/share/%d", id), nil, nil)
}
