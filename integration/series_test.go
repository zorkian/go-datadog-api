package integration

import (
	"testing"

	dd "github.com/zorkian/go-datadog-api"
)

func TestSeriesSubmit(t *testing.T) {
	metrics := []dd.Metric{{
		Metric: dd.String("test.metric"),
		Points: []dd.DataPoint{{dd.Float64(1.0), dd.Float64(2.0)}},
		Type:   dd.String("gauge"),
		Host:   dd.String("myhost"),
		Tags:   []string{"some:tag"},
		Unit:   dd.String("unit"),
	}}

	err := client.PostMetrics(metrics)
	if err != nil {
		t.Fatalf("Posting metrics failed when it shouldn't. (%s)", err)
	}
}
