package datadog

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchSyntheticsChecks(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/synthetics/checks/search_response.json")
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

	checks, err := datadogClient.SearchSyntheticsChecks("")
	if err != nil {
		t.Fatal(err)
	}

	expectedCnt := 3
	if cnt := len(checks); cnt != expectedCnt {
		t.Fatalf("expect %d checks. Got %d", expectedCnt, cnt)
	}
}

func TestGetSyntheticsCheck(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/synthetics/checks/get_response.json")
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

	c, err := datadogClient.GetSyntheticsCheck("xxx-xxx-xxx")
	if err != nil {
		t.Fatal(err)
	}

	expectedId := 1234
	if id := c.GetId(); id != expectedId {
		t.Fatalf("expect id %d. Got %d", expectedId, id)
	}

	expectedPublicId := "xxx-xxx-xxx"
	if publicId := c.GetPublicId(); publicId != expectedPublicId {
		t.Fatalf("expect public_id %s. Got %s", expectedPublicId, publicId)
	}

	expectedName := "Check on example.com"
	if name := c.GetName(); name != expectedName {
		t.Fatalf("expect name %s. Got %s", expectedName, name)
	}

	expectedType := "api"
	if checkType := c.GetType(); checkType != expectedType {
		t.Fatalf("expect type %s. Got %s", expectedType, checkType)
	}

	expectedCreatedAt := "2019-01-25T02:25:40.241032+00:00"
	if createdAt := c.GetCreatedAt(); createdAt != expectedCreatedAt {
		t.Fatalf("expect created_at %s. Got %s", expectedCreatedAt, createdAt)
	}

	expectedModifiedAt := "2019-02-09T18:11:12.801165+00:00"
	if modifiedAt := c.GetModifiedAt(); modifiedAt != expectedModifiedAt {
		t.Fatalf("expect modified_at %s. Got %s", expectedModifiedAt, modifiedAt)
	}

	expectedCheckStatus := "live"
	if checkStatus := c.GetCheckStatus(); checkStatus != expectedCheckStatus {
		t.Fatalf("expect check_status %s. Got %s", expectedCheckStatus, checkStatus)
	}

	expectedMessage := "Danger! @example@example.com"
	if message := c.GetMessage(); message != expectedMessage {
		t.Fatalf("expect message %s. Got %s", expectedMessage, message)
	}

	expectedTickEvery := 60
	options := c.GetOptions()
	if tickEvery := options.GetTickEvery(); tickEvery != expectedTickEvery {
		t.Fatalf("expect options.tick_every %d. Got %d", expectedTickEvery, tickEvery)
	}

	locations := c.Locations
	expectedLocationsCnt := 1
	if cnt := len(locations); cnt != expectedLocationsCnt {
		t.Fatalf("locations count should be %d. Got %d", expectedLocationsCnt, cnt)
	}
	expectedLocation := "aws:ap-northeast-1"
	if location := locations[0]; location != expectedLocation {
		t.Fatalf("expect location %s. Got %s", expectedLocation, location)
	}

	tags := c.Tags
	expectedTagsCnt := 1
	if cnt := len(tags); cnt != expectedTagsCnt {
		t.Fatalf("tags count should be %d. Got %d", expectedTagsCnt, cnt)
	}
	expectedTag := "example_tag"
	if tag := tags[0]; tag != expectedTag {
		t.Fatalf("expect tag %s. Got %s", expectedTag, tag)
	}

	createdBy := c.GetCreatedBy()
	assert.Equal(t, createdBy.GetEmail(), "example@example.com")
	assert.Equal(t, createdBy.GetHandle(), "example@example.com")
	assert.Equal(t, createdBy.GetId(), 123456)
	assert.Equal(t, createdBy.GetName(), "John Doe")

	modifiedBy := c.GetModifiedBy()
	assert.Equal(t, modifiedBy.GetEmail(), "example@example.com")
	assert.Equal(t, modifiedBy.GetHandle(), "example@example.com")
	assert.Equal(t, modifiedBy.GetId(), 123456)
	assert.Equal(t, modifiedBy.GetName(), "John Doe")

	config := c.GetConfig()

	request := config.GetRequest()
	assert.Equal(t, request.GetUrl(), "https://example.com/")
	assert.Equal(t, request.GetMethod(), "GET")
	assert.Equal(t, request.GetTimeout(), 30)

	assertions := config.Assertions
	assert.Equal(t, len(assertions), 3)

	assert.Equal(t, assertions[0].GetOperator(), "is")
	assert.Equal(t, assertions[0].GetProperty(), "content-type")
	assert.Equal(t, assertions[0].GetType(), "header")
	assert.Equal(t, assertions[0].Target.(string), "text/html; charset=UTF-8")

	assert.Equal(t, assertions[1].GetOperator(), "lessThan")
	assert.Equal(t, assertions[1].GetType(), "responseTime")
	assert.Equal(t, int(assertions[1].Target.(float64)), 4000)

	assert.Equal(t, assertions[2].GetOperator(), "is")
	assert.Equal(t, assertions[2].GetType(), "statusCode")
	assert.Equal(t, int(assertions[2].Target.(float64)), 200)
}
