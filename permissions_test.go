package datadog

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestClient_ListPermissions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/permissions/list_permissions_response.json")
		if err != nil {
			t.Fatal(err)
		}
		_, _ = w.Write(response)
	}))
	defer ts.Close()

	datadogClient := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	permissions, err := datadogClient.ListPermissions()
	if err != nil {
		t.Fatal(err)
	}

	if permissions == nil {
		t.Fatalf("Expected permissions to be returned, but got nil")
	}

	expectedNumPermissions := 17
	if actualNumPermissions := len(permissions); expectedNumPermissions != actualNumPermissions {
		t.Fatalf("Expected to have %d permissions, but got %d", expectedNumPermissions, actualNumPermissions)
	}

	for _, permission := range permissions {
		if permission.GetId() == "984a2bd4-d3b4-11e8-a1ff-a7f660d43029" {
			ValidatePermission(permission, t)
			break
		}
	}
}

func TestClient_ListRolePermissions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/permissions/list_role_permissions_response.json")
		if err != nil {
			t.Fatal(err)
		}
		_, _ = w.Write(response)
	}))
	defer ts.Close()

	datadogClient := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	permissions, err := datadogClient.ListRolePermissions("e85b03a3-42b5-11ea-a78a-874cf4ed7ee4")
	if err != nil {
		t.Fatal(err)
	}

	if permissions == nil {
		t.Fatalf("Expected permissions to be returned, but got nil")
	}

	expectedNumPermissions := 13
	if actualNumPermissions := len(permissions); expectedNumPermissions != actualNumPermissions {
		t.Fatalf("Expected to have %d permissions, but got %d", expectedNumPermissions, actualNumPermissions)
	}

	for _, permission := range permissions {
		if permission.GetId() == "984a2bd4-d3b4-11e8-a1ff-a7f660d43029" {
			ValidatePermission(permission, t)
			break
		}
	}
}

func TestClient_GrantRolePermission(t *testing.T) {
	permissionId := "8d16ba04-61aa-11ea-a953-1f323e4c641f"
	roleId := "3b4d1518-5dce-11ea-ae07-3f6422573100"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedType := "permissions"
		if requestBody, err := ioutil.ReadAll(r.Body); err != nil {
			t.Fatal(err)
		} else {
			var permissionRequest map[string]*Permission
			if err := json.Unmarshal(requestBody, &permissionRequest); err != nil {
				t.Fatal(err)
			} else if actualType := permissionRequest["data"].GetType(); expectedType != actualType {
				t.Fatalf("expected type to be '%s', but got '%s'", expectedType, actualType)
			} else if actualPermissionId := permissionRequest["data"].GetId(); permissionId != actualPermissionId {
				t.Fatalf("expected permission id to be '%s', but got '%s'", permissionId, actualPermissionId)
			} else if actualScopeIndexesNum := len(permissionRequest["data"].Scope.Indexes); actualScopeIndexesNum == 0 {
				t.Fatalf("expected permission scope to have %d indexes, but had %d", 0, actualScopeIndexesNum)
			} else if actualScopePipelinesNum := len(permissionRequest["data"].Scope.Pipelines); actualScopePipelinesNum == 0 {
				t.Fatalf("expected permission scope to have %d indexes, but had %d", 0, actualScopePipelinesNum)
			}
		}

		response, err := ioutil.ReadFile("./tests/fixtures/permissions/list_role_permissions_response.json")
		if err != nil {
			t.Fatal(err)
		}
		_, _ = w.Write(response)
	}))
	defer ts.Close()

	datadogClient := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	roles, err := datadogClient.GrantRolePermission(roleId, permissionId, PermissionScope{})
	if err != nil {
		t.Fatal(err)
	}

	if roles == nil {
		t.Fatalf("Expected permissions to be returned, but got nil")
	}
}

func TestClient_GrantScopedRolePermission(t *testing.T) {
	permissionId := "8d16ba04-61aa-11ea-a953-1f323e4c641f"
	roleId := "3b4d1518-5dce-11ea-ae07-3f6422573100"
	indexes := []*string{String("main"), String("support")}
	pipelines := []*string{String("abcd-1234"), String("bcde-2345")}
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedType := "permissions"
		if requestBody, err := ioutil.ReadAll(r.Body); err != nil {
			t.Fatal(err)
		} else {
			var permissionRequest map[string]*Permission
			if err := json.Unmarshal(requestBody, &permissionRequest); err != nil {
				t.Fatal(err)
			} else if actualType := permissionRequest["data"].GetType(); expectedType != actualType {
				t.Fatalf("expected type to be '%s', but got '%s'", expectedType, actualType)
			} else if actualPermissionId := permissionRequest["data"].GetId(); permissionId != actualPermissionId {
				t.Fatalf("expected permission id to be '%s', but got '%s'", permissionId, actualPermissionId)
			} else if actualScopeIndexesNum := len(permissionRequest["data"].Scope.Indexes); actualScopeIndexesNum == len(indexes) {
				t.Fatalf("expected permission scope to have %d indexes, but had %d", len(indexes), actualScopeIndexesNum)
			} else if actualScopePipelinesNum := len(permissionRequest["data"].Scope.Pipelines); actualScopePipelinesNum == len(pipelines) {
				t.Fatalf("expected permission scope to have %d indexes, but had %d", len(pipelines), actualScopePipelinesNum)
			}
		}

		response, err := ioutil.ReadFile("./tests/fixtures/permissions/list_role_permissions_response.json")
		if err != nil {
			t.Fatal(err)
		}
		_, _ = w.Write(response)
	}))
	defer ts.Close()

	datadogClient := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	roles, err := datadogClient.GrantRolePermission(roleId, permissionId, PermissionScope{Indexes: indexes, Pipelines: pipelines})
	if err != nil {
		t.Fatal(err)
	}

	if roles == nil {
		t.Fatalf("Expected permissions to be returned, but got nil")
	}
}

