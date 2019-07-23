package integration

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zorkian/go-datadog-api"
)

func TestServiceLevelObjectivesCreateGetUpdateAndDelete(t *testing.T) {
	expected := &datadog.ServiceLevelObjective{
		Name:        datadog.String("Integration Test SLO - 'Test Create, Update and Delete'"),
		Description: datadog.String("Integration test for SLOs"),
		Tags:        []string{"test:integration"},
		Thresholds: datadog.ServiceLevelObjectiveThresholds{
			{
				TimeFrame: datadog.String("7d"),
				Target:    datadog.Float64(99),
				Warning:   datadog.Float64(99.5),
			},
		},
		Type: &datadog.ServiceLevelObjectiveTypeMetric,
		Query: &datadog.ServiceLevelObjectiveMetricQuery{
			Numerator:   datadog.String("sum:my.metric{type:good}.as_count()"),
			Denominator: datadog.String("sum:my.metric{*}.as_count()"),
		},
	}

	// Create
	actual, err := client.CreateServiceLevelObjective(expected)
	assert.NoError(t, err)
	assert.NotEmpty(t, actual.GetID())
	assert.Equal(t, expected.Name, actual.Name)
	assert.Equal(t, expected.Description, actual.Description)
	assert.True(t, expected.Thresholds.Equal(actual.Thresholds))

	// Get
	found, err := client.GetServiceLevelObjective(actual.GetID())
	assert.NoError(t, err)
	assert.Equal(t, actual.GetID(), found.GetID())

	// Update
	actual.SetDescription("Integration test for SLOs - updated")
	actual.Thresholds = datadog.ServiceLevelObjectiveThresholds{
		{
			TimeFrame: datadog.String("7d"),
			Target:    datadog.Float64(99),
			Warning:   datadog.Float64(99.5),
		},
		{
			TimeFrame: datadog.String("30d"),
			Target:    datadog.Float64(99),
			Warning:   datadog.Float64(99.5),
		},
	}
	actual, err = client.UpdateServiceLevelObjective(actual)
	assert.NoError(t, err)
	assert.Equal(t, "Integration test for SLOs - updated", actual.GetDescription())
	assert.Len(t, actual.Thresholds, 2)

	// Delete
	err = client.DeleteServiceLevelObjective(actual.GetID())
	assert.NoError(t, err)
}

func TestServiceLevelObjectivesBulkTimeFrameDelete(t *testing.T) {
	expected1 := &datadog.ServiceLevelObjective{
		Name:        datadog.String("Integration Test SLO - 'Test Multi Time Frame Delete 1'"),
		Description: datadog.String("Integration test for SLOs"),
		Tags:        []string{"test:integration"},
		Thresholds: datadog.ServiceLevelObjectiveThresholds{
			{
				TimeFrame: datadog.String("7d"),
				Target:    datadog.Float64(99),
				Warning:   datadog.Float64(99.5),
			},
			{
				TimeFrame: datadog.String("30d"),
				Target:    datadog.Float64(99),
				Warning:   datadog.Float64(99.5),
			},
			{
				TimeFrame: datadog.String("90d"),
				Target:    datadog.Float64(99),
				Warning:   datadog.Float64(99.5),
			},
		},
		Type: &datadog.ServiceLevelObjectiveTypeMetric,
		Query: &datadog.ServiceLevelObjectiveMetricQuery{
			Numerator:   datadog.String("sum:my.metric{type:good}.as_count()"),
			Denominator: datadog.String("sum:my.metric{*}.as_count()"),
		},
	}
	expected2 := &datadog.ServiceLevelObjective{
		Name:        datadog.String("Integration Test SLO - 'Test Multi Time Frame Delete 2'"),
		Description: datadog.String("Integration test for SLOs"),
		Tags:        []string{"test:integration"},
		Thresholds: datadog.ServiceLevelObjectiveThresholds{
			{
				TimeFrame: datadog.String("7d"),
				Target:    datadog.Float64(99),
				Warning:   datadog.Float64(99.5),
			},
		},
		Type: &datadog.ServiceLevelObjectiveTypeMetric,
		Query: &datadog.ServiceLevelObjectiveMetricQuery{
			Numerator:   datadog.String("sum:my.metric{type:good}.as_count()"),
			Denominator: datadog.String("sum:my.metric{*}.as_count()"),
		},
	}

	// Create
	actual1, err := client.CreateServiceLevelObjective(expected1)
	assert.NoError(t, err)
	assert.NotEmpty(t, actual1.GetID())
	actual2, err := client.CreateServiceLevelObjective(expected2)
	assert.NoError(t, err)
	assert.NotEmpty(t, actual2.GetID())

	// Do multi-timeframe delete
	timeframesToDelete := map[string][]string{
		// delete only 2 of 3 timeframes from 1
		actual1.GetID(): {
			"30d", "90d",
		},
		// delete all timeframes from 2
		actual2.GetID(): {
			"7d",
		},
	}

	resp, err := client.DeleteServiceLevelObjectiveTimeFrames(timeframesToDelete)
	assert.EqualValues(t, []string{actual2.GetID()}, resp.DeletedIDs)
	assert.EqualValues(t, []string{actual1.GetID()}, resp.UpdatedIDs)

	// Delete
	err = client.DeleteServiceLevelObjective(actual1.GetID())
	assert.NoError(t, err)
}
