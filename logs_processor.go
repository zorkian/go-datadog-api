package datadog

import (
	"encoding/json"
	"fmt"
)

const (
	ArithmeticProcessorType = "arithmetic-processor"
	AttributeRemapperType   = "attribute-remapper"
	CategoryProcessorType   = "category-processor"
	DateRemapperType        = "date-remapper"
	GrokParserType          = "grok-parser"
	MessageRemapperType     = "message-remapper"
	NestedPipelineType      = "pipeline"
	ServiceRemapperType     = "service-remapper"
	StatusRemapperType      = "status-remapper"
	TraceIdRemapperType     = "trace-id-remapper"
	UrlParserType           = "url-parser"
	UserAgentParserType     = "user-agent-parser"
)

// LogsProcessor struct represents the processor object from Config API.
type LogsProcessor struct {
	Name       *string     `json:"name"`
	IsEnabled  *bool       `json:"is_enabled"`
	Type       *string     `json:"type"`
	Definition interface{} `json:"definition"`
}

// ArithmeticProcessor struct represents unique part of arithmetic processor
// object from config API.
type ArithmeticProcessor struct {
	Expression       *string `json:"expression"`
	Target           *string `json:"target"`
	IsReplaceMissing *bool   `json:"is_replace_missing"`
}

// AttributeRemapper struct represents unique part of attribute remapper object
// from config API.
type AttributeRemapper struct {
	Sources            []string `json:"sources"`
	SourceType         *string  `json:"source_type"`
	Target             *string  `json:"target"`
	TargetType         *string  `json:"target_type"`
	PreserveSource     *bool    `json:"preserve_source"`
	OverrideOnConflict *bool    `json:"override_on_conflict"`
}

// CategoryProcessor struct represents unique part of category processor object
// from config API.
type CategoryProcessor struct {
	Target     *string    `json:"target"`
	Categories []Category `json:"categories"`
}

// Category represents category object from config API.
type Category struct {
	Name   *string              `json:"name"`
	Filter *FilterConfiguration `json:"filter"`
}

// SourceRemapper represents the object from config API that contains
// only a list of sources.
type SourceRemapper struct {
	Sources []string `json:"sources"`
}

// GrokParser represents the grok parser processor object from config API.
type GrokParser struct {
	Source   *string   `json:"source"`
	GrokRule *GrokRule `json:"grok"`
}

// GrokRule represents the rules for grok parser from config API.
type GrokRule struct {
	SupportRules *string `json:"support_rules"`
	MatchRules   *string `json:"match_rules"`
}

// NestedPipeline represents the pipeline as processor from config API.
type NestedPipeline struct {
	Filter     *FilterConfiguration `json:"filter"`
	Processors []LogsProcessor      `json:"processors,omitempty"`
}

// UrlParser represents the url parser from config API.
type UrlParser struct {
	Sources                []string `json:"sources"`
	Target                 *string  `json:"target"`
	NormalizeEndingSlashes *bool    `json:"normalize_ending_slashes"`
}

// UserAgentParser represents the user agent parser from config API.
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

// MarshalJSON serializes logsprocessor struct to config API compatible json object.
func (processor *LogsProcessor) MarshalJSON() ([]byte, error) {
	var mapProcessor = make(map[string]interface{})
	switch *processor.Type {
	case ArithmeticProcessorType:
		if err := convert(processor.Definition.(ArithmeticProcessor), mapProcessor); err != nil {
			return nil, err
		}
	case AttributeRemapperType:
		if err := convert(processor.Definition.(AttributeRemapper), mapProcessor); err != nil {
			return nil, err
		}
	case CategoryProcessorType:
		if err := convert(processor.Definition.(CategoryProcessor), mapProcessor); err != nil {
			return nil, err
		}
	case DateRemapperType,
		MessageRemapperType,
		ServiceRemapperType,
		StatusRemapperType,
		TraceIdRemapperType:
		if err := convert(processor.Definition.(SourceRemapper), mapProcessor); err != nil {
			return nil, err
		}
	case GrokParserType:
		if err := convert(processor.Definition.(GrokParser), mapProcessor); err != nil {
			return nil, err
		}
	case NestedPipelineType:
		if err := convert(processor.Definition.(NestedPipeline), mapProcessor); err != nil {
			return nil, err
		}
	case UrlParserType:
		if err := convert(processor.Definition.(UrlParser), mapProcessor); err != nil {
			return nil, err
		}
	case UserAgentParserType:
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

// UnmarshalJSON deserializes the config API json object to LogsProcessor struct.
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
	case ArithmeticProcessorType:
		var arithmeticProcessor ArithmeticProcessor
		if err := json.Unmarshal(data, &arithmeticProcessor); err != nil {
			return err
		}
		processor.Definition = arithmeticProcessor
	case AttributeRemapperType:
		var attributeRemapper AttributeRemapper
		if err := json.Unmarshal(data, &attributeRemapper); err != nil {
			return err
		}
		processor.Definition = attributeRemapper
	case CategoryProcessorType:
		var categoryProcessor CategoryProcessor
		if err := json.Unmarshal(data, &categoryProcessor); err != nil {
			return err
		}
		processor.Definition = categoryProcessor
	case DateRemapperType,
		MessageRemapperType,
		ServiceRemapperType,
		StatusRemapperType,
		TraceIdRemapperType:
		var sourceRemapper SourceRemapper
		if err := json.Unmarshal(data, &sourceRemapper); err != nil {
			return err
		}
		processor.Definition = sourceRemapper
	case GrokParserType:
		var grokParser GrokParser
		if err := json.Unmarshal(data, &grokParser); err != nil {
			return err
		}
		processor.Definition = grokParser
	case NestedPipelineType:
		var nestedPipeline NestedPipeline
		if err := json.Unmarshal(data, &nestedPipeline); err != nil {
			return err
		}
		processor.Definition = nestedPipeline
	case UrlParserType:
		var urlParser UrlParser
		if err := json.Unmarshal(data, &urlParser); err != nil {
			return err
		}
		processor.Definition = urlParser
	case UserAgentParserType:
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
