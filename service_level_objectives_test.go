package datadog

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"sort"
	"testing"

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
				SLO:       Float64(99),
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

func testSLOGetMock(t *testing.T, fixturePath string) (*httptest.Server, *Client) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	return &ServiceLevelObjective{
		ID:          &id,
		Name:        sptr("Test SLO"),
		Description: sptr("Test Description"),
		Tags:        []string{"product:foo"},
		Thresholds: []*ServiceLevelObjectiveThreshold{
			{
				TimeFrame: String("7d"),
				SLO:       Float64(99),
				Warning:   Float64(99.5),
			},
			{
				TimeFrame: String("30d"),
				SLO:       Float64(98),
				Warning:   Float64(99),
			},
			{
				TimeFrame: String("90d"),
				SLO:       Float64(98),
				Warning:   Float64(99),
			},
		},
		Type:       &ServiceLevelObjectiveTypeMonitor,
		MonitorIDs: []int{1},
	}
}

func TestServiceLevelObjectiveIntegration(t *testing.T) {

	t.Run("CreateMonitor", func(t2 *testing.T) {
		ts, c := testSLOGetMock(t2, "./tests/fixtures/service_level_objectives/create_response.json")
		defer ts.Close()

		slo := getMockSLO("")
		created, err := c.CreateServiceLevelObjective(slo)
		assert.NoError(t2, err)
		assert.Equal(t2, "12345678901234567890123456789012", created.GetID())
	})

	t.Run("CreateMetric", func(t2 *testing.T) {
		ts, c := testSLOGetMock(t2, "./tests/fixtures/service_level_objectives/create_response_metric.json")
		defer ts.Close()

		slo := getMockSLO("")
		created, err := c.CreateServiceLevelObjective(slo)
		assert.NoError(t2, err)
		assert.Equal(t2, "abcdefabcdefabcdefabcdefabcdefab", created.GetID())
	})

	t.Run("Update", func(t2 *testing.T) {
		ts, c := testSLOGetMock(t2, "./tests/fixtures/service_level_objectives/update_response.json")
		defer ts.Close()

		slo := getMockSLO("12345678901234567890123456789012")
		slo, err := c.UpdateServiceLevelObjective(slo)
		assert.NoError(t2, err)
		assert.Equal(t2, "12345678901234567890123456789012", slo.GetID())
		assert.Equal(t2, 1563283900, slo.GetModifiedAt())
	})

	t.Run("Delete", func(t2 *testing.T) {
		ts, c := testSLOGetMock(t2, "./tests/fixtures/service_level_objectives/delete_response.json")
		defer ts.Close()

		slo := getMockSLO("12345678901234567890123456789012")
		err := c.DeleteServiceLevelObjective(slo.GetID())
		assert.NoError(t2, err)
	})

	t.Run("DeleteMany", func(t2 *testing.T) {
		ts, c := testSLOGetMock(t2, "./tests/fixtures/service_level_objectives/delete_many_response.json")
		defer ts.Close()

		err := c.DeleteServiceLevelObjectives(
			[]string{"12345678901234567890123456789012", "abcdefabcdefabcdefabcdefabcdefab"},
		)
		assert.NoError(t2, err)
	})

	t.Run("DeleteByTimeframe", func(t2 *testing.T) {
		ts, c := testSLOGetMock(t2, "./tests/fixtures/service_level_objectives/delete_by_timeframe_response.json")
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
		ts, c := testSLOGetMock(t2, "./tests/fixtures/service_level_objectives/get_by_id_response.json")
		defer ts.Close()

		slo, err := c.GetServiceLevelObjective("12345678901234567890123456789012")
		assert.NoError(t2, err)
		assert.Equal(t2, "12345678901234567890123456789012", slo.GetID())
	})

	t.Run("GetManyWithIDs", func(t2 *testing.T) {
		ts, c := testSLOGetMock(t2, "./tests/fixtures/service_level_objectives/get_many_response.json")
		defer ts.Close()

		slos, err := c.GetServiceLevelObjectives([]string{"12345678901234567890123456789012", "abcdefabcdefabcdefabcdefabcdefab"})
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

	t.Run("Search", func(t2 *testing.T) {
		ts, c := testSLOGetMock(t2, "./tests/fixtures/service_level_objectives/search_response.json")
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
				SLO:       Float64(99.9),
			},
			{
				TimeFrame: String("7d"),
				SLO:       Float64(98.9),
			},
			{
				TimeFrame: String("90d"),
				SLO:       Float64(97.9),
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
			SLO:       Float64(99.9),
		}
		threshold2 := &ServiceLevelObjectiveThreshold{
			TimeFrame: String("30d"),
			SLO:       Float64(99.9),
		}
		assert.True(t2, threshold1.Equal(threshold2))
		threshold3 := &ServiceLevelObjectiveThreshold{
			TimeFrame: String("30d"),
			SLO:       Float64(0.9),
		}

		assert.False(t2, threshold3.Equal(threshold2))

		threshold4 := &ServiceLevelObjectiveThreshold{
			TimeFrame: String("7d"),
			SLO:       Float64(99.9),
		}
		assert.False(t2, threshold2.Equal(threshold4))
	})

}
