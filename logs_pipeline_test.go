package datadog

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogsPipelineGetAll(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response, err := ioutil.ReadFile("./tests/fixtures/logs/pipeline_response.json")
		if err != nil {
			t.Fatal(err)
		}
		w.Write(response)
	}))

	defer ts.Close()

	client := Client{
		baseUrl:    ts.URL,
		HttpClient: http.DefaultClient,
	}

	pipeline, err := client.GetLogsPipeline("dbJLomG9Tz-DYnAR5w-ilA")
	assert.Nil(t, err)
	assert.Equal(t, expectedPipeline, pipeline)
}

var expectedPipeline = &LogsPipeline{
	Id:         String("dbJLomG9Tz-DYnAR5w-ilA"),
	Type:       String("pipeline"),
	Name:       String("Test pipeline"),
	IsEnabled:  Bool(true),
	IsReadOnly: Bool(false),
	Filter: &FilterConfiguration{
		Query: String("source:test"),
	},
	Processors: []LogsProcessor{
		{
			Name:      String("nested pipeline"),
			IsEnabled: Bool(true),
		//	Definition: NestedPipeline{
		//		Type: String("pipeline"),
		//		Filter: &FilterConfiguration{
		//			Query: String("service:nest"),
		//		},
		//		Processors: []LogsProcessor{
		//			{
		//				Name:      String("test arithmetic processor"),
		//				IsEnabled: Bool(true),
		//				Definition: ArithmeticProcessor{
		//					Type:             String("arithmetic-processor"),
		//					Expression:       String("(time1-time2)*1000"),
		//					Target:           String("my_arithmetic"),
		//					IsReplaceMissing: Bool(false),
		//				},
		//			}, {
		//				Name:      String("test trace Id processor"),
		//				IsEnabled: Bool(true),
		//				Definition: SourceRemapper{
		//					Type:    String("trace-id-remapper"),
		//					Sources: []string{"dummy_trace_id1", "dummy_trace_id2"},
		//				},
		//			},
		//		},
		//	},
		//}, {
		//	Name:      String("test grok parser"),
		//	IsEnabled: Bool(true),
		//	Definition: GrokParser{
		//		Type:   String("grok-parser"),
		//		Source: String("text"),
		//		GrokRule: &GrokRule{
		//			SupportRules: String("date_parser %{date(\"yyyy-MM-dd HH:mm:ss,SSS\"):timestamp}"),
		//			MatchRules:   String("rule %{date(\"yyyy-MM-dd HH:mm:ss,SSS\"):timestamp}"),
		//		},
		//	},
		//}, {
		//	Name:      String("test remapper"),
		//	IsEnabled: Bool(true),
		//	Definition: AttributeRemapper{
		//		Type:               String("attribute-remapper"),
		//		Sources:            []string{"tag_1"},
		//		SourceType:         String("tag"),
		//		Target:             String("tag_3"),
		//		TargetType:         String("tag"),
		//		PreserveSource:     Bool(false),
		//		OverrideOnConflict: Bool(true),
		//	},
		//}, {
		//	Name:      String("test user-agent parser"),
		//	IsEnabled: Bool(true),
		//	Definition: UserAgentParser{
		//		Type:      String("user-agent-parser"),
		//		Sources:   []string{"user_agent"},
		//		Target:    String("my_agent.details"),
		//		IsEncoded: Bool(false),
		//	},
		//}, {
		//	Name:      String("test url parser"),
		//	IsEnabled: Bool(true),
		//	Definition: UrlParser{
		//		Type:                   String("url-parser"),
		//		Sources:                []string{"http_test"},
		//		Target:                 String("http_test.details"),
		//		NormalizeEndingSlashes: Bool(false),
		//	},
		//}, {
		//	Name:      String("test date remapper"),
		//	IsEnabled: Bool(true),
		//	Definition: SourceRemapper{
		//		Type:    String("date-remapper"),
		//		Sources: []string{"attribute_1", "attribute_2"},
		//	},
		//}, {
		//	Name:      String("test message remapper"),
		//	IsEnabled: Bool(true),
		//	Definition: SourceRemapper{
		//		Type:    String("message-remapper"),
		//		Sources: []string{"attribute_1", "attribute_2"},
		//	},
		//}, {
		//	Name:      String("test status remapper"),
		//	IsEnabled: Bool(true),
		//	Definition: SourceRemapper{
		//		Type:    String("status-remapper"),
		//		Sources: []string{"attribute_1", "attribute_2"},
		//	},
		//}, {
		//	Name:      String("test service remapper"),
		//	IsEnabled: Bool(true),
		//	Definition: SourceRemapper{
		//		Type:    String("service-remapper"),
		//		Sources: []string{"attribute_1", "attribute_2"},
		//	},
		//}, {
		//	Name:      String("test category processor"),
		//	IsEnabled: Bool(true),
		//	Definition: CategoryProcessor{
		//		Type:   String("category-processor"),
		//		Target: String("test_category"),
		//		Categories: []Category{
		//			{
		//				Name: String("5xx"),
		//				Filter: &FilterConfiguration{
		//					Query: String("status_code:[500 TO 599]"),
		//				},
		//			},
		//			{
		//				Name: String("4xx"),
		//				Filter: &FilterConfiguration{
		//					Query: String("status_code:[400 TO 499]"),
		//				},
		//			},
		//		},
		//	},
		},
	},
}
