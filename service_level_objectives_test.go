package datadog

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func sptr(i string) *string {
	return &i
}

func TestServiceLevelObjectiveSerialization(t *testing.T) {
	slo := ServiceLevelObjective{
		ID:          sptr("12345678901234567890123456789012"),
		Name:        sptr("Test SLO"),
		Description: sptr("Test Description"),
		Tags:        []string{"product:foo"},
		Thresholds: []*ServiceLevelObjectiveThreshold{
			{
				TimeFrame: String("7d"),
				Target:    Float64(99),
				Warning:   Float64(99.5),
			},
		},
		Type:       &ServiceLevelObjectiveTypeMonitor,
		MonitorIDs: []int{1},
	}

	raw, err := json.Marshal(&slo)
	assert.NoError(t, err)
	assert.NotEmpty(t, raw)

	var deserializedSLO ServiceLevelObjective

	err = json.Unmarshal(raw, &deserializedSLO)
	assert.NoError(t, err)
	assert.Equal(t, slo.ID, deserializedSLO.ID)
	assert.Equal(t, slo.Name, deserializedSLO.Name)
	assert.Equal(t, slo.Description, deserializedSLO.Description)
	assert.EqualValues(t, slo.Tags, deserializedSLO.Tags)
	assert.EqualValues(t, slo.Thresholds, deserializedSLO.Thresholds)
	assert.Equal(t, slo.Type, deserializedSLO.Type)
	assert.EqualValues(t, slo.MonitorIDs, deserializedSLO.MonitorIDs)
	assert.Nil(t, deserializedSLO.Groups)
}

const sloTestFixturePrefix = "./tests/fixtures/service_level_objectives/"

func testSLOGetMock(t *testing.T, expectedInputFixturePath, fixturePath string) (*httptest.Server, *Client) {

	if expectedInputFixturePath != "" {
		expectedInputFixturePath = sloTestFixturePrefix + expectedInputFixturePath
	}
	fixturePath = sloTestFixturePrefix + fixturePath

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if expectedInputFixturePath != "" && r != nil {
			body := r.Body
			if body == nil {
				t.Fatal("nil body received")
			}
			defer body.Close()
			payload, err := ioutil.ReadAll(body)
			if err != nil {
				t.Fatal(err)
			}

			var payloadContent interface{}
			err = json.Unmarshal(payload, &payloadContent)
			if err != nil {
				t.Fatal(err)
			}

			expectedPayload, err := ioutil.ReadFile(expectedInputFixturePath)
			if err != nil {
				t.Fatal(err)
			}
			var expectedPayloadContent interface{}
			err = json.Unmarshal(expectedPayload, &expectedPayloadContent)
			if err != nil {
				t.Fatal(err)
			}

			if !assert.Equal(t, expectedPayloadContent, payloadContent) {
				t.Fatalf("expected input did not match actual input")
			}
		}

		response, err := ioutil.ReadFile(fixturePath)
		if err != nil {
			t.Fatal(err)
		}
		w.Write(response)
	}))

	datadogClient := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}
	return ts, &datadogClient
}

func getMockSLO(id string) *ServiceLevelObjective {
	var sloID *string = nil
	if id != "" {
		sloID = &id
	}
	return &ServiceLevelObjective{
		ID:          sloID,
		Name:        sptr("Test SLO"),
		Description: sptr("test slo description"),
		Tags:        []string{"product:foo"},
		Thresholds: []*ServiceLevelObjectiveThreshold{
			{
				TimeFrame: String("7d"),
				Target:    Float64(99),
				Warning:   Float64(99.5),
			},
			{
				TimeFrame: String("30d"),
				Target:    Float64(98),
				Warning:   Float64(99),
			},
			{
				TimeFrame: String("90d"),
				Target:    Float64(98),
				Warning:   Float64(99),
			},
		},
	}
}

