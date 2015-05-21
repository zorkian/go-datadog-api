package datadog

import (
	"fmt"
	"strconv"
)

type Widget interface{}

func NewTimeseriesWidget(
	x, y, width, height int,
	title bool, titleAlign string, titleSize TextSize, titleText string,
	timeframe string,
	requests []TimeseriesRequest) Widget {
	return &TimeseriesWidget{
		X:          x,
		Y:          y,
		Width:      width,
		Height:     height,
		Title:      title,
		TitleAlign: titleAlign,
		TitleSize:  titleSize,
		TitleText:  titleText,
		Timeframe:  timeframe,
		TileDef: TileDef{
			Viz:      "timeseries",
			Requests: requests,
		},
		Type: "timeseries",
	}
}

type TimeseriesWidget struct {
	BoardId    int      `json:"board_id,omitempty"`
	Height     int      `json:"height"`
	Legend     bool     `json:"legend"`
	TileDef    TileDef  `json:"tile_def"`
	Timeframe  string   `json:"timeframe"`
	Title      bool     `json:"title"`
	TitleAlign string   `json:"string"`
	TitleSize  TextSize `json:"title_size"`
	TitleText  string   `json:"title_text"`
	Type       string   `json:"string"`
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

func NewQueryValueWidget(
	x, y, width, height int,
	title bool, titleAlign string, titleSize TextSize, titleText,
	textAlign string, textSize TextSize,
	timeframe, timeframeAggregator,
	aggregator, query string) Widget {
	return &QueryValueWidget{
		X:                   x,
		Y:                   y,
		Width:               width,
		Height:              height,
		Title:               title,
		TitleAlign:          titleAlign,
		TitleSize:           titleSize,
		TitleText:           titleText,
		TextAlign:           textAlign,
		TextSize:            textSize,
		Timeframe:           timeframe,
		TimeframeAggregator: timeframeAggregator,
		Aggregator:          aggregator,
		Query:               query,
		Type:                "query_value",
	}
}

type QueryValueWidget struct {
	Timeframe           string              `json:"timeframe"`
	TimeframeAggregator string              `json:"aggr"`
	Aggregator          string              `json:"aggregator"`
	BoardId             int                 `json:"board_id,omitempty"`
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

func NewToplistWidget(
	x, y, width, height int,
	title bool, titleAlign string, titleSize TextSize, titleText string,
	timeframe string,
	request TimeseriesRequest) Widget {
	return &ToplistWidget{
		X:          x,
		Y:          y,
		Width:      width,
		Height:     height,
		Title:      title,
		TitleAlign: titleAlign,
		TitleSize:  titleSize,
		TitleText:  titleText,
		TileDef: TileDef{
			Viz:      "toplist",
			Requests: []TimeseriesRequest{request},
		},
		Timeframe: timeframe,
		Type:      "toplist",
	}
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

func NewEventStreamWidget(
	x, y, width, height int,
	title bool, titleAlign string, titleSize TextSize, titleText string,
	timeframe string,
	query, eventSize string) Widget {
	return &EventStreamWidget{
		X:          x,
		Y:          y,
		Width:      width,
		Height:     height,
		Title:      title,
		TitleAlign: titleAlign,
		TitleSize:  titleSize,
		TitleText:  titleText,
		Timeframe:  timeframe,
		Query:      query,
		EventSize:  eventSize,
		Type:       "event_stream",
	}
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

func NewFreeTextWidget(x, y, width, height int, text string, size int, align string) Widget {
	return &FreeTextWidget{
		X:         x,
		Y:         y,
		Width:     width,
		Height:    height,
		Text:      text,
		FontSize:  fmt.Sprintf("%d", size),
		TextAlign: align,
		Type:      "free_text",
	}
}

type FreeTextWidget struct {
	BoardId    int      `json:"board_id,omitempty"`
	Color      string   `json:"color,omitempty"`
	FontSize   string   `json:"font_size,omitempty"`
	Height     int      `json:"height,omitempty"`
	Text       string   `json:"text"`
	TextAlign  string   `json:"text_align"`
	Title      bool     `json:"title"`
	TitleAlign string   `json:"title_align"`
	TitleSize  TextSize `json:"title_size"`
	TitleText  string   `json:"title_text"`
	Type       string   `json:"type"`
	Width      int      `json:"width"`
	X          int      `json:"x"`
	Y          int      `json:"y"`
}
