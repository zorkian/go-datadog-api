package datadog

type TextSize struct {
	Size *int
	Auto *bool
}

type TileDef struct {
	Events   []TileDefEvent      `mapstructure:"events"    json:"events,omitempty"`
	Markers  []TimeseriesMarker  `mapstructure:"markers"   json:"markers,omitempty"`
	Requests []TimeseriesRequest `mapstructure:"requests"  json:"requests,omitempty"`
	Viz      *string             `mapstructure:"viz"       json:"viz,omitempty"`
}

type Time struct {
	LiveSpan	*string		`mapstructure:"live_span" json:"live_span,omitempty"`
}

type TimeseriesRequest struct {
	Query              *string                 `mapstructure:"q"                   json:"q,omitempty"`
	Type               *string                 `mapstructure:"type"                json:"type,omitempty"`
	ConditionalFormats []ConditionalFormat     `mapstructure:"conditional_formats" json:"conditional_formats,omitempty"`
	Style              *TimeseriesRequestStyle `mapstructure:"style"               json:"style,omitempty"`
}

type TimeseriesRequestStyle struct {
	Palette *string `mapstructure:"palette" json:"palette,omitempty"`
}

type TimeseriesMarker struct {
	Label *string `mapstructure:"label" json:"label,omitempty"`
	Type  *string `mapstructure:"type"  json:"type,omitempty"`
	Value *string `mapstructure:"value" json:"value,omitempty"`
}

type TileDefEvent struct {
	Query *string `mapstructure:"q" json:"q"`
}

type AlertValueWidget struct {
	TitleSize    *int    `mapstructure:"title_size"    json:"title_size,omitempty"`
	Title        *bool   `mapstructure:"title"         json:"title,omitempty"`
	TitleAlign   *string `mapstructure:"title_align"   json:"title_align,omitempty"`
	TextAlign    *string `mapstructure:"text_align"    json:"text_align,omitempty"`
	TitleText    *string `mapstructure:"title_text"    json:"title_text,omitempty"`
	Precision    *int    `mapstructure:"precision"     json:"precision,omitempty"`
	AlertId      *int    `mapstructure:"alert_id"      json:"alert_id,omitempty"`
	Time	     *Time   `mapstructure:"time"          json:"time,omitempty"`
	AddTimeframe *bool   `mapstructure:"add_timeframe" json:"add_timeframe,omitempty"`
	Y            *int    `mapstructure:"y"             json:"y,omitempty"`
	X            *int    `mapstructure:"x"             json:"x,omitempty"`
	TextSize     *string `mapstructure:"text_size"     json:"text_size,omitempty"`
	Height       *int    `mapstructure:"height"        json:"height,omitempty"`
	Width        *int    `mapstructure:"width"         json:"width,omitempty"`
	Type         *string `mapstructure:"type"          json:"type,omitempty"`
	Unit         *string `mapstructure:"unit"          json:"unit,omitempty"`
}

type ChangeWidget struct {
	TitleSize  *int     `mapstructure:"title_size"    json:"title_size,omitempty"`
	Title      *bool    `mapstructure:"title"         json:"title,omitempty"`
	TitleAlign *string  `mapstructure:"title_align"   json:"title_align,omitempty"`
	TitleText  *string  `mapstructure:"title_text"    json:"title_text,omitempty"`
	Height     *int     `mapstructure:"height"        json:"height,omitempty"`
	Width      *int     `mapstructure:"width"         json:"width,omitempty"`
	X          *int     `mapstructure:"y"             json:"y,omitempty"`
	Y          *int     `mapstructure:"x"             json:"x,omitempty"`
	Aggregator *string  `mapstructure:"aggregator"    json:"aggregator,omitempty"`
	TileDef    *TileDef `mapstructure:"tile_def"      json:"tile_def,omitempty"`
}

