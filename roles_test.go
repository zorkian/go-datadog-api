package datadog

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strings"
	"testing"
	"time"
)

func TestClient_ListRoles(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/roles/list_roles_response.json")
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

	roles, err := datadogClient.ListRoles(10, 0, SortNameAsc, "")
	if err != nil {
		t.Fatal(err)
	}

	if roles == nil {
		t.Fatalf("Expected roles to be returned, but got nil")
	}

	expectedTotalCount := 3
	if totalCount := roles.RoleMetadata.Page.GetTotalCount(); expectedTotalCount != totalCount {
		t.Fatalf("expected %d as the total count, but got %d", expectedTotalCount, totalCount)
	}

	expectedFilteredCount := 3
	if filteredCount := roles.RoleMetadata.Page.GetTotalFilteredCount(); expectedFilteredCount != filteredCount {
		t.Fatalf("expected %d as the filtered count, but got %d", expectedFilteredCount, filteredCount)
	}

	if totalInList := len(roles.RoleData); expectedFilteredCount != totalInList {
		t.Fatalf("expected %d as the total number of roles in the list, but got %d", expectedFilteredCount, totalInList)
	}
}

func TestClient_ListRoles_Filtered(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/roles/list_roles_filtered_response.json")
		if err != nil {
			t.Fatal(err)
		}
		_, _ = w.Write(response)

		expectedFilterString := "filter=Datadog+Admin+Role"
		queryParams := strings.Split(r.RequestURI, "?")[1]

		var actualFilterString string

		for _, param := range strings.Split(queryParams, "&") {
			if strings.Contains(param, "filter=") {
				actualFilterString = param
			}
		}

		if expectedFilterString != actualFilterString {
			t.Fatalf("expected filter string to be '%s', but was '%s'", expectedFilterString, actualFilterString)
		}
	}))
	defer ts.Close()

	datadogClient := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	roles, err := datadogClient.ListRoles(10, 0, SortNameAsc, "Datadog Admin Role")
	if err != nil {
		t.Fatal(err)
	}

	if roles == nil {
		t.Fatalf("Expected roles to be returned, but got nil")
	}

	expectedTotalCount := 3
	if totalCount := roles.RoleMetadata.Page.GetTotalCount(); expectedTotalCount != totalCount {
		t.Fatalf("expected %d as the total count, but got %d", expectedTotalCount, totalCount)
	}

	expectedFilteredCount := 1
	if filteredCount := roles.RoleMetadata.Page.GetTotalFilteredCount(); expectedFilteredCount != filteredCount {
		t.Fatalf("expected %d as the filtered count, but got %d", expectedFilteredCount, filteredCount)
	}

	if totalInList := len(roles.RoleData); expectedFilteredCount != totalInList {
		t.Fatalf("expected %d as the total number of roles in the list, but got %d", expectedFilteredCount, totalInList)
	}

	ValidateAdminRole(roles.RoleData[0], t)
}

func TestClient_ListRoles_NegativePageSize(t *testing.T) {
	datadogClient := Client{
		baseUrl:    "",
		HttpClient: http.DefaultClient,
	}

	roles, err := datadogClient.ListRoles(-1, 0, SortNameAsc, "")
	if roles != nil {
		t.Fatalf("expected an error and got roles instead")
	}

	expectedError := "invalid page size, Value of 'page_size' should be 1 or more"
	if err == nil {
		t.Fatalf("expected error '%s', but no error was thrown", expectedError)
	} else if expectedError != err.Error() {
		t.Fatalf("expected error to be '%s', but was '%s'", expectedError, err)
	}
}

func TestClient_ListRoles_NegativePageNumber(t *testing.T) {
	datadogClient := Client{
		baseUrl:    "",
		HttpClient: http.DefaultClient,
	}

	roles, err := datadogClient.ListRoles(10, -1, SortNameAsc, "")
	if roles != nil {
		t.Fatalf("expected an error and got roles instead")
	}

	expectedError := "invalid page number, Value of 'page_number' should be 0 or more"
	if err == nil {
		t.Fatalf("expected error '%s', but no error was thrown", expectedError)
	} else if expectedError != err.Error() {
		t.Fatalf("expected error to be '%s', but was '%s'", expectedError, err)
	}
}

