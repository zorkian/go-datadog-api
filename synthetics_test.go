package datadog

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListSyntheticsTests(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/synthetics/tests/list_tests.json")
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

	tests, err := datadogClient.GetSyntheticsTests()
	if err != nil {
		t.Fatal(err)
	}

	expectedCnt := 3
	if cnt := len(tests); cnt != expectedCnt {
		t.Fatalf("expect %d tests. Got %d", expectedCnt, cnt)
	}
}

func TestGetSyntheticsTestApi(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/synthetics/tests/get_test_api.json")
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

	c, err := datadogClient.GetSyntheticsTest("xxx-xxx-xxx")
	if err != nil {
		t.Fatal(err)
	}

	expectedPublicId := "xxx-xxx-xxx"
	if publicId := c.GetPublicId(); publicId != expectedPublicId {
		t.Fatalf("expect public_id %s. Got %s", expectedPublicId, publicId)
	}

	expectedMonitorId := 666
	if monitorId := c.GetMonitorId(); monitorId != expectedMonitorId {
		t.Fatalf("expect monitor_id %d. Got %d", expectedMonitorId, monitorId)
	}

	expectedName := "Check on example.com"
	if name := c.GetName(); name != expectedName {
		t.Fatalf("expect name %s. Got %s", expectedName, name)
	}

	expectedType := "api"
	if actualType := c.GetType(); actualType != expectedType {
		t.Fatalf("expect type %s. Got %s", expectedType, actualType)
	}

	expectedCreatedAt := "2019-01-25T02:25:40.241032+00:00"
	if createdAt := c.GetCreatedAt(); createdAt != expectedCreatedAt {
		t.Fatalf("expect created_at %s. Got %s", expectedCreatedAt, createdAt)
	}

	expectedModifiedAt := "2019-02-09T18:11:12.801165+00:00"
	if modifiedAt := c.GetModifiedAt(); modifiedAt != expectedModifiedAt {
		t.Fatalf("expect modified_at %s. Got %s", expectedModifiedAt, modifiedAt)
	}

	expectedStatus := "live"
	if status := c.GetStatus(); status != expectedStatus {
		t.Fatalf("expect status %s. Got %s", expectedStatus, status)
	}

	expectedMessage := "Danger! @example@example.com"
	if message := c.GetMessage(); message != expectedMessage {
		t.Fatalf("expect message %s. Got %s", expectedMessage, message)
	}

	options := c.GetOptions()

	expectedTickEvery := 60
	if tickEvery := options.GetTickEvery(); tickEvery != expectedTickEvery {
		t.Fatalf("expect options.tick_every %d. Got %d", expectedTickEvery, tickEvery)
	}

	expectedMinFailureDuration := 30
	if minFailureDuration := options.GetMinFailureDuration(); minFailureDuration != expectedMinFailureDuration {
		t.Fatalf("expect options.min_failure_duration %d. Got %d", expectedMinFailureDuration, expectedMinFailureDuration)
	}

	expectedMinLocationFailed := 3
	if minLocationFailed := options.GetMinLocationFailed(); minLocationFailed != expectedMinLocationFailed {
		t.Fatalf("expect options.min_location_failed %d. Got %d", expectedMinLocationFailed, minLocationFailed)
	}

	expectedFollowRedirects := true
	if followRedirects := options.GetFollowRedirects(); followRedirects != expectedFollowRedirects {
		t.Fatalf("expect options.follow_redirects %v. Got %v", expectedFollowRedirects, followRedirects)
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

func TestGetSyntheticsTestBrowser(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/synthetics/tests/get_test_browser.json")
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

	c, err := datadogClient.GetSyntheticsTest("xxx-xxx-xxx")
	if err != nil {
		t.Fatal(err)
	}

	expectedPublicId := "xxx-xxx-xxx"
	if publicId := c.GetPublicId(); publicId != expectedPublicId {
		t.Fatalf("expect public_id %s. Got %s", expectedPublicId, publicId)
	}

	expectedMonitorId := 666
	if monitorId := c.GetMonitorId(); monitorId != expectedMonitorId {
		t.Fatalf("expect monitor_id %d. Got %d", expectedMonitorId, monitorId)
	}

	expectedName := "Check on example.com"
	if name := c.GetName(); name != expectedName {
		t.Fatalf("expect name %s. Got %s", expectedName, name)
	}

	expectedType := "browser"
	if actualType := c.GetType(); actualType != expectedType {
		t.Fatalf("expect type %s. Got %s", expectedType, actualType)
	}

	expectedCreatedAt := "2019-01-25T02:25:40.241032+00:00"
	if createdAt := c.GetCreatedAt(); createdAt != expectedCreatedAt {
		t.Fatalf("expect created_at %s. Got %s", expectedCreatedAt, createdAt)
	}

	expectedModifiedAt := "2019-02-09T18:11:12.801165+00:00"
	if modifiedAt := c.GetModifiedAt(); modifiedAt != expectedModifiedAt {
		t.Fatalf("expect modified_at %s. Got %s", expectedModifiedAt, modifiedAt)
	}

	expectedStatus := "live"
	if status := c.GetStatus(); status != expectedStatus {
		t.Fatalf("expect status %s. Got %s", expectedStatus, status)
	}

	expectedMessage := "Danger! @example@example.com"
	if message := c.GetMessage(); message != expectedMessage {
		t.Fatalf("expect message %s. Got %s", expectedMessage, message)
	}

	options := c.GetOptions()

	expectedTickEvery := 900
	if tickEvery := options.GetTickEvery(); tickEvery != expectedTickEvery {
		t.Fatalf("expect options.tick_every %d. Got %d", expectedTickEvery, tickEvery)
	}

	deviceIds := options.DeviceIds
	expectedDevicesCnt := 1
	if cnt := len(deviceIds); cnt != expectedDevicesCnt {
		t.Fatalf("device_ids count should be %d. Got %d", expectedDevicesCnt, cnt)
	}

	expectedDeviceId := "laptop_large"
	if deviceId := deviceIds[0]; deviceId != expectedDeviceId {
		t.Fatalf("expect device_id %s. Got %s", expectedDeviceId, deviceId)
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