type GraphWidget struct {
	TitleSize  *int     `mapstructure:"title_size"  json:"title_size,omitempty"`
	Title      *bool    `mapstructure:"title"       json:"title,omitempty"`
	TitleAlign *string  `mapstructure:"title_align" json:"title_align,omitempty"`
	TitleText  *string  `mapstructure:"title_text"  json:"title_text,omitempty"`
	Height     *int     `mapstructure:"height"      json:"height,omitempty"`
	Width      *int     `mapstructure:"width"       json:"width,omitempty"`
	X          *int     `mapstructure:"y"           json:"y,omitempty"`
	Y          *int     `mapstructure:"x"           json:"x,omitempty"`
	Type       *string  `mapstructure:"type"        json:"type,omitempty"`
	Time	   *Time    `mapstructure:"time"        json:"time,omitempty"`
	LegendSize *int     `mapstructure:"legend_size" json:"legend_size,omitempty"`
	Legend     *bool    `mapstructure:"legend"      json:"legend,omitempty"`
	TileDef    *TileDef `mapstructure:"tile_def"    json:"tile_def,omitempty"`
}

type EventTimelineWidget struct {
	TitleSize  *int    `mapstructure:"title_size"  json:"title_size,omitempty"`
	Title      *bool   `mapstructure:"title"       json:"title,omitempty"`
	TitleAlign *string `mapstructure:"title_align" json:"title_align,omitempty"`
	TitleText  *string `mapstructure:"title_text"  json:"title_text,omitempty"`
	Height     *int    `mapstructure:"height"      json:"height,omitempty"`
	Width      *int    `mapstructure:"width"       json:"width,omitempty"`
	X          *int    `mapstructure:"y"           json:"y,omitempty"`
	Y          *int    `mapstructure:"x"           json:"x,omitempty"`
	Type       *string `mapstructure:"type"        json:"type,omitempty"`
	Time       *Time   `mapstructure:"time"        json:"time,omitempty"`
	Query      *string `mapstructure:"query"       json:"query,omitempty"`
}

type AlertGraphWidget struct {
	TitleSize    *int    `mapstructure:"title_size"     json:"title_size,omitempty"`
	VizType      *string `mapstructure:"timeseries"     json:"timeseries,omitempty"`
	Title        *bool   `mapstructure:"title"          json:"title,omitempty"`
	TitleAlign   *string `mapstructure:"title_align"    json:"title_align,omitempty"`
	TitleText    *string `mapstructure:"title_text"     json:"title_text,omitempty"`
	Height       *int    `mapstructure:"height"         json:"height,omitempty"`
	Width        *int    `mapstructure:"width"          json:"width,omitempty"`
	X            *int    `mapstructure:"y"              json:"y,omitempty"`
	Y            *int    `mapstructure:"x"              json:"x,omitempty"`
	AlertId      *int    `mapstructure:"alert_id"       json:"alert_id,omitempty"`
	Time         *Time   `mapstructure:"time"           json:"time,omitempty"`
	Type         *string `mapstructure:"type"           json:"type,omitempty"`
	AddTimeframe *bool   `mapstructure:"add_timeframe"  json:"add_timeframe,omitempty"`
}

type HostMapWidget struct {
	TitleSize  *int     `mapstructure:"title_size"   json:"title_size,omitempty"`
	Title      *bool    `mapstructure:"title"        json:"title,omitempty"`
	TitleAlign *string  `mapstructure:"title_align"  json:"title_align,omitempty"`
	TitleText  *string  `mapstructure:"title_text"   json:"title_text,omitempty"`
	Height     *int     `mapstructure:"height"       json:"height,omitempty"`
	Width      *int     `mapstructure:"width"        json:"width,omitempty"`
	X          *int     `mapstructure:"y"            json:"y,omitempty"`
	Y          *int     `mapstructure:"x"            json:"x,omitempty"`
	Query      *string  `mapstructure:"query"        json:"query,omitempty"`
	Time       *Time    `mapstructure:"time"         json:"time,omitempty"`
	LegendSize *int     `mapstructure:"legend_size"  json:"legend_size,omitempty"`
	Type       *string  `mapstructure:"type"         json:"type,omitempty"`
	Legend     *bool    `mapstructure:"legend"       json:"legend,omitempty"`
	TileDef    *TileDef `mapstructure:"tile_def"     json:"tile_def,omitempty"`
}

