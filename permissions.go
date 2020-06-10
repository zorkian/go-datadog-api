package datadog

import (
	"fmt"
	"time"
)

type PermissionsResponse struct {
	Data []*Permission `json:"data,omitempty"`
}

type Permission struct {
	Type       *string               `json:"type"`
	Id         *string               `json:"id,omitempty"`
	Attributes *PermissionAttributes `json:"attributes,omitempty"`
}

type PermissionAttributes struct {
	Created     *time.Time `json:"created,omitempty"`
	Description *string    `json:"description,omitempty"`
	DisplayName *string    `json:"display_name,omitempty"`
	DisplayType *string    `json:"display_type,omitempty"`
	GroupName   *string    `json:"group_name,omitempty"`
	Name        *string    `json:"name,omitempty"`
	Restricted  *bool      `json:"restricted,omitempty"`
}

func (client *Client) ListPermissions() ([]*Permission, error) {
	var permissionsResponse PermissionsResponse

	if err := client.doJsonRequest("GET", "/v2/permissions", nil, &permissionsResponse); err != nil {
		return nil, err
	}

	return permissionsResponse.Data, nil
}

func (client *Client) ListRolePermissions(roleId string) ([]*Permission, error) {
	var permissionsResponse PermissionsResponse

	uri := fmt.Sprintf("/v2/roles/%s/permissions", roleId)

	if err := client.doJsonRequest("GET", uri, nil, &permissionsResponse); err != nil {
		return nil, err
	}

	return permissionsResponse.Data, nil
}

func (client *Client) GrantRolePermission(roleId string, permissionId string) ([]*Permission, error) {
	var permissionsResponse PermissionsResponse

	uri := fmt.Sprintf("/v2/roles/%s/permissions", roleId)

	permissionRequest := Permission{
		Type: String("permissions"),
		Id:   String(permissionId),
	}

	data := DataWrapper{Data: permissionRequest}

	if err := client.doJsonRequest("POST", uri, data, &permissionsResponse); err != nil {
		return nil, err
	}

	return permissionsResponse.Data, nil
}

func (client *Client) RevokeRolePermission(roleId string, permissionId string) ([]*Permission, error) {
	var permissionsResponse PermissionsResponse

	uri := fmt.Sprintf("/v2/roles/%s/permissions", roleId)

	permissionRequest := Permission{
		Type: String("permissions"),
		Id:   String(permissionId),
	}

	data := DataWrapper{Data: permissionRequest}

	if err := client.doJsonRequest("DELETE", uri, data, &permissionsResponse); err != nil {
		return nil, err
	}

	return permissionsResponse.Data, nil
}
