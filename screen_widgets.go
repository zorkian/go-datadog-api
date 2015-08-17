package datadog

import (
	"fmt"
	"strconv"
)

type TimeseriesWidget struct {
	Height     int      `json:"height"`
	Legend     bool     `json:"legend"`
	TileDef    TileDef  `json:"tile_def"`
	Timeframe  string   `json:"timeframe"`
	Title      bool     `json:"title"`
	TitleAlign string   `json:"title_align"`
	TitleSize  TextSize `json:"title_size"`
	TitleText  string   `json:"title_text"`
	Type       string   `json:"type"`
	Width      int      `json:"width"`
	X          int      `json:"x"`
	Y          int      `json:"y"`
}

type TextSize struct {
	Size int
	Auto bool
}

func (size *TextSize) UnmarshalJSON(data []byte) error {
	if string(data) == "\"auto\"" {
		size.Auto = true
		return nil
	}

	num, err := strconv.Atoi(string(data))
	if err != nil {
		return err
	}
	size.Size = num

	return nil
}
func (size *TextSize) MarshalJSON() ([]byte, error) {
	if size.Auto {
		return []byte("\"auto\""), nil
	}

	return []byte(fmt.Sprintf("%d", size.Size)), nil
}

type TileDef struct {
	Events   []TileDefEvent      `json:"events"`
	Requests []TimeseriesRequest `json:"requests"`
	Viz      string              `json:"viz"`
}

func NewTimeseriesRequest(rtype string, query string) TimeseriesRequest {
	return TimeseriesRequest{
		Query: query,
		Type:  rtype,
	}
}

type TimeseriesRequest struct {
	Query              string                 `json:"q"`
	Type               string                 `json:"type,omitempty"`
	ConditionalFormats []ConditionalFormat    `json:"conditional_formats,omitempty"`
	Style              TimeseriesRequestStyle `json:"style,omitempty"`
}

type TimeseriesRequestStyle struct {
	Palette string `json:"palette"`
}

type TileDefEvent struct {
	Query string `json:"q"`
}

type QueryValueWidget struct {
	Timeframe           string              `json:"timeframe"`
	TimeframeAggregator string              `json:"aggr"`
	Aggregator          string              `json:"aggregator"`
	CalcFunc            string              `json:"calc_func"`
	ConditionalFormats  []ConditionalFormat `json:"conditional_formats"`
	Height              int                 `json:"height"`
	IsValidQuery        bool                `json:is_valid_query,omitempty"`
	Metric              string              `json:"metric"`
	MetricType          string              `json:"metric_type"`
	Precision           int                 `json:"precision"`
	Query               string              `json:"query"`
	ResultCalcFunc      string              `json:"res_calc_func"`
	Tags                []string            `json:"tags"`
	TextAlign           string              `json:"text_align"`
	TextSize            TextSize            `json:"text_size"`
	Title               bool                `json:"title"`
	TitleAlign          string              `json:"title_align"`
	TitleSize           TextSize            `json:"title_size"`
	TitleText           string              `json:"title_text"`
	Type                string              `json:"type"`
	Unit                string              `json:"auto"`
	Width               int                 `json:"width"`
	X                   int                 `json:"x"`
	Y                   int                 `json:"y"`
}
type ConditionalFormat struct {
	Color      string `json:"color"`
	Comparator string `json:"comparator"`
	Inverted   bool   `json:"invert"`
	Value      int    `json:"value"`
}

type ToplistWidget struct {
	Height     int      `json:"height"`
	Legend     bool     `json:"legend"`
	LegendSize string   `json:"legend_size"`
	TileDef    TileDef  `json:"tile_def"`
	Timeframe  string   `json:"timeframe"`
	Title      bool     `json:"title"`
	TitleAlign string   `json:"title_align"`
	TitleSize  TextSize `json:"title_size"`
	TitleText  string   `json:"title_text"`
	Type       string   `json:"type"`
	Width      int      `json:"width"`
	X          int      `json:"x"`
	Y          int      `json:"y"`
}

type EventStreamWidget struct {
	EventSize  string   `json:"event_size"`
	Height     int      `json:"height"`
	Query      string   `json:"query"`
	Timeframe  string   `json:"timeframe"`
	Title      bool     `json:"title"`
	TitleAlign string   `json:"title_align"`
	TitleSize  TextSize `json:"title_size"`
	TitleText  string   `json:"title_text"`
	Type       string   `json:"type"`
	Width      int      `json:"width"`
	X          int      `json:"x"`
	Y          int      `json:"y"`
}

type FreeTextWidget struct {
	Color     string `json:"color,omitempty"`
	FontSize  string `json:"font_size,omitempty"`
	Height    int    `json:"height,omitempty"`
	Text      string `json:"text"`
	TextAlign string `json:"text_align"`
	Type      string `json:"type"`
	Width     int    `json:"width"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
}

type ImageWidget struct {
	Height     int      `json:"height"`
	Sizing     string   `json:"sizing"`
	Title      bool     `json:"title"`
	TitleAlign string   `json:"title_align"`
	TitleSize  TextSize `json:"title_size"`
	TitleText  string   `json:"title_text"`
	Type       string   `json:"type"`
	Url        string   `json:"url"`
	Width      int      `json:"width"`
	X          int      `json:"x"`
	Y          int      `json:"y"`
}
