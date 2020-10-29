package datadog

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

type HostActionResp struct {
	Action   string `json:"action"`
	Hostname string `json:"hostname"`
	Message  string `json:"message,omitempty"`
}

type HostActionMute struct {
	Message  *string `json:"message,omitempty"`
	EndTime  *string `json:"end,omitempty"`
	Override *bool   `json:"override,omitempty"`
}

// MuteHost mutes all monitors for the given host
func (client *Client) MuteHost(host string, action *HostActionMute) (*HostActionResp, error) {
	var out HostActionResp
	uri := "/v1/host/" + host + "/mute"
	if err := client.doJsonRequest("POST", uri, action, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UnmuteHost unmutes all monitors for the given host
func (client *Client) UnmuteHost(host string) (*HostActionResp, error) {
	var out HostActionResp
	uri := "/v1/host/" + host + "/unmute"
	if err := client.doJsonRequest("POST", uri, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// HostTotalsResp defines response to GET /v1/hosts/totals.
type HostTotalsResp struct {
	TotalUp     *int `json:"total_up"`
	TotalActive *int `json:"total_active"`
}

// GetHostTotals returns number of total active hosts and total up hosts.
// Active means the host has reported in the past hour, and up means it has reported in the past two hours.
func (client *Client) GetHostTotals() (*HostTotalsResp, error) {
	var out HostTotalsResp
	uri := "/v1/hosts/totals"
	if err := client.doJsonRequest("GET", uri, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// HostSearchResp defines response to GET /v1/hosts.
type HostSearchResp struct {
	ExactTotalMatching bool   `json:"exact_total_matching"`
	TotalMatching      int    `json:"total_matching"`
	TotalReturned      int    `json:"total_returned"`
	HostList           []Host `json:"host_list"`
}

type Host struct {
	LastReportedTime int64                  `json:"last_reported_time"`
	Name             string                 `json:"name"`
	IsMuted          bool                   `json:"is_muted"`
	MuteTimeout      int                    `json:"mute_timeout"`
	Apps             []string               `json:"apps"`
	TagsBySource     map[string][]string    `json:"tags_by_source"`
	Up               bool                   `json:"up"`
	Metrics          map[string]float64     `json:"metrics"`
	Sources          []string               `json:"sources"`
	Meta             map[string]interface{} `json:"meta"`
	HostName         string                 `json:"host_name"`
	AwsID            *string                `json:"aws_id"`
	ID               int64                  `json:"id"`
	Aliases          []string               `json:"aliases"`
}

type HostSearchRequest struct {
	Filter        string
	SortField     string
	SortDirection string
	Start         int
	Count         int
	FromTs        time.Time
}

// GetHosts searches through the hosts facet, returning matching hosts.
func (client *Client) GetHosts(req HostSearchRequest) (*HostSearchResp, error) {
	v := url.Values{}

	if req.Filter != "" {
		v.Add("filter", req.Filter)
	}
	if req.SortField != "" {
		v.Add("sort_field", req.SortField)
	}
	if req.SortDirection != "" {
		v.Add("sort_dir", req.SortDirection)
	}
	if req.Start >= 0 {
		v.Add("start", strconv.Itoa(req.Start))
	}
	if req.Count >= 0 {
		v.Add("count", strconv.Itoa(req.Count))
	}
	if !req.FromTs.IsZero() {
		v.Add("from", fmt.Sprintf("%d", req.FromTs.Unix()))
	}

	var out HostSearchResp
	uri := "/v1/hosts?" + v.Encode()
	if err := client.doJsonRequest(http.MethodGet, uri, nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}