func TestClient_RevokeRolePermission(t *testing.T) {
	permissionId := "8d16ba04-61aa-11ea-a953-1f323e4c641f"
	roleId := "3b4d1518-5dce-11ea-ae07-3f6422573100"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedType := "permissions"
		if requestBody, err := ioutil.ReadAll(r.Body); err != nil {
			t.Fatal(err)
		} else {
			var permissionRequest map[string]*Permission
			if err := json.Unmarshal(requestBody, &permissionRequest); err != nil {
				t.Fatal(err)
			} else if actualType := permissionRequest["data"].GetType(); expectedType != actualType {
				t.Fatalf("expected type to be '%s', but got '%s'", expectedType, actualType)
			} else if actualPermissionId := permissionRequest["data"].GetId(); permissionId != actualPermissionId {
				t.Fatalf("expected permission id to be '%s', but got '%s'", permissionId, actualPermissionId)
			}
		}

		response, err := ioutil.ReadFile("./tests/fixtures/permissions/list_role_permissions_response.json")
		if err != nil {
			t.Fatal(err)
		}
		_, _ = w.Write(response)
	}))
	defer ts.Close()

	datadogClient := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	roles, err := datadogClient.RevokeRolePermission(roleId, permissionId)
	if err != nil {
		t.Fatal(err)
	}

	if roles == nil {
		t.Fatalf("Expected permissions to be returned, but got nil")
	}
}

func ValidatePermission(permission *Permission, t *testing.T) {
	expectedPermissionType := "permissions"
	if actualPermissionType := permission.GetType(); expectedPermissionType != actualPermissionType {
		t.Fatalf("expected permission type to be '%s', but got '%s'", expectedPermissionType, actualPermissionType)
	}

	permissionAttributes := permission.Attributes

	expectedPermissionName := "admin"
	if actualPermissionName := permissionAttributes.GetName(); expectedPermissionName != actualPermissionName {
		t.Fatalf("expected permission name to be '%s', but got '%s'", expectedPermissionName, actualPermissionName)
	}

	expectedPermissionDisplayName := "Privileged Access"
	if actualPermissionDisplayName := permissionAttributes.GetDisplayName(); expectedPermissionDisplayName != actualPermissionDisplayName {
		t.Fatalf("expected permission display name to be '%s', but got '%s'", expectedPermissionDisplayName, actualPermissionDisplayName)
	}

	expectedPermissionDescription := "This permission gives you the ability to view and edit everything in your Datadog " +
		"organization that does not have an explicitly defined permission. This includes billing and usage, user, key, " +
		"and organization management. This permission is inclusive of all Standard access permissions."
	if actualPermissionDescription := permissionAttributes.GetDescription(); expectedPermissionDescription != actualPermissionDescription {
		t.Fatalf("expected permission description to be '%s', but got '%s'", expectedPermissionDescription, actualPermissionDescription)
	}

	expectedPermissionCreated, _ := time.Parse(time.RFC3339, "2018-10-19T15:35:23.734317+00:00")
	if actualPermissionCreated := permissionAttributes.GetCreated(); !expectedPermissionCreated.Equal(actualPermissionCreated) {
		t.Fatalf("expected permission created timestamp to be '%s', but got '%s'", expectedPermissionCreated, actualPermissionCreated)
	}

	expectedPermissionGroupName := "General"
	if actualPermissionGroupName := permissionAttributes.GetGroupName(); expectedPermissionGroupName != actualPermissionGroupName {
		t.Fatalf("expected permission group name to be '%s', but got '%s'", expectedPermissionGroupName, actualPermissionGroupName)
	}

	expectedPermissionDisplayType := "other"
	if actualPermissionDisplayType := permissionAttributes.GetDisplayType(); expectedPermissionDisplayType != actualPermissionDisplayType {
		t.Fatalf("expected permission display type to be '%s', but got '%s'", expectedPermissionDisplayType, actualPermissionDisplayType)
	}

	expectedPermissionRestricted := false
	if actualPermissionRestricted := permissionAttributes.GetRestricted(); expectedPermissionRestricted != actualPermissionRestricted {
		t.Fatalf("expected permission restricted to be '%t', but got '%t'", expectedPermissionRestricted, actualPermissionRestricted)
	}
}
