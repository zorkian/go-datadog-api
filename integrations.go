/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

type servicePD struct {
	ServiceName *string `json:"service"`
	ServiceKey  *string `json:"key"`
}

type integrationPD struct {
	Services  []servicePD `json:"services"`
	Subdomain *string     `json:"subdomain"`
	Schedules []string    `json:"schedules"`
	APIToken  *string     `json:"api_token"`
}

type ServicePDRequest struct {
	ServiceName *string `json:"service_name"`
	ServiceKey  *string `json:"service_key"`
}

type IntegrationPDRequest struct {
	Services  []ServicePDRequest `json:"services,omitempty"`
	Subdomain *string            `json:"subdomain,omitempty"`
	Schedules []string           `json:"schedules,omitempty"`
	APIToken  *string            `json:"api_token,omitempty"`
	RunCheck  *bool              `json:"run_check,omitempty"`
}

// CreateIntegrationPD creates new Pagerduty Integrations.
func (client *Client) CreateIntegrationPD(pdIntegration *IntegrationPDRequest) error {
	return client.doJsonRequest("POST", "/v1/integration/pagerduty", pdIntegration, nil)
}

// UpdateIntegrationPD updates the Pagerduty Integration.
// This will replace the existing values with the new values
func (client *Client) UpdateIntegrationPD(pdIntegration *IntegrationPDRequest) error {
	return client.doJsonRequest("PUT", "/v1/integration/pagerduty", pdIntegration, nil)
}

// GetIntegrationPD gets all the Pagerduty Integrations from the system.
func (client *Client) GetIntegrationPD() (*integrationPD, error) {
	var out integrationPD
	if err := client.doJsonRequest("GET", "/v1/integration/pagerduty", nil, &out); err != nil {
		return nil, err
	}

	return &out, nil
}

// DeleteIntegrationPD remove the PD Integration from the system.
func (client *Client) DeleteIntegrationPD() error {
	return client.doJsonRequest("DELETE", "/v1/integration/pagerduty", nil, nil)
}
