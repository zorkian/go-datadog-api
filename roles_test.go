package datadog

import (
	"encoding/json"
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
	roleName := "Test Role"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if response, err := ioutil.ReadFile("./tests/fixtures/roles/create_role_response.json"); err != nil {
			t.Fatal(err)
		} else {
			_, _ = w.Write(response)
		}

		expectedType := "roles"
		if requestBody, err := ioutil.ReadAll(r.Body); err != nil {
			t.Fatal(err)
		} else {
			var roleRequest map[string]*Role
			if err := json.Unmarshal(requestBody, &roleRequest); err != nil {
				t.Fatal(err)
			} else if actualType := roleRequest["data"].GetType(); expectedType != actualType {
				t.Fatalf("expected type to be '%s', but got '%s'", expectedType, actualType)
			} else if actualRoleName := roleRequest["data"].Attributes.GetName(); actualRoleName != roleName {
				t.Fatalf("expected role name to be '%s', but got '%s'", roleName, actualRoleName)
			}
		}
	}))
	defer ts.Close()

	datadogClient := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	role, err := datadogClient.CreateRole(roleName)
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

func TestClient_UpdateRoleName(t *testing.T) {
	roleId := "3b4d1518-5dce-11ea-ae07-3f6422573100"
	roleName := "Test Role Revised"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if response, err := ioutil.ReadFile("./tests/fixtures/roles/create_role_response.json"); err != nil {
			t.Fatal(err)
		} else {
			_, _ = w.Write(response)
		}

		expectedType := "roles"
		if requestBody, err := ioutil.ReadAll(r.Body); err != nil {
			t.Fatal(err)
		} else {
			var roleRequest map[string]*Role
			if err := json.Unmarshal(requestBody, &roleRequest); err != nil {
				t.Fatal(err)
			} else if actualType := roleRequest["data"].GetType(); expectedType != actualType {
				t.Fatalf("expected type to be '%s', but got '%s'", expectedType, actualType)
			} else if actualRoleId := roleRequest["data"].GetId(); roleId != actualRoleId {
				t.Fatalf("expected role id to be '%s', but got '%s'", roleId, actualRoleId)
			} else if actualRoleName := roleRequest["data"].Attributes.GetName(); roleName != actualRoleName {
				t.Fatalf("expected role name to be '%s', but got '%s'", roleName, actualRoleName)
			}
		}
	}))
	defer ts.Close()

	datadogClient := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	role, err := datadogClient.UpdateRoleName(roleId, roleName)
	if err != nil {
		t.Fatal(err)
	}

	if role == nil {
		t.Fatalf("expected role to be returned, but was nil")
	}

	expectedId := roleId
	if actualId := role.GetId(); expectedId != actualId {
		t.Fatalf("expected id to be '%s', but was '%s'", expectedId, actualId)
	}
}

func TestClient_DeleteRole(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(204)
		_, _ = w.Write(make([]byte, 0))
	}))
	defer ts.Close()

	datadogClient := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	err := datadogClient.DeleteRole("3b4d1518-5dce-11ea-ae07-3f6422573100")
	if err != nil {
		t.Fatal(err)
	}
}

func TestClient_ListRoleUsers(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/roles/list_role_users_response.json")
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

	users, err := datadogClient.ListRoleUsers("e85b03a3-42b5-11ea-a78a-874cf4ed7ee4", 10, 0, SortNameAsc, "")
	if err != nil {
		t.Fatal(err)
	}

	if users == nil {
		t.Fatalf("expected users to be returned, but got nil")
	}

	expectedUserTotalCount := 2
	if actualUserTotalCount := users.Meta.Page.GetTotalCount(); expectedUserTotalCount != actualUserTotalCount {
		t.Fatalf("expected users total count to be %d, but got %d", expectedUserTotalCount, actualUserTotalCount)
	}

	expectedUserFilteredCount := 2
	if actualUserFilteredCount := users.Meta.Page.GetTotalFilteredCount(); expectedUserFilteredCount != actualUserFilteredCount {
		t.Fatalf("expected users filtered count to be %d, but got %d", expectedUserFilteredCount, actualUserFilteredCount)
	}

	if actualUserCount := len(users.Data); expectedUserFilteredCount != actualUserCount {
		t.Fatalf("expected number of users returned to be %d, but got %d", expectedUserFilteredCount, actualUserCount)
	}
}

