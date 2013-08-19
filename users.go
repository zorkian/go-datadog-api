/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

import (
	"errors"
)

// reqInviteUsers contains email addresses to send invitations to.
type reqInviteUsers struct {
	Emails []string `json:"emails"`
}

// InviteUsers takes a slice of email addresses and sends invitations to them.
func (self *Client) InviteUsers(emails []string) error {
	return errors.New("datadog API docs don't list the endpoint")

	//	return self.doJsonRequest("POST", "/v1/alert",
	//		reqInviteUsers{Emails: emails}, nil)
}
