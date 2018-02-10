package datadog

// HostActionResp represents the response object for the Hosts API endpoints.
type HostActionResp struct {
	Action   string `json:"action"`
	Hostname string `json:"hostname"`
	Message  string `json:"message,omitempty"`
}

// HostActionMute represents the request object used for the Mute Host API endpoint .
type HostActionMute struct {
	Message  *string `json:"message,omitempty"`
	EndTime  *string `json:"end,omitempty"`
	Override *bool   `json:"override,omitempty"`
}

// MuteHost mutes all monitors for the given host
func (client *Client) MuteHost(host string, action *HostActionMute) (*HostActionResp, error) {
	var out HostActionResp
	uri := "/v1/host/" + host + "/mute"
	if err := client.doJSONRequest("POST", uri, action, &out); err != nil {
		return nil, err
	}
	return &out, nil
}

// UnmuteHost unmutes all monitors for the given host
func (client *Client) UnmuteHost(host string) (*HostActionResp, error) {
	var out HostActionResp
	uri := "/v1/host/" + host + "/unmute"
	if err := client.doJSONRequest("POST", uri, nil, &out); err != nil {
		return nil, err
	}
	return &out, nil
}
