package datadog

import (
	"time"
)

const logsListPath = "/v1/logs-queries/list"

// LogsListRequest represents the request body sent to the list API
type LogsListRequest struct {
	Index   *string    `json:"index,omitempty"`
	Limit   *int       `json:"limit,omitempty"`
	Query   *string    `json:"query"`
	Sort    *string    `json:"sort,omitempty"`
	StartAt *string    `json:"startAt,omitempty"`
	Time    *QueryTime `json:"time"`
}

// QueryTime represents the time object for the request sent to the list API
type QueryTime struct {
	TimeFrom *string `json:"from"`
	TimeTo   *string `json:"to"`
	TimeZone *string `json:"timezone,omitempty"`
}

// LogsList represents the base API response returned by the list API
type LogsList struct {
	Logs      []Logs  `json:"logs"`
	NextLogID *string `json:"nextLogId"`
	Status    *string `json:"status"`
}

func (l *LogsList) next() bool {
	if l.NextLogID != nil || l.Status == nil {
		return true
	}

	return false
}

// Logs represents the data of a log entry and contains the UUID of that entry (which
// is used for the StartAt option in an API request) as well as the content of that log
type Logs struct {
	ID      *string `json:"id"`
	Content Content `json:"content"`
}

// Content respresents the actual log content returned by the list API
type Content struct {
	Timestamp  *time.Time `json:"timestamp"`
	Tags       []string   `json:"tags,omitempty"`
	Attributes Attributes `json:"attributes,omitempty"`
	Host       *string    `json:"host"`
	Service    *string    `json:"service"`
	Message    *string    `json:"message"`
}

// Attributes represents the Content attribute object from the list API
type Attributes map[string]interface{}

// GetLogsList gets a page of log entries based on the values in the provided LogListRequest
func (client *Client) GetLogsList(logsRequest *LogsListRequest) (logsList *LogsList, err error) {
	out := &LogsList{}

	if err = client.doJsonRequest("POST", logsListPath, logsRequest, out); err != nil {
		return nil, err
	}

	return out, nil
}

// GetLogsListPages calls GetLogsList and handles the pagination performed by the 'logs-queries/list' API
func (client *Client) GetLogsListPages(logsRequest *LogsListRequest) (logs []Logs, err error) {

	response, err := client.GetLogsList(logsRequest)
	if err != nil {
		return response.Logs, err
	}

	logs = append(logs, response.Logs...)

	for response.next() && err == nil {
		logsRequest.StartAt = response.NextLogID
		response, err = client.GetLogsList(logsRequest)
		logs = append(logs, response.Logs...)
	}

	return logs, err
}
