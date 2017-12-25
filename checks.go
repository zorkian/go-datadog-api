package datadog

// Check represents the request object used in the service check endpoint
type Check struct {
	Check     *string  `json:"check,omitempty"`
	HostName  *string  `json:"host_name,omitempty"`
	Status    *Status  `json:"status,omitempty"`
	Timestamp *string  `json:"timestamp,omitempty"`
	Message   *string  `json:"message,omitempty"`
	Tags      []string `json:"tags,omitempty"`
}

// Status refers to the status of the check run.
type Status int

// Status codes for the service check run
const (
	OK Status = iota
	WARNING
	CRITICAL
	UNKNOWN
)

// PostCheck posts the result of a check run to the server
func (client *Client) PostCheck(check Check) error {
	return client.doJSONRequest("POST", "/v1/check_run",
		check, nil)
}
