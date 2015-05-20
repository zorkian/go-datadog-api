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

// Screenboard represents a user created screenboard. This is the full screenboard
// struct when we load a screenboard in detail.
type Screenboard struct {
	Id                int                `json:"id,omitempty"`
	Title             string             `json:"board_title"`
	Height            int                `json:"height,omitempty"`
	Width             string             `json:"width,omitempty"`
	Shared            bool               `json:"shared"`
	Templated         bool               `json:"templated,omitempty"`
	TemplateVariables []TemplateVariable `json:"template_variables,omitempty"`
	Widgets           []Widget           `json:"widgets"`
}
type TemplateVariable struct {
	Default string `json:"default"`
	Name    string `json:"name"`
	Prefix  string `json:"prefix"`
}

func (s *Screenboard) UnmarshalJSON(data []byte) error {
	dest := struct {
		Id                int                `json:"id"`
		Title             string             `json:"board_title"`
		Height            int                `json:"height"`
		Width             string             `json:"width"`
		Shared            bool               `json:"shared"`
		Templated         bool               `json:"templated"`
		TemplateVariables []TemplateVariable `json:"template_variables"`
		Widgets           []json.RawMessage  `json:"widgets"`
	}{}
	err := json.Unmarshal(data, &dest)
	if err != nil {
		return err
	}

	widgets := []Widget{}
	for _, rawWidget := range dest.Widgets {
		typeStruct := struct {
			Type string `json:"type"`
		}{}
		if err := json.Unmarshal(rawWidget, &typeStruct); err != nil {
			return fmt.Errorf("Could not detect widget type: %s", err)
		}

		widget, err := unmarshalWidget(typeStruct.Type, rawWidget)
		if err != nil {
			return fmt.Errorf("Could not unmarshal widget (type %s): %s", typeStruct.Type, err)
		}

		widgets = append(widgets, widget)
	}

	s.Id = dest.Id
	s.Title = dest.Title
	s.Height = dest.Height
	s.Width = dest.Width
	s.Shared = dest.Shared
	s.Templated = dest.Templated
	s.TemplateVariables = dest.TemplateVariables
	s.Widgets = widgets

	return nil
}
func unmarshalWidget(widgetType string, data json.RawMessage) (Widget, error) {
	var dest Widget

	switch widgetType {
	case "timeseries":
		dest = &TimeseriesWidget{}
	case "query_value":
		dest = &QueryValueWidget{}
	case "event_stream":
		dest = &EventStreamWidget{}
	case "free_text":
		dest = &FreeTextWidget{}
	case "toplist":
		dest = &ToplistWidget{}
	default:
		return nil, fmt.Errorf("Could not unmarshal unknown widget type %s.", widgetType)
	}

	if err := json.Unmarshal(data, dest); err != nil {
		return nil, err
	}

	return dest, nil
}

// ScreenboardLite represents a user created screenboard. This is the mini
// struct when we load the summaries.
type ScreenboardLite struct {
	Id       int    `json:"id"`
	Resource string `json:"resource"`
	Title    string `json:"title"`
}

// reqGetScreenboards from /api/v1/screen
type reqGetScreenboards struct {
	Screenboards []*ScreenboardLite `json:"screenboards"`
}

// GetScreenboard returns a single screenboard created on this account.
func (self *Client) GetScreenboard(id int) (*Screenboard, error) {
	out := &Screenboard{}
	err := self.doJsonRequest("GET", fmt.Sprintf("/v1/screen/%d", id), nil, out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GetScreenboards returns a list of all screenboards created on this account.
func (self *Client) GetScreenboards() ([]*ScreenboardLite, error) {
	var out reqGetScreenboards
	err := self.doJsonRequest("GET", "/v1/screen", nil, &out)
	if err != nil {
		return nil, err
	}
	return out.Screenboards, nil
}

// DeleteScreenboard deletes a screenboard by the identifier.
func (self *Client) DeleteScreenboard(id int) error {
	return self.doJsonRequest("DELETE", fmt.Sprintf("/v1/screen/%d", id), nil, nil)
}

// CreateScreenboard creates a new screenboard when given a Screenboard struct. Note
// that the Id, Resource, Url and similar elements are not used in creation.
func (self *Client) CreateScreenboard(board *Screenboard) (*Screenboard, error) {
	out := &Screenboard{}
	if err := self.doJsonRequest("POST", "/v1/screen", board, out); err != nil {
		return nil, err
	}
	return out, nil
}

// UpdateScreenboard in essence takes a Screenboard struct and persists it back to
// the server. Use this if you've updated your local and need to push it back.
func (self *Client) UpdateScreenboard(board *Screenboard) error {
	return self.doJsonRequest("PUT", fmt.Sprintf("/v1/screen/%d", board.Id), board, nil)
}