type CheckStatusWidget struct {
	TitleSize  *int    `mapstructure:"title_size"   json:"title_size,omitempty"`
	Title      *bool   `mapstructure:"title"        json:"title,omitempty"`
	TitleAlign *string `mapstructure:"title_align"  json:"title_align,omitempty"`
	TextAlign  *string `mapstructure:"text_align"   json:"text_align,omitempty"`
	TitleText  *string `mapstructure:"title_text"   json:"title_text,omitempty"`
	Height     *int    `mapstructure:"height"       json:"height,omitempty"`
	Width      *int    `mapstructure:"width"        json:"width,omitempty"`
	X          *int    `mapstructure:"y"            json:"y,omitempty"`
	Y          *int    `mapstructure:"x"            json:"x,omitempty"`
	Tags       *string `mapstructure:"tags"         json:"tags,omitempty"`
	Time       *Time   `mapstructure:"time"         json:"time,omitempty"`
	TextSize   *string `mapstructure:"text_size"    json:"text_size,omitempty"`
	Type       *string `mapstructure:"type"         json:"type,omitempty"`
	Check      *string `mapstructure:"check"        json:"check,omitempty"`
	Group      *string `mapstructure:"group"        json:"group,omitempty"`
	Grouping   *string `mapstructure:"grouping"     json:"grouping,omitempty"`
}

type IFrameWidget struct {
	TitleSize  *int    `mapstructure:"title_size"   json:"title_size,omitempty"`
	Title      *bool   `mapstructure:"title"        json:"title,omitempty"`
	Url        *string `mapstructure:"url"          json:"url,omitempty"`
	TitleAlign *string `mapstructure:"title_align"  json:"title_align,omitempty"`
	TitleText  *string `mapstructure:"title_text"   json:"title_text,omitempty"`
	Height     *int    `mapstructure:"height"       json:"height,omitempty"`
	Width      *int    `mapstructure:"width"        json:"width,omitempty"`
	X          *int    `mapstructure:"y"            json:"y,omitempty"`
	Y          *int    `mapstructure:"x"            json:"x,omitempty"`
	Type       *string `mapstructure:"type"         json:"type,omitempty"`
}

type NoteWidget struct {
	TitleSize    *int    `mapstructure:"title_size"     json:"title_size,omitempty"`
	Title        *bool   `mapstructure:"title"          json:"title,omitempty"`
	RefreshEvery *int    `mapstructure:"refresh_every"  json:"refresh_every,omitempty"`
	TickPos      *string `mapstructure:"tick_pos"       json:"tick_pos,omitempty"`
	TitleAlign   *string `mapstructure:"title_align"    json:"title_align,omitempty"`
	TickEdge     *string `mapstructure:"tick_edge"      json:"tick_edge,omitempty"`
	TextAlign    *string `mapstructure:"text_align"     json:"text_align,omitempty"`
	TitleText    *string `mapstructure:"title_text"     json:"title_text,omitempty"`
	Height       *int    `mapstructure:"height"         json:"height,omitempty"`
	Color        *string `mapstructure:"bgcolor"        json:"bgcolor,omitempty"`
	Html         *string `mapstructure:"html"           json:"html,omitempty"`
	Y            *int    `mapstructure:"y"              json:"y,omitempty"`
	X            *int    `mapstructure:"x"              json:"x,omitempty"`
	FontSize     *int    `mapstructure:"font_size"      json:"font_size,omitempty"`
	Tick         *bool   `mapstructure:"tick"           json:"tick,omitempty"`
	Note         *string `mapstructure:"type"           json:"type,omitempty"`
	Width        *int    `mapstructure:"width"          json:"width,omitempty"`
	AutoRefresh  *bool   `mapstructure:"auto_refresh"   json:"auto_refresh,omitempty"`
}

