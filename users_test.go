package datadog

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/users_response.json")
		if err != nil {
			t.Fatal(err)
		}
		w.Write(response)
	}))
	defer ts.Close()

	datadogClient := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	user, err := datadogClient.GetUser("test@datadoghq.com")
	if err != nil {
		t.Fatal(err)
	}

	expectedHandle := "test@datadoghq.com"
	if handle := user.GetHandle(); handle != expectedHandle {
		t.Fatalf("expect handle %s. Got %s", expectedHandle, handle)
	}

	expectedName := "test user"
	if name := user.GetName(); name != expectedName {
		t.Fatalf("expect name %s. Got %s", expectedName, name)
	}

	expectedEmail := "test@datadoghq.com"
	if email := user.GetEmail(); email != expectedEmail {
		t.Fatalf("expect email %s. Got %s", expectedEmail, email)
	}

	expectedAccessRole := "st"
	if accessRole := user.GetAccessRole(); accessRole != expectedAccessRole {
		t.Fatalf("expect access role %s. Got %s", expectedAccessRole, accessRole)
	}

	expectedIsAdmin := false
	if isAdmin := user.GetIsAdmin(); isAdmin != expectedIsAdmin {
		t.Fatalf("expect is_admin %t. Got %v", expectedIsAdmin, isAdmin)
	}

	expectedIsVerified := true
	if isVerified := user.GetVerified(); isVerified != expectedIsVerified {
		t.Fatalf("expect is_verified %t. Got %v", expectedIsVerified, isVerified)
	}

	expectedIsDisabled := false
	if isVerified := user.GetDisabled(); isVerified != expectedIsDisabled {
		t.Fatalf("expect is_disabled %t. Got %v", expectedIsDisabled, isVerified)
	}

	expectedRole := ""
	if role := user.GetRole(); role != expectedRole {
		t.Fatalf("expect role %s. Got %s", expectedRole, role)
	}
}