func TestClient_ListRoleUsers_Filtered(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/roles/list_role_users_filtered_response.json")
		if err != nil {
			t.Fatal(err)
		}
		_, _ = w.Write(response)

		expectedFilterString := "filter=Jane+Doe"
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

	users, err := datadogClient.ListRoleUsers("e85b03a3-42b5-11ea-a78a-874cf4ed7ee4", 10, 0, SortNameAsc, "Jane Doe")
	if err != nil {
		t.Fatal(err)
	}

	if users == nil {
		t.Fatalf("Expected roles to be returned, but got nil")
	}

	expectedUserTotalCount := 2
	if actualUserTotalCount := users.Meta.Page.GetTotalCount(); expectedUserTotalCount != actualUserTotalCount {
		t.Fatalf("expected users total count to be %d, but got %d", expectedUserTotalCount, actualUserTotalCount)
	}

	expectedUserFilteredCount := 1
	if actualUserFilteredCount := users.Meta.Page.GetTotalFilteredCount(); expectedUserFilteredCount != actualUserFilteredCount {
		t.Fatalf("expected users filtered count to be %d, but got %d", expectedUserFilteredCount, actualUserFilteredCount)
	}

	if actualUserCount := len(users.Data); expectedUserFilteredCount != actualUserCount {
		t.Fatalf("expected number of users returned to be %d, but got %d", expectedUserFilteredCount, actualUserCount)
	}

	userData := users.Data[0]

	expectedType := "users"
	if actualType := userData.GetType(); expectedType != actualType {
		t.Fatalf("expected user type to be '%s', but got '%s'", expectedType, actualType)
	}

	expectedId := "bc611b4a-4447-11ea-a78a-f7832a4ee550"
	if actualId := userData.GetId(); expectedId != actualId {
		t.Fatalf("expected user id to be '%s', but got '%s'", expectedId, actualId)
	}

	if !userData.HasAttributes() {
		t.Fatalf("expected user to have attributes, but they were nil")
	}

	userAttributes := userData.Attributes

	expectedName := "Jane Doe"
	if actualName := userAttributes.GetName(); expectedName != actualName {
		t.Fatalf("expected user name to be '%s', but got '%s'", expectedName, actualName)
	}

	expectedHandle := "jdoe"
	if actualHandle := userAttributes.GetHandle(); expectedHandle != actualHandle {
		t.Fatalf("expected user name to be '%s', but got '%s'", expectedHandle, actualHandle)
	}

	expectedCreatedAt, _ := time.Parse(time.RFC3339, "2020-01-31T16:35:48.205684+00:00")
	if actualCreatedAt := userAttributes.GetCreatedAt(); !expectedCreatedAt.Equal(actualCreatedAt) {
		t.Fatalf("expected user created timestamp to be '%s', but got '%s'", expectedCreatedAt, actualCreatedAt)
	}

	expectedEmail := "jdoe@example.com"
	if actualEmail := userAttributes.GetEmail(); expectedEmail != actualEmail {
		t.Fatalf("expected user email to be '%s', but got '%s'", expectedEmail, actualEmail)
	}

	expectedIcon := "https://secure.gravatar.com/avatar/66666666677777777888888889999999?s=48&d=retro"
	if actualIcon := userAttributes.GetIcon(); expectedIcon != actualIcon {
		t.Fatalf("expected user icon to be '%s', but got '%s'", expectedIcon, actualIcon)
	}

	expectedTitle := "DataDog Admin"
	if actualTitle := userAttributes.GetTitle(); expectedTitle != actualTitle {
		t.Fatalf("expected user title to be '%s', but got '%s'", expectedTitle, actualTitle)
	}

	expectedVerified := true
	if actualVerified := userAttributes.GetVerified(); expectedVerified != actualVerified {
		t.Fatalf("expected user verified to be %t, but got %t", expectedVerified, actualVerified)
	}

	expectedDisabled := false
	if actualDisabled := userAttributes.GetDisabled(); expectedDisabled != actualDisabled {
		t.Fatalf("expected user disabled to be %t, but got %t", expectedDisabled, actualDisabled)
	}

	expectedAllowedLoginMethodsCount := 0
	if actualAllowedLoginMethodsCount := len(userAttributes.AllowedLoginMethods); expectedAllowedLoginMethodsCount != actualAllowedLoginMethodsCount {
		t.Fatalf("expected number of user allowed login methods to be %d, but got %d", expectedAllowedLoginMethodsCount, actualAllowedLoginMethodsCount)
	}

	expectedStatus := "Active"
	if actualStatus := userAttributes.GetStatus(); expectedStatus != actualStatus {
		t.Fatalf("expected user status to be '%s', but got '%s'", expectedStatus, actualStatus)
	}

	if !userData.HasRelationships() {
		t.Fatalf("expected user to have relationships, but they were nil")
	}

	if !userData.Relationships.HasRoles() {
		t.Fatalf("expected user relationships to have roles, but they were nil")
	}

	userRoles := userData.Relationships.Roles.RoleData

	expectedNumRoles := 1
	if actualNumRoles := len(userRoles); expectedNumRoles != actualNumRoles {
		t.Fatalf("expected number of user allowed login methods to be %d, but got %d", expectedNumRoles, actualNumRoles)
	}

	userRoleData := userRoles[0]

	expectedRoleType := "roles"
	if actualRoleType := userRoleData.GetType(); expectedRoleType != actualRoleType {
		t.Fatalf("expected user role type to be '%s', but it was '%s'", expectedRoleType, actualRoleType)
	}

	expectedRoleId := "e85b03a3-42b5-11ea-a78a-874cf4ed7ee4"
	if actualRoleId := userRoleData.GetId(); expectedRoleId != actualRoleId {
		t.Fatalf("expected user role id to be '%s', but it was '%s'", expectedRoleId, actualRoleId)
	}

	if !userData.Relationships.HasOrg() {
		t.Fatalf("expected user relationships to have an org, but it was nil")
	}

	if !userData.Relationships.Org.HasData() {
		t.Fatalf("expected user org to have data, but it was nil")
	}

	userOrg := userData.Relationships.Org.Data

	expectedOrgType := "orgs"
	if actualOrgType := userOrg.GetType(); expectedOrgType != actualOrgType {
		t.Fatalf("expected org type to be '%s', but it was '%s'", expectedOrgType, actualOrgType)
	}

	expectedOrgId := "e85b03a2-42b5-11ea-a78a-ab1bc28a6c51"
	if actualOrgId := userOrg.GetId(); expectedOrgId != actualOrgId {
		t.Fatalf("expected org id to be '%s', but it was '%s'", expectedOrgId, actualOrgId)
	}
}

