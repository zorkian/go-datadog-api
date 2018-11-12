package datadog

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type Alias Yaxis

type YAxisTestSuite struct {
	suite.Suite
	yJSON                     []byte
	yMarshalledJSON           []byte
	y                         Yaxis
	yAutoMinMaxJSON           []byte
	yAutoMinMaxMarshalledJSON []byte
	yAutoMinMax               Yaxis
}

func (suite *YAxisTestSuite) SetupTest() {
	// Custom Y.Min, Y.Max
	suite.yJSON = []byte(`{"min":0,"max":1,"scale":"linear","includeZero":true,"units":true}`)
	suite.yMarshalledJSON = suite.yJSON
	yMinFloat := float64(0)
	yMaxFloat := float64(1)
	yScale := "linear"
	yIncludeZero := true
	yIncludeUnits := true
	suite.y = Yaxis{
		Min:          &yMinFloat,
		AutoMin:      false,
		Max:          &yMaxFloat,
		AutoMax:      false,
		Scale:        &yScale,
		IncludeZero:  &yIncludeZero,
		IncludeUnits: &yIncludeUnits,
	}
	// Auto Y.Min, Y.Max
	suite.yAutoMinMaxJSON = []byte(`{"min":"auto","max":"auto","scale":"linear","includeZero":true,"units":true}`)
	suite.yAutoMinMaxMarshalledJSON = []byte(`{"scale":"linear","includeZero":true,"units":true}`)
	suite.yAutoMinMax = Yaxis{
		Min:          nil,
		AutoMin:      true,
		Max:          nil,
		AutoMax:      true,
		Scale:        &yScale,
		IncludeZero:  &yIncludeZero,
		IncludeUnits: &yIncludeUnits,
	}
}

func (suite *YAxisTestSuite) TestYAxisJSONMarshalCustomMinMax() {
	jsonData, err := json.Marshal(suite.y)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.yMarshalledJSON, jsonData)
}

func (suite *YAxisTestSuite) TestYAxisJSONUnmarshalCustomMinMax() {
	var res Yaxis
	err := json.Unmarshal(suite.yJSON, &res)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.y, res)
}

func (suite *YAxisTestSuite) TestYAxisJSONMarshalAutoMinMax() {
	jsonData, err := json.Marshal(suite.yAutoMinMax)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.yAutoMinMaxMarshalledJSON, jsonData)
}

func (suite *YAxisTestSuite) TestYAxisJSONUnmarshalAutoMinMax() {
	var res Yaxis
	err := json.Unmarshal(suite.yAutoMinMaxJSON, &res)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), suite.yAutoMinMax, res)
}

func TestYAxisTestSuite(t *testing.T) {
	suite.Run(t, new(YAxisTestSuite))
}

func TestDashboardGetters(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/dashboards_response.json")
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

	dashboards, err := datadogClient.GetDashboards()
	if err != nil {
		t.Fatal(err)
	}

	if len(dashboards) != 2 {
		t.Fatalf("expect response with 2 dashboards. Got %d", len(dashboards))
	}

	d1 := dashboards[0]

	expectedID := 123
	if id := d1.GetId(); id != expectedID {
		t.Fatalf("expect ID %d. Got %d", expectedID, id)
	}

	expectedTitle := "Dashboard 1"
	if title := d1.GetTitle(); title != expectedTitle {
		t.Fatalf("expect title %s. Got %s", expectedTitle, title)
	}

	expectedDescription := "created by user1"
	if description := d1.GetDescription(); description != expectedDescription {
		t.Fatalf("expect description %s. Got %s", expectedDescription, description)
	}

	expectedCreateTime := "2018-10-05T12:32:01.000000+00:00"
	createTime, ok := d1.GetCreatedOk()
	if !ok {
		t.Fatalf("expect to have a created field")
	}

	if createTime != expectedCreateTime {
		t.Fatalf("expect create time %s. Got %s", expectedCreateTime, createTime)
	}

	expectedReadOnly := false
	readOnly, ok := d1.GetReadOnlyOk()
	if !ok {
		t.Fatalf("expect to have a read_only field")
	}

	if readOnly != expectedReadOnly {
		t.Fatalf("expect read_only %v. Got %v", expectedReadOnly, readOnly)
	}

	expectedModified := "2018-09-11T06:38:09.000000+00:00"
	modified, ok := d1.GetModifiedOk()
	if !ok {
		t.Fatalf("expect to have a modified field")
	}

	if modified != expectedModified {
		t.Fatalf("expect modified %s. Got %s", expectedModified, modified)
	}

	createdBy, ok := d1.GetCreatedByOk()
	if !ok {
		t.Fatal("expect created_by field to exist")
	}

	validateCreatedBy(t, &createdBy)
}

func validateCreatedBy(t *testing.T, cb *CreatedBy) {
	expectedAccessRole := "adm"
	accessRole, ok := cb.GetAccessRoleOk()
	if !ok {
		t.Fatal("expect to have access_role field")
	}

	if accessRole != expectedAccessRole {
		t.Fatalf("expect access_role %s. Got %s", expectedAccessRole, accessRole)
	}

	expectedDisabled := true
	disabled, ok := cb.GetDisabledOk()
	if !ok {
		t.Fatal("expect to have disabled field")
	}

	if disabled != expectedDisabled {
		t.Fatalf("expect disabled %v. Got %v", expectedDisabled, disabled)
	}

	expectedEmail := "john.doe@company.com"
	email, ok := cb.GetEmailOk()
	if !ok {
		t.Fatal("expect to have email field")
	}

	if email != expectedEmail {
		t.Fatalf("expect email %s. Got %s", expectedEmail, email)
	}

	expectedHandle := "john.doe@company.com"
	handle, ok := cb.GetHandleOk()
	if !ok {
		t.Fatal("expect to have handle field")
	}

	if handle != expectedHandle {
		t.Fatalf("expect handle %s. Got %s", expectedHandle, handle)
	}

	expectedIcon := "https://pics.site.com/123.jpg"
	icon, ok := cb.GetIconOk()
	if !ok {
		t.Fatal("expect to have icon field")
	}

	if icon != expectedIcon {
		t.Fatalf("expect icon %s. Got %s", expectedIcon, icon)
	}

	expectedName := "John Doe"
	name, ok := cb.GetNameOk()
	if !ok {
		t.Fatal("expect to have name field")
	}

	if name != expectedName {
		t.Fatalf("expect name %s. Got %s", expectedName, name)
	}

	expectedVerified := true
	verified, ok := cb.GetVerifiedOk()
	if !ok {
		t.Fatal("expect to have verified field")
	}

	if verified != expectedVerified {
		t.Fatalf("expect verify %v. Got %v", expectedVerified, verified)
	}
}
