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
					Type:      datadog.String("pipeline"),
					Definition: datadog.NestedPipeline{
						Filter: &datadog.FilterConfiguration{
							Query: datadog.String("service:nest"),
						},
						Processors: []datadog.LogsProcessor{
							{
								Name:      datadog.String("test arithmetic processor"),
								IsEnabled: datadog.Bool(true),
								Type:      datadog.String("arithmetic-processor"),
								Definition: datadog.ArithmeticProcessor{
									Expression:       datadog.String("(time1-time2)*1000"),
									Target:           datadog.String("my_arithmetic"),
									IsReplaceMissing: datadog.Bool(false),
								},
							}, {
								Name:      datadog.String("test trace Id processor"),
								IsEnabled: datadog.Bool(true),
								Type:      datadog.String("trace-id-remapper"),
								Definition: datadog.SourceRemapper{
									Sources: []string{"dummy_trace_id1", "dummy_trace_id2"},
								},
							},
						},
					},
				}, {
					Name:      datadog.String("test grok parser"),
					IsEnabled: datadog.Bool(true),
					Type:      datadog.String("grok-parser"),
					Definition: datadog.GrokParser{
						Source: datadog.String("text"),
						GrokRule: &datadog.GrokRule{
							SupportRules: datadog.String("date_parser %{date(\"yyyy-MM-dd HH:mm:ss,SSS\"):timestamp}"),
							MatchRules:   datadog.String("rule %{date(\"yyyy-MM-dd HH:mm:ss,SSS\"):timestamp}"),
						},
					},
				}, {
					Name:      datadog.String("test remapper"),
					IsEnabled: datadog.Bool(true),
					Type:      datadog.String("attribute-remapper"),
					Definition: datadog.AttributeRemapper{
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
					Type:      datadog.String("user-agent-parser"),
					Definition: datadog.UserAgentParser{
						Sources:   []string{"user_agent"},
						Target:    datadog.String("my_agent.details"),
						IsEncoded: datadog.Bool(false),
					},
				}, {
					Name:      datadog.String("test url parser"),
					IsEnabled: datadog.Bool(true),
					Type:      datadog.String("url-parser"),
					Definition: datadog.UrlParser{
						Sources:                []string{"http_test"},
						Target:                 datadog.String("http_test.details"),
						NormalizeEndingSlashes: datadog.Bool(false),
					},
				}, {
					Name:      datadog.String("test date remapper"),
					IsEnabled: datadog.Bool(true),
					Type:      datadog.String("date-remapper"),
					Definition: datadog.SourceRemapper{
						Sources: []string{"attribute_1", "attribute_2"},
					},
				}, {
					Name:      datadog.String("test message remapper"),
					IsEnabled: datadog.Bool(true),
					Type:      datadog.String("message-remapper"),
					Definition: datadog.SourceRemapper{
						Sources: []string{"attribute_1", "attribute_2"},
					},
				}, {
					Name:      datadog.String("test status remapper"),
					IsEnabled: datadog.Bool(true),
					Type:      datadog.String("status-remapper"),
					Definition: datadog.SourceRemapper{
						Sources: []string{"attribute_1", "attribute_2"},
					},
				}, {
					Name:      datadog.String("test service remapper"),
					IsEnabled: datadog.Bool(true),
					Type:      datadog.String("service-remapper"),
					Definition: datadog.SourceRemapper{
						Sources: []string{"attribute_1", "attribute_2"},
					},
				}, {
					Name:      datadog.String("test category processor"),
					IsEnabled: datadog.Bool(true),
					Type:      datadog.String("category-processor"),
					Definition: datadog.CategoryProcessor{
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
				}, {
					Name:      datadog.String("test string builder processor"),
					IsEnabled: datadog.Bool(true),
					Type:      datadog.String("string-builder-processor"),
					Definition: datadog.StringBuilderProcessor{
						Template: datadog.String("hello %{user.name}"),
						IsReplaceMissing: datadog.Bool(false),
						Target: datadog.String("target"),
					},
				}, {
				Name: datadog.String("geo ip parser test"),
				IsEnabled: datadog.Bool(false),
				Type: datadog.String("geo-ip-parser"),
				Definition: datadog.GeoIPParser{
				Sources: []string{"source1", "source2"},
				Target: datadog.String("target"),
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