func TestClient_AddRoleUser(t *testing.T) {
	userId := "4014ebde-64e0-11ea-b5b2-8be7e92064db"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedType := "users"
		if requestBody, err := ioutil.ReadAll(r.Body); err != nil {
			t.Fatal(err)
		} else {
			var userRequest map[string]*Permission
			if err := json.Unmarshal(requestBody, &userRequest); err != nil {
				t.Fatal(err)
			} else if actualType := userRequest["data"].GetType(); expectedType != actualType {
				t.Fatalf("expected type to be '%s', but got '%s'", expectedType, actualType)
			} else if actualUserId := userRequest["data"].GetId(); userId != actualUserId {
				t.Fatalf("expected user id to be '%s', but got '%s'", userId, actualUserId)
			}
		}

		response, err := ioutil.ReadFile("./tests/fixtures/roles/list_role_users_response.json")
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

	users, err := datadogClient.AddRoleUser("e85b03a3-42b5-11ea-a78a-874cf4ed7ee4", "4014ebde-64e0-11ea-b5b2-8be7e92064db")
	if err != nil {
		t.Fatal(err)
	}

	if users == nil {
		t.Fatalf("expected users to be returned, but got nil")
	}

	expectedUserTotalCount := 2
	if actualUserTotalCount := users.Meta.Page.GetTotalCount(); expectedUserTotalCount != actualUserTotalCount {
		t.Fatalf("expected users total count to be %d, but got %d", expectedUserTotalCount, actualUserTotalCount)
	}

	expectedUserFilteredCount := 2
	if actualUserFilteredCount := users.Meta.Page.GetTotalFilteredCount(); expectedUserFilteredCount != actualUserFilteredCount {
		t.Fatalf("expected users filtered count to be %d, but got %d", expectedUserFilteredCount, actualUserFilteredCount)
	}

	if actualUserCount := len(users.Data); expectedUserFilteredCount != actualUserCount {
		t.Fatalf("expected number of users returned to be %d, but got %d", expectedUserFilteredCount, actualUserCount)
	}
}

func TestClient_RemoveRoleUser(t *testing.T) {
	userId := "4014ebde-64e0-11ea-b5b2-8be7e92064db"
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expectedType := "users"
		if requestBody, err := ioutil.ReadAll(r.Body); err != nil {
			t.Fatal(err)
		} else {
			var userRequest map[string]*Permission
			if err := json.Unmarshal(requestBody, &userRequest); err != nil {
				t.Fatal(err)
			} else if actualType := userRequest["data"].GetType(); expectedType != actualType {
				t.Fatalf("expected type to be '%s', but got '%s'", expectedType, actualType)
			} else if actualUserId := userRequest["data"].GetId(); userId != actualUserId {
				t.Fatalf("expected user id to be '%s', but got '%s'", userId, actualUserId)
			}
		}

		response, err := ioutil.ReadFile("./tests/fixtures/roles/list_role_users_response.json")
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

	users, err := datadogClient.RemoveRoleUser("e85b03a3-42b5-11ea-a78a-874cf4ed7ee4", userId)
	if err != nil {
		t.Fatal(err)
	}

	if users == nil {
		t.Fatalf("expected users to be returned, but got nil")
	}

	expectedUserTotalCount := 2
	if actualUserTotalCount := users.Meta.Page.GetTotalCount(); expectedUserTotalCount != actualUserTotalCount {
		t.Fatalf("expected users total count to be %d, but got %d", expectedUserTotalCount, actualUserTotalCount)
	}

	expectedUserFilteredCount := 2
	if actualUserFilteredCount := users.Meta.Page.GetTotalFilteredCount(); expectedUserFilteredCount != actualUserFilteredCount {
		t.Fatalf("expected users filtered count to be %d, but got %d", expectedUserFilteredCount, actualUserFilteredCount)
	}

	if actualUserCount := len(users.Data); expectedUserFilteredCount != actualUserCount {
		t.Fatalf("expected number of users returned to be %d, but got %d", expectedUserFilteredCount, actualUserCount)
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
