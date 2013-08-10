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
	"errors"
	"io/ioutil"
	"net/http"
)

// uriForAPI is to be called with something like "/v1/events" and it will give
// the proper request URI to be posted to.
func (self *Client) uriForAPI(api string) string {
	return "https://app.datadoghq.com/api" + api + "?api_key=" +
		self.apiKey + "&application_key=" + self.appKey
}

// doJsonRequest is the simplest type of request: a method on a URI that returns
// some JSON result which we unmarshal into the passed interface.
func (self *Client) doJsonRequest(method, api string, out interface{}) error {
	req, err := http.NewRequest(method, self.uriForAPI(api), nil)
	if err != nil {
		return err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return errors.New("API error: " + resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// If we got no body, by default let's just make an empty JSON dict. This
	// saves us some work in other parts of the code.
	if len(body) == 0 {
		body = []byte{'{', '}'}
	}

	err = json.Unmarshal(body, &out)
	if err != nil {
		return err
	}
	return nil
}

// doSimpleRequest performs a requested method on the API and returns the status
// code and an error message.
func (self *Client) doSimpleRequest(method, api string) (int, error) {
	req, err := http.NewRequest(method, self.uriForAPI(api), nil)
	if err != nil {
		return 0, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode != 200 {
		return resp.StatusCode, errors.New("API error: " + resp.Status)
	}

	return resp.StatusCode, nil
}
