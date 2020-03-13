package datadog

import (
	"fmt"
	"time"
)

type Sort string

const (
	SortNameAsc        Sort = "name"
	SortNameDesc       Sort = "-name"
	SortModifiedAtAsc  Sort = "modified_at"
	SortModifiedAtDesc Sort = "-modified_at"
	SortUserCountAsc   Sort = "user_count"
	SortUserCountDesc  Sort = "-user_count"
)

type ListRolesResponse struct {
	RoleMetadata *RoleMetadata `json:"meta,omitempty"`
	RoleData     []*Role       `json:"data,omitempty"`
}

type RoleResponse struct {
	RoleData *Role `json:"data,omitempty"`
}

type RoleMetadata struct {
	Page *Page `json:"page,omitempty"`
}

type Page struct {
	TotalFilteredCount *int `json:"total_filtered_count,omitempty"`
	TotalCount         *int `json:"total_count,omitempty"`
}

type Role struct {
	Type          *string            `json:"type,omitempty"`
	Id            *string            `json:"id,omitempty"`
	Attributes    *RoleAttributes    `json:"attributes,omitempty"`
	Relationships *RoleRelationships `json:"relationships,omitempty"`
}

type RoleAttributes struct {
	Name       *string    `json:"name,omitempty"`
	CreatedAt  *time.Time `json:"created_at,omitempty"`
	ModifiedAt *time.Time `json:"modified_at,omitempty"`
	UserCount  *int       `json:"user_count,omitempty"`
}

type RoleRelationships struct {
	Permissions *PermissionsResponse `json:"permissions,omitempty"`
}

type RoleUsersResponse struct {
	Meta *RoleMetadata `json:"meta,omitempty"`
	Data []*UserV2     `json:"data,omitempty"`
}

func (client *Client) ListRoles(pageSize int, pageNumber int, sort Sort, filter string) (*ListRolesResponse, error) {
	var roleResponse ListRolesResponse

	if pageSize < 1 {
		return nil, fmt.Errorf("invalid page size, Value of 'page_size' should be 1 or more")
	}

	if pageNumber < 0 {
		return nil, fmt.Errorf("invalid page number, Value of 'page_number' should be 0 or more")
	}

	uri := fmt.Sprintf("/v2/roles?page[size]=%d&page[number]=%d&sort=%s&filter=%s",
		pageSize, pageNumber, sort, filter)

	if err := client.doJsonRequest("GET", uri, nil, &roleResponse); err != nil {
		return nil, err
	}

	return &roleResponse, nil
}

func (client *Client) GetRoleById(id string) (*Role, error) {
	var roleResponse RoleResponse

	uri := fmt.Sprintf("/v2/roles/%s", id)

	if err := client.doJsonRequest("GET", uri, nil, &roleResponse); err != nil {
		return nil, err
	}

	return roleResponse.RoleData, nil
}

func (client *Client) CreateRole(name string) (*Role, error) {

	var roleResponse RoleResponse

	roleRequest := Role{
		Type:       String("roles"),
		Attributes: &RoleAttributes{Name: String(name)},
	}

	data := DataWrapper{Data: roleRequest}

	if err := client.doJsonRequest("POST", "/v2/roles", data, &roleResponse); err != nil {
		return nil, err
	}

	return roleResponse.RoleData, nil
}

func (client *Client) UpdateRoleName(id string, name string) (*Role, error) {
	var roleResponse RoleResponse

	roleRequest := Role{
		Type:       String("roles"),
		Id:         String(id),
		Attributes: &RoleAttributes{Name: String(name)},
	}

	data := DataWrapper{Data: roleRequest}

	uri := fmt.Sprintf("/v2/roles/%s", id)

	if err := client.doJsonRequest("PATCH", uri, data, &roleResponse); err != nil {
		return nil, err
	}

	return roleResponse.RoleData, nil
}

func (client *Client) DeleteRole(id string) error {
	uri := fmt.Sprintf("/v2/roles/%s", id)

	if err := client.doJsonRequest("DELETE", uri, nil, nil); err != nil {
		return err
	}

	return nil
}

func (client *Client) ListRoleUsers(roleId string, pageSize int, pageNumber int, sort Sort, filter string) (*RoleUsersResponse, error) {
	var roleUsersResponse RoleUsersResponse

	if pageSize < 1 {
		return nil, fmt.Errorf("invalid page size, Value of 'page_size' should be 1 or more")
	}

	if pageNumber < 0 {
		return nil, fmt.Errorf("invalid page number, Value of 'page_number' should be 0 or more")
	}

	uri := fmt.Sprintf("/v2/roles/%s/users?page[size]=%d&page[number]=%d&sort=%s&filter=%s",
		roleId, pageSize, pageNumber, sort, filter)

	if err := client.doJsonRequest("GET", uri, nil, &roleUsersResponse); err != nil {
		return nil, err
	}

	return &roleUsersResponse, nil
}

func (client *Client) AddRoleUser(roleId string, userId string) (*RoleUsersResponse, error) {
	var roleUsersResponse RoleUsersResponse

	uri := fmt.Sprintf("/v2/roles/%s/users", roleId)

	userRequest := UserV2{
		Type: String("users"),
		Id:   String(userId),
	}

	data := DataWrapper{Data: userRequest}

	if err := client.doJsonRequest("POST", uri, data, &roleUsersResponse); err != nil {
		return nil, err
	}

	return &roleUsersResponse, nil
}

func (client *Client) RemoveRoleUser(roleId string, userId string) (*RoleUsersResponse, error) {
	var roleUsersResponse RoleUsersResponse

	uri := fmt.Sprintf("/v2/roles/%s/users", roleId)

	userRequest := UserV2{
		Type: String("users"),
		Id:   String(userId),
	}

	data := DataWrapper{Data: userRequest}

	if err := client.doJsonRequest("DELETE", uri, data, &roleUsersResponse); err != nil {
		return nil, err
	}

	return &roleUsersResponse, nil
}
