package integration

import (
	"github.com/stretchr/testify/assert"
	"github.com/zorkian/go-datadog-api"
	"testing"
)

func TestLogsIndexGet(t *testing.T) {
	logsIndex, err := client.GetLogsIndex("main")
	assert.Nil(t, err)

	updateLogsIndex := datadog.LogsIndex{
		Filter: &datadog.FilterConfiguration{Query: datadog.String("updated query")},
		ExclusionFilters: []datadog.ExclusionFilter{
			{
				Name:      datadog.String("updated Filter 1"),
				IsEnabled: datadog.Bool(false),
				Filter: &datadog.Filter{
					Query:      datadog.String("source:agent"),
					SampleRate: datadog.Float64(0.3),
				},
			}, {
				Name:      datadog.String("updated Filter 2"),
				IsEnabled: datadog.Bool(false),
				Filter: &datadog.Filter{
					Query:      datadog.String("source:info"),
					SampleRate: datadog.Float64(0.2),
				},
			}, {
				Name:      datadog.String("updated Filter 3"),
				IsEnabled: datadog.Bool(false),
				Filter: &datadog.Filter{
					Query:      datadog.String("source:warn"),
					SampleRate: datadog.Float64(1.0),
				},
			},
		},
	}
	updatedLogsIndex, err := client.UpdateLogsIndex(*logsIndex.Name, &updateLogsIndex)
	defer assertRevertChange(t, logsIndex)
	assert.Nil(t, err)
	assert.Equal(t, &datadog.LogsIndex{
		Name:             logsIndex.Name,
		NumRetentionDays: logsIndex.NumRetentionDays,
		DailyLimit:       logsIndex.DailyLimit,
		IsRateLimited:    logsIndex.IsRateLimited,
		Filter:           updateLogsIndex.Filter,
		ExclusionFilters: updateLogsIndex.ExclusionFilters,
	}, updatedLogsIndex)

}

func TestUpdateIndexList(t *testing.T) {
	indexList, err := client.GetLogsIndexList()
	assert.Nil(t, err)
	size := len(indexList.IndexNames)
	updateList := make([]string, size)
	for i, name := range indexList.IndexNames {
		updateList[size-1-i] = name
	}
	updateIndexList := &datadog.LogsIndexList{IndexNames: updateList}
	updatedIndexList, err := client.UpdateLogsIndexList(updateIndexList)
	defer assertRevertOrder(t, indexList)
	assert.Nil(t, err)
	assert.Equal(t, updateIndexList, updatedIndexList)
}

func assertRevertOrder(t *testing.T, indexList *datadog.LogsIndexList) {
	revertedList, err := client.UpdateLogsIndexList(indexList)
	assert.Nil(t, err)
	assert.Equal(t, indexList, revertedList)
}

func assertRevertChange(t *testing.T, logsIndex *datadog.LogsIndex) {
	revertLogsIndex, err := client.UpdateLogsIndex(*logsIndex.Name, &datadog.LogsIndex{
		Filter:           logsIndex.Filter,
		ExclusionFilters: logsIndex.ExclusionFilters,
	})
	assert.Nil(t, err)
	assert.Equal(t, revertLogsIndex, logsIndex)
}
