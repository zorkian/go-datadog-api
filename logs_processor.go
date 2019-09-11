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
	Type       *string     `json:"type"`
	Definition interface{} `json:"definition"`
}

type ArithmeticProcessor struct {
	Expression       *string `json:"expression"`
	Target           *string `json:"target"`
	IsReplaceMissing *bool   `json:"is_replace_missing"`
}

type AttributeRemapper struct {
	Sources            []string `json:"sources"`
	SourceType         *string  `json:"source_type"`
	Target             *string  `json:"target"`
	TargetType         *string  `json:"target_type"`
	PreserveSource     *bool    `json:"preserve_source"`
	OverrideOnConflict *bool    `json:"override_on_conflict"`
}

type CategoryProcessor struct {
	Target     *string    `json:"target"`
	Categories []Category `json:"categories"`
}

type Category struct {
	Name   *string              `json:"name"`
	Filter *FilterConfiguration `json:"filter"`
}

type SourceRemapper struct {
	Sources []string `json:"sources"`
}

type GrokParser struct {
	Source   *string   `json:"source"`
	GrokRule *GrokRule `json:"grok"`
}

type GrokRule struct {
	SupportRules *string `json:"support_rules"`
	MatchRules   *string `json:"match_rules"`
}

type NestedPipeline struct {
	Filter     *FilterConfiguration `json:"filter"`
	Processors []LogsProcessor      `json:"processors,omitempty"`
}

type UrlParser struct {
	Sources                []string `json:"sources"`
	Target                 *string  `json:"target"`
	NormalizeEndingSlashes *bool    `json:"normalize_ending_slashes"`
}

type UserAgentParser struct {
	Sources   []string `json:"sources"`
	Target    *string  `json:"target"`
	IsEncoded *bool    `json:"is_encoded"`
}

// convert converts the first argument of type interface{} to a map of string and interface{}
// and sign the result to the second argument.
func convert(definition interface{}, processor map[string]interface{}) error {
	inrec, err := json.Marshal(definition)
	if err != nil {
		return err
	}
	if err = json.Unmarshal(inrec, &processor); err != nil {
		return err
	}
	return nil
}

func (processor *LogsProcessor) MarshalJSON() ([]byte, error) {
	var mapProcessor = make(map[string]interface{})
	switch *processor.Type {
	case ARITHMETIC_PROCESSOR:
		if err := convert(processor.Definition.(ArithmeticProcessor), mapProcessor); err != nil {
			return nil, err
		}
	case ATTIBUTE_REMAPPER_PROCESSOR:
		if err := convert(processor.Definition.(AttributeRemapper), mapProcessor); err != nil {
			return nil, err
		}
	case CATEGORY_PROCESSOR:
		if err := convert(processor.Definition.(CategoryProcessor), mapProcessor); err != nil {
			return nil, err
		}
	case DATE_REMAPPER_PROCESSOR,
		MESSAGE_REMAPPER_PROCESSOR,
		SERVICE_REMAPPER_PROCESSOR,
		STATUS_REMAPPER_PROCESSOR,
		TRACE_ID_REMAPPER_PROCESSOR:
		if err := convert(processor.Definition.(SourceRemapper), mapProcessor); err != nil {
			return nil, err
		}
	case GROK_PARSER_PROCESSOR:
		if err := convert(processor.Definition.(GrokParser), mapProcessor); err != nil {
			return nil, err
		}
	case NESTED_PIPELINE_PROCESSOR:
		if err := convert(processor.Definition.(NestedPipeline), mapProcessor); err != nil {
			return nil, err
		}
	case URL_PARSER_PROCESSOR:
		if err := convert(processor.Definition.(UrlParser), mapProcessor); err != nil {
			return nil, err
		}
	case USER_AGENT_PARSER_PROCESSOR:
		if err := convert(processor.Definition.(UserAgentParser), mapProcessor); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("cannot marshal processor of type: %s", *processor.Type)
	}
	mapProcessor["name"] = processor.Name
	mapProcessor["is_enabled"] = processor.IsEnabled
	mapProcessor["type"] = processor.Type
	jsn, err := json.Marshal(mapProcessor)
	if err != nil {
		return nil, err
	}
	return jsn, err
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
	processor.Type = processorHandler.Type

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
	case DATE_REMAPPER_PROCESSOR,
		MESSAGE_REMAPPER_PROCESSOR,
		SERVICE_REMAPPER_PROCESSOR,
		STATUS_REMAPPER_PROCESSOR,
		TRACE_ID_REMAPPER_PROCESSOR:
		var sourceRemapper SourceRemapper
		if err := json.Unmarshal(data, &sourceRemapper); err != nil {
			return err
		}
		processor.Definition = sourceRemapper
	case GROK_PARSER_PROCESSOR:
		var grokParser GrokParser
		if err := json.Unmarshal(data, &grokParser); err != nil {
			return err
		}
		processor.Definition = grokParser
	case NESTED_PIPELINE_PROCESSOR:
		var nestedPipeline NestedPipeline
		if err := json.Unmarshal(data, &nestedPipeline); err != nil {
			return err
		}
		processor.Definition = nestedPipeline
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