func TestServiceLevelObjectiveIntegration(t *testing.T) {

	t.Run("CreateMonitor", func(t2 *testing.T) {
		ts, c := testSLOGetMock(t2, "create_request_monitor.json", "create_response_monitor.json")
		defer ts.Close()

		slo := getMockSLO("")
		slo.SetType(ServiceLevelObjectiveTypeMonitor)
		slo.MonitorIDs = []int{1}
		created, err := c.CreateServiceLevelObjective(slo)
		assert.NoError(t2, err)
		assert.Equal(t2, "12345678901234567890123456789012", created.GetID())
	})

	t.Run("CreateMetric", func(t2 *testing.T) {
		ts, c := testSLOGetMock(t2, "create_request_metric.json", "create_response_metric.json")
		defer ts.Close()

		slo := getMockSLO("")
		slo.SetType(ServiceLevelObjectiveTypeMetric)
		slo.SetQuery(ServiceLevelObjectiveMetricQuery{
			Numerator:   String("sum:my.metric{type:good}.as_count()"),
			Denominator: String("sum:my.metric{*}.as_count()"),
		})
		created, err := c.CreateServiceLevelObjective(slo)
		assert.NoError(t2, err)
		assert.Equal(t2, "abcdefabcdefabcdefabcdefabcdefab", created.GetID())
	})

	t.Run("Update", func(t2 *testing.T) {
		ts, c := testSLOGetMock(t2, "update_request.json", "update_response.json")
		defer ts.Close()

		slo := getMockSLO("12345678901234567890123456789012")
		slo.SetType(ServiceLevelObjectiveTypeMonitor)
		slo.MonitorIDs = []int{1}
		slo, err := c.UpdateServiceLevelObjective(slo)
		assert.NoError(t2, err)
		assert.Equal(t2, "12345678901234567890123456789012", slo.GetID())
		assert.Equal(t2, 1563283900, slo.GetModifiedAt())
	})

	t.Run("Delete", func(t2 *testing.T) {
		ts, c := testSLOGetMock(t2, "", "delete_response.json")
		defer ts.Close()

		slo := getMockSLO("12345678901234567890123456789012")
		err := c.DeleteServiceLevelObjective(slo.GetID())
		assert.NoError(t2, err)
	})

	t.Run("DeleteMany", func(t2 *testing.T) {
		ts, c := testSLOGetMock(t2, "delete_many_request.json", "delete_many_response.json")
		defer ts.Close()

		err := c.DeleteServiceLevelObjectives(
			[]string{"12345678901234567890123456789012", "abcdefabcdefabcdefabcdefabcdefab"},
		)
		assert.NoError(t2, err)
	})

	t.Run("DeleteByTimeframe", func(t2 *testing.T) {
		ts, c := testSLOGetMock(t2, "delete_by_timeframe_request.json", "delete_by_timeframe_response.json")
		defer ts.Close()

		/* Some Context for this test case:  This is useful for doing individual time-frame deletes across different SLOs (used by the web list view bulk delete)

		`12345678901234567890123456789012` was defined with 2 time frames: "7d" and "30d"
		`abcdefabcdefabcdefabcdefabcdefab` was defined with 2 time frames: "30d" and "90d"

		When we delete `7d` from `12345678901234567890123456789012` we still have `30d` timeframe remaining, hence this is "updated"
		When we delete `30d` and `90d` from `abcdefabcdefabcdefabcdefabcdefab` we are left with 0 time frames, hence this is "deleted"
		     and the entire SLO config is deleted
		*/
		resp, err := c.DeleteServiceLevelObjectiveTimeFrames(map[string][]string{
			"12345678901234567890123456789012": {"7d"},
			"abcdefabcdefabcdefabcdefabcdefab": {"30d", "90d"},
		})
		assert.NoError(t2, err)
		assert.EqualValues(t2, resp.UpdatedIDs, []string{"12345678901234567890123456789012"})
		assert.EqualValues(t2, resp.DeletedIDs, []string{"abcdefabcdefabcdefabcdefabcdefab"})
	})

	t.Run("GetByID", func(t2 *testing.T) {
		ts, c := testSLOGetMock(t2, "", "get_by_id_response.json")
		defer ts.Close()

		slo, err := c.GetServiceLevelObjective("12345678901234567890123456789012")
		assert.NoError(t2, err)
		assert.Equal(t2, "12345678901234567890123456789012", slo.GetID())
	})

	t.Run("SearchWithIDs", func(t2 *testing.T) {
		ts, c := testSLOGetMock(t2, "", "get_many_response.json")
		defer ts.Close()

		slos, err := c.SearchServiceLevelObjectives(1000, 0, "", []string{"12345678901234567890123456789012", "abcdefabcdefabcdefabcdefabcdefab"})
		assert.NoError(t2, err)
		assert.Len(t2, slos, 2)

		contains := func(slos []*ServiceLevelObjective, id string) bool {
			for _, slo := range slos {
				if slo.GetID() == id {
					return true
				}
			}
			return false
		}
		assert.True(t2, contains(slos, "12345678901234567890123456789012"))
		assert.True(t2, contains(slos, "abcdefabcdefabcdefabcdefabcdefab"))
	})

	t.Run("SearchWithQuery", func(t2 *testing.T) {
		ts, c := testSLOGetMock(t2, "", "search_response.json")
		defer ts.Close()

		slos, err := c.SearchServiceLevelObjectives(1000, 0, "service:foo AND team:a", nil)
		assert.NoError(t2, err)
		assert.Len(t2, slos, 1)
		assert.Equal(t2, "12345678901234567890123456789012", slos[0].GetID())
	})

	t.Run("thresholds are sortable by duration", func(t2 *testing.T) {
		thresholds := ServiceLevelObjectiveThresholds{
			{
				TimeFrame: String("30d"),
				Target:    Float64(99.9),
			},
			{
				TimeFrame: String("7d"),
				Target:    Float64(98.9),
			},
			{
				TimeFrame: String("90d"),
				Target:    Float64(97.9),
			},
		}

		sort.Sort(thresholds)
		assert.Equal(t2, "7d", thresholds[0].GetTimeFrame())
		assert.Equal(t2, "30d", thresholds[1].GetTimeFrame())
		assert.Equal(t2, "90d", thresholds[2].GetTimeFrame())
	})

	t.Run("thresholds are comparable", func(t2 *testing.T) {
		threshold1 := &ServiceLevelObjectiveThreshold{
			TimeFrame: String("30d"),
			Target:    Float64(99.9),
		}
		threshold2 := &ServiceLevelObjectiveThreshold{
			TimeFrame: String("30d"),
			Target:    Float64(99.9),
		}
		assert.True(t2, threshold1.Equal(threshold2))
		threshold3 := &ServiceLevelObjectiveThreshold{
			TimeFrame: String("30d"),
			Target:    Float64(0.9),
		}

		assert.False(t2, threshold3.Equal(threshold2))

		threshold4 := &ServiceLevelObjectiveThreshold{
			TimeFrame: String("7d"),
			Target:    Float64(99.9),
		}
		assert.False(t2, threshold2.Equal(threshold4))
	})

	t.Run("CheckCanDelete", func(t2 *testing.T) {
		ts, c := testSLOGetMock(t2, "", "check_can_delete_response.json")
		defer ts.Close()

		resp, err := c.CheckCanDeleteServiceLevelObjectives(
			[]string{"12345678901234567890123456789012", "abcdefabcdefabcdefabcdefabcdefab"},
		)
		assert.NoError(t2, err)
		assert.EqualValues(t2, []string{"12345678901234567890123456789012"}, resp.Data.OK)
		assert.EqualValues(t2,
			map[string]string{
				"abcdefabcdefabcdefabcdefabcdefab": "SLO abcdefabcdefabcdefabcdefabcdefab is used in dashboard 123-456-789",
			},
			resp.Errors,
		)
	})

	t.Run("GetServiceLevelObjectiveHistory-Metric", func(t2 *testing.T) {
		ts, c := testSLOGetMock(t2, "", "get_history_metric_response.json")
		defer ts.Close()

		resp, err := c.GetServiceLevelObjectiveHistory(
			"12345678901234567890123456789012",
			time.Unix(1571162100, 0),
			time.Unix(1571766900, 0),
		)
		assert.NoError(t2, err)
		assert.Nil(t2, resp.Error)
		assert.Equal(t2, float32(100), resp.Data.Overall.Uptime)
		assert.Equal(t2, json.Number("3698988"), resp.Data.Metrics.Numerator.Sum)
		assert.Equal(t2, json.Number("3698988"), resp.Data.Metrics.Denominator.Sum)
	})

	t.Run("GetServiceLevelObjectiveHistory-Monitor", func(t2 *testing.T) {
		ts, c := testSLOGetMock(t2, "", "get_history_monitor_response.json")
		defer ts.Close()

		resp, err := c.GetServiceLevelObjectiveHistory(
			"12345678901234567890123456789012",
			time.Unix(1571162100, 0),
			time.Unix(1571766900, 0),
		)
		assert.NoError(t2, err)
		assert.Nil(t2, resp.Error)
		assert.Equal(t2, float32(6.765872955322266), resp.Data.Overall.Uptime)
		assert.Len(t2, resp.Data.Groups, 1)
		assert.Equal(t2, float32(6.765872955322266), resp.Data.Groups[0].Uptime)
		assert.Equal(t2, "some:tag", resp.Data.Groups[0].Name)
	})

}