type TimeseriesWidget struct {
	Height     *int      `mapstructure:"height"       json:"height,omitempty"`
	Legend     *bool     `mapstructure:"legend"       json:"legend,omitempty"`
	TileDef    *TileDef  `mapstructure:"tile_def"     json:"tile_def,omitempty"`
	Time       *Time   `mapstructure:"time"           json:"time,omitempty"`
	Title      *bool     `mapstructure:"title"        json:"title,omitempty"`
	TitleAlign *string   `mapstructure:"title_align"  json:"title_align,omitempty"`
	TitleSize  *TextSize `mapstructure:"title_size"   json:"title_size,omitempty"`
	TitleText  *string   `mapstructure:"title_text"   json:"title_text,omitempty"`
	Type       *string   `mapstructure:"type"         json:"type,omitempty"`
	Width      *int      `mapstructure:"width"        json:"width,omitempty"`
	X          *int      `mapstructure:"x"            json:"x,omitempty"`
	Y          *int      `mapstructure:"y"            json:"y,omitempty"`
}

type QueryValueWidget struct {
	Time                *Time               `mapstructure:"time"                 json:"time,omitempty"`
	TimeframeAggregator *string             `mapstructure:"aggr"                 json:"aggr,omitempty"`
	Aggregator          *string             `mapstructure:"aggregator"           json:"aggregator,omitempty"`
	CalcFunc            *string             `mapstructure:"calc_func"            json:"calc_func,omitempty"`
	ConditionalFormats  []ConditionalFormat `mapstructure:"conditional_formats"  json:"conditional_formats,omitempty"`
	Height              *int                `mapstructure:"height"               json:"height,omitempty"`
	IsValidQuery        *bool               `mapstructure:"is_valid_query"       json:"is_valid_query,omitempty,omitempty"`
	Metric              *string             `mapstructure:"metric"               json:"metric,omitempty"`
	MetricType          *string             `mapstructure:"metric_type"          json:"metric_type,omitempty"`
	Precision           *int                `mapstructure:"precision"            json:"precision,omitempty"`
	Query               *string             `mapstructure:"query"                json:"query,omitempty"`
	ResultCalcFunc      *string             `mapstructure:"res_calc_func"        json:"res_calc_func,omitempty"`
	Tags                []string            `mapstructure:"tags"                 json:"tags,omitempty"`
	TextAlign           *string             `mapstructure:"text_align"           json:"text_align,omitempty"`
	TextSize            *TextSize           `mapstructure:"text_size"            json:"text_size,omitempty"`
	Title               *bool               `mapstructure:"title"                json:"title,omitempty"`
	TitleAlign          *string             `mapstructure:"title_align"          json:"title_align,omitempty"`
	TitleSize           *TextSize           `mapstructure:"title_size"           json:"title_size,omitempty"`
	TitleText           *string             `mapstructure:"title_text"           json:"title_text,omitempty"`
	Type                *string             `mapstructure:"type"                 json:"type,omitempty"`
	Unit                *string             `mapstructure:"auto"                 json:"auto,omitempty"`
	Width               *int                `mapstructure:"width"                json:"width,omitempty"`
	X                   *int                `mapstructure:"x"                    json:"x,omitempty"`
	Y                   *int                `mapstructure:"y"                    json:"y,omitempty"`
}
type ConditionalFormat struct {
	Color      *string `mapstructure:"color"       json:"color,omitempty"`
	Comparator *string `mapstructure:"comparator"  json:"comparator,omitempty"`
	Inverted   *bool   `mapstructure:"invert"      json:"invert,omitempty"`
	Value      *int    `mapstructure:"value"       json:"value,omitempty"`
}

