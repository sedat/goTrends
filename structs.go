package gotrends

type DataHouse struct {
	Trends   []Widget `json:"widgets"`
	Country  string
	Keyword  string
	Property string
}

type Trends struct {
	Widgets                       []Widget        `json:"widgets"`
	Keywords                      []TrendsKeyword `json:"keywords"`
	TimeRanges                    []string        `json:"timeRanges"`
	Examples                      []interface{}   `json:"examples"`
	ShareText                     string          `json:"shareText"`
	ShouldShowMultiHeatMapMessage bool            `json:"shouldShowMultiHeatMapMessage"`
}

type TrendsKeyword struct {
	Keyword string `json:"keyword"`
	Name    string `json:"name"`
	Type    string `json:"type"`
}

type Widget struct {
	Request             Request    `json:"request"`
	LineAnnotationText  *string    `json:"lineAnnotationText,omitempty"`
	Bullets             []Bullet   `json:"bullets"`
	ShowLegend          *bool      `json:"showLegend,omitempty"`
	ShowAverages        *bool      `json:"showAverages,omitempty"`
	HelpDialog          HelpDialog `json:"helpDialog"`
	Token               string     `json:"token"`
	ID                  string     `json:"id"`
	Type                string     `json:"type"`
	Title               string     `json:"title"`
	Template            string     `json:"template"`
	EmbedTemplate       string     `json:"embedTemplate"`
	Version             string     `json:"version"`
	IsLong              bool       `json:"isLong"`
	IsCurated           bool       `json:"isCurated"`
	Geo                 *string    `json:"geo,omitempty"`
	Resolution          *string    `json:"resolution,omitempty"`
	SearchInterestLabel *string    `json:"searchInterestLabel,omitempty"`
	DisplayMode         *string    `json:"displayMode,omitempty"`
	Color               *string    `json:"color,omitempty"`
	Index               *int64     `json:"index,omitempty"`
	Bullet              *string    `json:"bullet,omitempty"`
	KeywordName         *string    `json:"keywordName,omitempty"`
}

type Bullet struct {
	Text string `json:"text"`
}

type HelpDialog struct {
	Title   string  `json:"title"`
	Content string  `json:"content"`
	URL     *string `json:"url,omitempty"`
}

type Request struct {
	Time               *string             `json:"time,omitempty"`
	Resolution         *string             `json:"resolution,omitempty"`
	Locale             *string             `json:"locale,omitempty"`
	ComparisonItem     []Restriction       `json:"comparisonItem"`
	RequestOptions     RequestOptions      `json:"requestOptions"`
	Geo                *Geo                `json:"geo,omitempty"`
	Restriction        *Restriction        `json:"restriction,omitempty"`
	KeywordType        *string             `json:"keywordType,omitempty"`
	Metric             []string            `json:"metric"`
	TrendinessSettings *TrendinessSettings `json:"trendinessSettings,omitempty"`
	Language           *string             `json:"language,omitempty"`
}

type Restriction struct {
	Geo                            *Geo                       `json:"geo,omitempty"`
	ComplexKeywordsRestriction     ComplexKeywordsRestriction `json:"complexKeywordsRestriction"`
	Time                           *string                    `json:"time,omitempty"`
	OriginalTimeRangeForExploreURL *string                    `json:"originalTimeRangeForExploreUrl,omitempty"`
}

type ComplexKeywordsRestriction struct {
	Keyword []ComplexKeywordsRestrictionKeyword `json:"keyword"`
}

type ComplexKeywordsRestrictionKeyword struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Geo struct {
	Country string `json:"country"`
}

type RequestOptions struct {
	Property string `json:"property"`
	Backend  string `json:"backend"`
	Category int64  `json:"category"`
}

type TrendinessSettings struct {
	CompareTime string `json:"compareTime"`
}