func TestClient_GetRoleById(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/roles/get_role_by_id_response.json")
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

	role, err := datadogClient.GetRoleById("e85b03a3-42b5-11ea-a78a-874cf4ed7ee4")
	if err != nil {
		t.Fatal(err)
	}

	if role == nil {
		t.Fatalf("Expected role to be returned, but it was nil")
	}

	ValidateAdminRole(role, t)
}

func TestClient_CreateRole(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/roles/create_role_response.json")
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

	role, err := datadogClient.CreateRole("Test Role")
	if err != nil {
		t.Fatal(err)
	}

	if role == nil {
		t.Fatalf("expected role to be returned, but was nil")
	}

	expectedId := "3b4d1518-5dce-11ea-ae07-3f6422573100"
	if actualId := role.GetId(); expectedId != actualId {
		t.Fatalf("expected id to be '%s', but was '%s'", expectedId, actualId)
	}
}

// Helper Functions

func ValidateAdminRole(roleData *Role, t *testing.T) {
	expectedType := "roles"
	if actualType := roleData.GetType(); expectedType != actualType {
		t.Fatalf("expected type to be '%s', but it was '%s'", expectedType, actualType)
	}

	expectedId := "e85b03a3-42b5-11ea-a78a-874cf4ed7ee4"
	if actualId := roleData.GetId(); expectedId != actualId {
		t.Fatalf("expected id to be '%s', but it was '%s'", expectedId, actualId)
	}

	if !roleData.HasAttributes() {
		t.Fatalf("expected role to have attributes, but they were nil")
	}

	roleAttributes := roleData.Attributes

	expectedName := "Datadog Admin Role"
	if actualName := roleAttributes.GetName(); expectedName != actualName {
		t.Fatalf("expected name to be '%s', but it was '%s'", expectedName, actualName)
	}

	expectedCreatedTime, _ := time.Parse(time.RFC3339, "2020-01-29T16:39:24.321785+00:00")
	if actualCreatedTime := roleAttributes.GetCreatedAt(); !expectedCreatedTime.Equal(actualCreatedTime) {
		t.Fatalf("expected created time to be '%s', but it was '%s'", expectedCreatedTime, actualCreatedTime)
	}

	expectedModifiedTime, _ := time.Parse(time.RFC3339, "2020-01-29T16:39:24.321785+00:00")
	if actualModifiedTime := roleAttributes.GetModifiedAt(); !expectedModifiedTime.Equal(actualModifiedTime) {
		t.Fatalf("expected modified time to be '%s', but it was '%s'", expectedModifiedTime, actualModifiedTime)
	}

	expectedUserCount := 1
	if actualUserCount := roleAttributes.GetUserCount(); expectedUserCount != actualUserCount {
		t.Fatalf("expected user count to be %d, but it was %d", expectedUserCount, actualUserCount)
	}

	if !roleData.HasRelationships() {
		t.Fatalf("expected role to have relationships, but they were nil")
	}

	if !roleData.Relationships.HasPermissions() {
		t.Fatalf("expected role relationships to have permissions, but they were nil")
	}

	permissionsData := roleData.Relationships.Permissions.Data

	expectedPermissionsNumber := 13
	if actualPermissionsNumber := len(permissionsData); actualPermissionsNumber != expectedPermissionsNumber {
		t.Fatalf("expected to have %d permissions, but got %d", expectedPermissionsNumber, actualPermissionsNumber)
	}

	expectedPermissionType := "permissions"
	if actualPermissionType := permissionsData[0].GetType(); expectedPermissionType != actualPermissionType {
		t.Fatalf("expected permission type to be '%s', but it was '%s'", expectedPermissionType, actualPermissionType)
	}

	uuidRegex := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	if actualPermissionId := permissionsData[0].GetId(); uuidRegex.MatchString(actualPermissionId) {
		t.Fatalf("expected permission id to be an uuid, but it was '%s'", actualPermissionId)
	}
}
