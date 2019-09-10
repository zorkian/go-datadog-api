package integration

import (
	"github.com/stretchr/testify/assert"
	"github.com/zorkian/go-datadog-api"
	"log"
	"testing"
)

func TestLogsPipelineCrud(t *testing.T) {
	createdPipeline, err := client.CreateLogsPipeline(
		&datadog.LogsPipeline{
			Name:      datadog.String("Test pipeline"),
			IsEnabled: datadog.Bool(true),
			Filter: &datadog.FilterConfiguration{
				Query: datadog.String("source:test"),
			},
			Processors: []datadog.LogsProcessor{
				{
					Name:      datadog.String("nested pipeline"),
					IsEnabled: datadog.Bool(true),
					Definition: datadog.NestedPipeline{
						Type: datadog.String("pipeline"),
						Filter: &datadog.FilterConfiguration{
							Query: datadog.String("service:nest"),
						},
						Processors: []datadog.LogsProcessor{
							{
								Name:      datadog.String("test arithmetic processor"),
								IsEnabled: datadog.Bool(true),
								Definition: datadog.ArithmeticProcessor{
									Type:             datadog.String("arithmetic-processor"),
									Expression:       datadog.String("(time1-time2)*1000"),
									Target:           datadog.String("my_arithmetic"),
									IsReplaceMissing: datadog.Bool(false),
								},
							}, {
								Name:      datadog.String("test trace Id processor"),
								IsEnabled: datadog.Bool(true),
								Definition: datadog.SourceRemapper{
									Type:    datadog.String("trace-id-remapper"),
									Sources: []string{"dummy_trace_id1", "dummy_trace_id2"},
								},
							},
						},
					},
				}, {
					Name:      datadog.String("test grok parser"),
					IsEnabled: datadog.Bool(true),
					Definition: datadog.GrokParser{
						Type:   datadog.String("grok-parser"),
						Source: datadog.String("text"),
						GrokRule: &datadog.GrokRule{
							SupportRules: datadog.String("date_parser %{date(\"yyyy-MM-dd HH:mm:ss,SSS\"):timestamp}"),
							MatchRules:   datadog.String("rule %{date(\"yyyy-MM-dd HH:mm:ss,SSS\"):timestamp}"),
						},
					},
				}, {
					Name:      datadog.String("test remapper"),
					IsEnabled: datadog.Bool(true),
					Definition: datadog.AttributeRemapper{
						Type:               datadog.String("attribute-remapper"),
						Sources:            []string{"tag_1"},
						SourceType:         datadog.String("tag"),
						Target:             datadog.String("tag_3"),
						TargetType:         datadog.String("tag"),
						PreserveSource:     datadog.Bool(false),
						OverrideOnConflict: datadog.Bool(true),
					},
				}, {
					Name:      datadog.String("test user-agent parser"),
					IsEnabled: datadog.Bool(true),
					Definition: datadog.UserAgentParser{
						Type:      datadog.String("user-agent-parser"),
						Sources:   []string{"user_agent"},
						Target:    datadog.String("my_agent.details"),
						IsEncoded: datadog.Bool(false),
					},
				}, {
					Name:      datadog.String("test url parser"),
					IsEnabled: datadog.Bool(true),
					Definition: datadog.UrlParser{
						Type:                   datadog.String("url-parser"),
						Sources:                []string{"http_test"},
						Target:                 datadog.String("http_test.details"),
						NormalizeEndingSlashes: datadog.Bool(false),
					},
				}, {
					Name:      datadog.String("test date remapper"),
					IsEnabled: datadog.Bool(true),
					Definition: datadog.SourceRemapper{
						Type:    datadog.String("date-remapper"),
						Sources: []string{"attribute_1", "attribute_2"},
					},
				}, {
					Name:      datadog.String("test message remapper"),
					IsEnabled: datadog.Bool(true),
					Definition: datadog.SourceRemapper{
						Type:    datadog.String("message-remapper"),
						Sources: []string{"attribute_1", "attribute_2"},
					},
				}, {
					Name:      datadog.String("test status remapper"),
					IsEnabled: datadog.Bool(true),
					Definition: datadog.SourceRemapper{
						Type:    datadog.String("status-remapper"),
						Sources: []string{"attribute_1", "attribute_2"},
					},
				}, {
					Name:      datadog.String("test service remapper"),
					IsEnabled: datadog.Bool(true),
					Definition: datadog.SourceRemapper{
						Type:    datadog.String("service-remapper"),
						Sources: []string{"attribute_1", "attribute_2"},
					},
				}, {
					Name:      datadog.String("test category processor"),
					IsEnabled: datadog.Bool(true),
					Definition: datadog.CategoryProcessor{
						Type:   datadog.String("category-processor"),
						Target: datadog.String("test_category"),
						Categories: []datadog.Category{
							{
								Name: datadog.String("5xx"),
								Filter: &datadog.FilterConfiguration{
									Query: datadog.String("status_code:[500 TO 599]"),
								},
							},
							{
								Name: datadog.String("4xx"),
								Filter: &datadog.FilterConfiguration{
									Query: datadog.String("status_code:[400 TO 499]"),
								},
							},
						},
					},
				},
			},
		})

	assert.Nil(t, err)

	defer assertPipelineDelete(t, *createdPipeline.Id)

	if err != nil {
		log.Fatalf("fatal: %s\n", err)
	}
	pipeline, err := client.GetLogsPipeline(*createdPipeline.Id)

	assert.Equal(t, createdPipeline, pipeline)
	assert.Nil(t, err)
	updatedPipeline, err := client.UpdateLogsPipeline(
		*pipeline.Id,
		&datadog.LogsPipeline{
			Name:      datadog.String("updated pipeline"),
			IsEnabled: datadog.Bool(false),
			Filter: &datadog.FilterConfiguration{
				Query: datadog.String("source:kafka"),
			},
		})
	assert.Nil(t, err)
	pipeline, err = client.GetLogsPipeline(*updatedPipeline.Id)
	assert.Equal(t, updatedPipeline, pipeline)
	assert.Nil(t, err)
}

func assertPipelineDelete(t *testing.T, id string) {
	err := client.DeleteLogsPipeline(id)
	assert.Nil(t, err)
	pipeline, err := client.GetLogsPipeline(id)
	assert.Nil(t, pipeline)
}
