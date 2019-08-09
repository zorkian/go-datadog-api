/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2019 by authors and contributors.
 */

package datadog

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

// IP ranges US: https://ip-ranges.datadoghq.com
// EU: https://ip-ranges.datadoghq.eu
// Same structure
type IPRanges struct {
	Agents     map[string][]string `json:"agents"`
	API        map[string][]string `json:"api"`
	Apm        map[string][]string `json:"apm"`
	Logs       map[string][]string `json:"logs"`
	Process    map[string][]string `json:"process"`
	Synthetics map[string][]string `json:"synthetics"`
	Webhooks   map[string][]string `json:"webhooks"`
}

// GetIPRanges returns all IP addresses by section: agents, api, apm, logs, process, synthetics, webhooks
func (client *Client) GetIPRanges() (*IPRanges, error) {

	baseURL := client.GetBaseUrl()

	var urlIPRanges string

	// Check if we're looking at the IP ranges for EU or US using the BaseURL domain
	if strings.Contains(baseURL, ".eu") == true {
		urlIPRanges = "https://ip-ranges.datadoghq.eu"
	} else {
		urlIPRanges = "https://ip-ranges.datadoghq.com"
	}

	// HTTP request to inspect the response's body
	clientHTTP := &http.Client{}

	req, err := http.NewRequest("GET", urlIPRanges, nil)
	if err != nil {
		return nil, fmt.Errorf("Error creating request: %s", err)
	}

	resp, err := clientHTTP.Do(req)
	if err != nil {
		return nil, fmt.Errorf("Error during making a request: %s", urlIPRanges)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP request error. Response code: %d", resp.StatusCode)
	}

	body, readErr := ioutil.ReadAll(resp.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	// Unmarshal the response body into the IPRanges struct
	var ipRanges IPRanges
	json.Unmarshal([]byte(body), &ipRanges)

	return &ipRanges, nil
}