type ToplistWidget struct {
	Height     *int      `mapstructure:"height"       json:"height,omitempty"`
	Legend     *bool     `mapstructure:"legend"       json:"legend,omitempty"`
	LegendSize *int      `mapstructure:"legend_size"  json:"legend_size,omitempty"`
	TileDef    *TileDef  `mapstructure:"tile_def"     json:"tile_def,omitempty"`
	Title      *bool     `mapstructure:"title"        json:"title,omitempty"`
	TitleAlign *string   `mapstructure:"title_align"  json:"title_align,omitempty"`
	TitleSize  *TextSize `mapstructure:"title_size"   json:"title_size,omitempty"`
	TitleText  *string   `mapstructure:"title_text"   json:"title_text,omitempty"`
	Type       *string   `mapstructure:"type"         json:"type,omitempty"`
	Time       *Time     `mapstructure:"time"         json:"time,omitempty"`
	Width      *int      `mapstructure:"width"        json:"width,omitempty"`
	X          *int      `mapstructure:"x"            json:"x,omitempty"`
	Y          *int      `mapstructure:"y"            json:"y,omitempty"`
}

type EventStreamWidget struct {
	EventSize  *string   `mapstructure:"event_size"   json:"event_size,omitempty"`
	Height     *int      `mapstructure:"height"       json:"height,omitempty"`
	Query      *string   `mapstructure:"query"        json:"query,omitempty"`
	Time       *Time     `mapstructure:"time"         json:"time,omitempty"`
	Title      *bool     `mapstructure:"title"        json:"title,omitempty"`
	TitleAlign *string   `mapstructure:"title_align"  json:"title_align,omitempty"`
	TitleSize  *TextSize `mapstructure:"title_size"   json:"title_size,omitempty"`
	TitleText  *string   `mapstructure:"title_text"   json:"title_text,omitempty"`
	Type       *string   `mapstructure:"type"         json:"type,omitempty"`
	Width      *int      `mapstructure:"width"        json:"width,omitempty"`
	X          *int      `mapstructure:"x"            json:"x,omitempty"`
	Y          *int      `mapstructure:"y"            json:"y,omitempty"`
}

type FreeTextWidget struct {
	Color     *string `mapstructure:"color"       json:"color,omitempty"`
	FontSize  *string `mapstructure:"font_size"   json:"font_size,omitempty"`
	Height    *int    `mapstructure:"height"      json:"height,omitempty"`
	Text      *string `mapstructure:"text"        json:"text,omitempty"`
	TextAlign *string `mapstructure:"text_align"  json:"text_align,omitempty"`
	Type      *string `mapstructure:"type"        json:"type,omitempty"`
	Width     *int    `mapstructure:"width"       json:"width,omitempty"`
	X         *int    `mapstructure:"x"           json:"x,omitempty"`
	Y         *int    `mapstructure:"y"           json:"y,omitempty"`
}

type ImageWidget struct {
	Height     *int      `mapstructure:"height"       json:"height,omitempty"`
	Sizing     *string   `mapstructure:"sizing"       json:"sizing,omitempty"`
	Title      *bool     `mapstructure:"title"        json:"title,omitempty"`
	TitleAlign *string   `mapstructure:"title_align"  json:"title_align,omitempty"`
	TitleSize  *TextSize `mapstructure:"title_size"   json:"title_size,omitempty"`
	TitleText  *string   `mapstructure:"title_text"   json:"title_text,omitempty"`
	Type       *string   `mapstructure:"type"         json:"type,omitempty"`
	Url        *string   `mapstructure:"url"          json:"url,omitempty"`
	Width      *int      `mapstructure:"width"        json:"width,omitempty"`
	X          *int      `mapstructure:"x"            json:"x,omitempty"`
	Y          *int      `mapstructure:"y"            json:"y,omitempty"`
}
