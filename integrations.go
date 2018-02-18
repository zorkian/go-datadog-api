/*
 * Datadog API for Go
 *
 * Please see the included LICENSE file for licensing information.
 *
 * Copyright 2013 by authors and contributors.
 */

package datadog

type SlackChannel struct {
	TransferAllUserComments *string `json:"transfer_all_user_comments"`
	ChannelName *string `json:"channel_name"`
	Account *string `json:"account"`
}

type SlackServiceHooks struct {
	Url *string `json:"url"`
	Account *string `json:"account"`
}

type reqSlackIntegration struct {
	Channels []SlackChannel `json:"channels,omitempty"`
	ServiceHooks []SlackServiceHooks`json:"channels,service_hooks"`
}


// Slack integration returns a list of all integrations with slack created on this account.
func (client *Client) GetSlackIntegrations() ([]SlackChannel,[]SlackServiceHooks , error) {
	var out reqSlackIntegration
	if err := client.doJsonRequest("GET", "/v1/integration/slack", nil, &out); err != nil {
		return nil, nil, err
	}
	return out.Channels, out.ServiceHooks, nil
}

