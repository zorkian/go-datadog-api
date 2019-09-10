package datadog

import (
	"encoding/json"
	"fmt"
)

const (
	ARITHMETIC_PROCESSOR        = "arithmetic-processor"
	ATTIBUTE_REMAPPER_PROCESSOR = "attribute-remapper"
	CATEGORY_PROCESSOR          = "category-processor"
	DATE_REMAPPER_PROCESSOR     = "date-remapper"
	GROK_PARSER_PROCESSOR       = "grok-parser"
	MESSAGE_REMAPPER_PROCESSOR  = "message-remapper"
	NESTED_PIPELINE_PROCESSOR   = "pipeline"
	SERVICE_REMAPPER_PROCESSOR  = "service-remapper"
	STATUS_REMAPPER_PROCESSOR   = "status-remapper"
	TRACE_ID_REMAPPER_PROCESSOR = "trace-id-remapper"
	URL_PARSER_PROCESSOR        = "url-parser"
	USER_AGENT_PARSER_PROCESSOR = "user-agent-parser"
)

type LogsProcessor struct {
	Name       *string     `json:"name"`
	IsEnabled  *bool       `json:"is_enabled"`
	Definition interface{} `json:"definition"`
}

type ArithmeticProcessor struct {
	Type             *string `json:"type"`
	Expression       *string `json:"expression"`
	Target           *string `json:"target"`
	IsReplaceMissing *bool   `json:"is_replace_missing"`
}

type AttributeRemapper struct {
	Type               *string  `json:"type"`
	Sources            []string `json:"sources"`
	SourceType         *string  `json:"source_type"`
	Target             *string  `json:"target"`
	TargetType         *string  `json:"target_type"`
	PreserveSource     *bool    `json:"preserve_source"`
	OverrideOnConflict *bool    `json:"override_on_conflict"`
}

type CategoryProcessor struct {
	Type       *string    `json:"type"`
	Target     *string    `json:"target"`
	Categories []Category `json:"categories"`
}

type Category struct {
	Name   *string              `json:"name"`
	Filter *FilterConfiguration `json:"filter"`
}

type SourceRemapper struct {
	Type    *string  `json:"type"`
	Sources []string `json:"sources"`
}

type GrokParser struct {
	Type     *string   `json:"type"`
	Source   *string   `json:"source"`
	GrokRule *GrokRule `json:"grok"`
}

type GrokRule struct {
	SupportRules *string `json:"support_rules"`
	MatchRules   *string `json:"match_rules"`
}

type NestedPipeline struct {
	Type       *string              `json:"type"`
	Filter     *FilterConfiguration `json:"filter"`
	Processors []LogsProcessor      `json:"processors,omitempty"`
}

type UrlParser struct {
	Type                   *string  `json:"type"`
	Sources                []string `json:"sources"`
	Target                 *string  `json:"target"`
	NormalizeEndingSlashes *bool    `json:"normalize_ending_slashes"`
}

type UserAgentParser struct {
	Type      *string  `json:"type"`
	Sources   []string `json:"sources"`
	Target    *string  `json:"target"`
	IsEncoded *bool    `json:"is_encoded"`
}

func (processor *LogsProcessor) UnmarshalJSON(data []byte) error {
	var processorHandler struct {
		Type      *string `json:"type"`
		Name      *string `json:"name"`
		IsEnabled *bool   `json:"is_enabled"`
	}
	if err := json.Unmarshal(data, &processorHandler); err != nil {
		return err
	}

	processor.Name = processorHandler.Name
	processor.IsEnabled = processorHandler.IsEnabled

	switch *processorHandler.Type {
	case ARITHMETIC_PROCESSOR:
		var arithmeticProcessor ArithmeticProcessor
		if err := json.Unmarshal(data, &arithmeticProcessor); err != nil {
			return err
		}
		processor.Definition = arithmeticProcessor
	case ATTIBUTE_REMAPPER_PROCESSOR:
		var attributeRemapper AttributeRemapper
		if err := json.Unmarshal(data, &attributeRemapper); err != nil {
			return err
		}
		processor.Definition = attributeRemapper
	case CATEGORY_PROCESSOR:
		var categoryProcessor CategoryProcessor
		if err := json.Unmarshal(data, &categoryProcessor); err != nil {
			return err
		}
		processor.Definition = categoryProcessor
	case DATE_REMAPPER_PROCESSOR:
		var dateRemapper SourceRemapper
		if err := json.Unmarshal(data, &dateRemapper); err != nil {
			return err
		}
		processor.Definition = dateRemapper
	case GROK_PARSER_PROCESSOR:
		var grokParser GrokParser
		if err := json.Unmarshal(data, &grokParser); err != nil {
			return err
		}
		processor.Definition = grokParser
	case MESSAGE_REMAPPER_PROCESSOR:
		var messageRemapper SourceRemapper
		if err := json.Unmarshal(data, &messageRemapper); err != nil {
			return err
		}
		processor.Definition = messageRemapper
	case NESTED_PIPELINE_PROCESSOR:
		var nestedPipeline NestedPipeline
		if err := json.Unmarshal(data, &nestedPipeline); err != nil {
			return err
		}
		processor.Definition = nestedPipeline
	case SERVICE_REMAPPER_PROCESSOR:
		var serviceRemapper SourceRemapper
		if err := json.Unmarshal(data, &serviceRemapper); err != nil {
			return err
		}
		processor.Definition = serviceRemapper
	case STATUS_REMAPPER_PROCESSOR:
		var statusRemapper SourceRemapper
		if err := json.Unmarshal(data, &statusRemapper); err != nil {
			return err
		}
		processor.Definition = statusRemapper
	case TRACE_ID_REMAPPER_PROCESSOR:
		var traceIdRemapper SourceRemapper
		if err := json.Unmarshal(data, &traceIdRemapper); err != nil {
			return err
		}
		processor.Definition = traceIdRemapper
	case URL_PARSER_PROCESSOR:
		var urlParser UrlParser
		if err := json.Unmarshal(data, &urlParser); err != nil {
			return err
		}
		processor.Definition = urlParser
	case USER_AGENT_PARSER_PROCESSOR:
		var userAgentParser UserAgentParser
		if err := json.Unmarshal(data, &userAgentParser); err != nil {
			return err
		}
		processor.Definition = userAgentParser
	default:
		return fmt.Errorf("cannot unmarshal processor of type: %s", *processorHandler.Type)
	}
	return nil
}
