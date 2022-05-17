package aviation

import (
	"context"
)

// Taf represents a taf object.
type Taf struct {
	AWG
	Data TafData `xml:"data,omitempty"`
}

// TafData is part of taf object.
type TafData struct {
	Text       string `xml:",chardata"`
	NumResults string `xml:"num_results,attr,omitempty"`
	TAF        []TAF  `xml:"TAF,omitempty"`
}

// TAF is part of TafData object.
type TAF struct {
	Text          string     `xml:",chardata"`
	RawText       string     `xml:"raw_text"`
	StationID     string     `xml:"station_id"`
	IssueTime     string     `xml:"issue_time"`
	BulletinTime  string     `xml:"bulletin_time"`
	ValidTimeFrom string     `xml:"valid_time_from"`
	ValidTimeTo   string     `xml:"valid_time_to"`
	Latitude      string     `xml:"latitude"`
	Longitude     string     `xml:"longitude"`
	ElevationM    string     `xml:"elevation_m"`
	Forecast      []Forecast `xml:"forecast"`
}

// Forecast is part of taf object.
type Forecast struct {
	Text                string       `xml:",chardata"`
	FcstTimeFrom        string       `xml:"fcst_time_from"`
	FcstTimeTo          string       `xml:"fcst_time_to"`
	WindDirDegrees      string       `xml:"wind_dir_degrees"`
	WindSpeedKt         string       `xml:"wind_speed_kt"`
	VisibilityStatuteMi string       `xml:"visibility_statute_mi"`
	SkyCondition        SkyCondition `xml:"sky_condition"`
	ChangeIndicator     string       `xml:"change_indicator"`
	Probability         string       `xml:"probability"`
	WxString            string       `xml:"wx_string"`
	TimeBecoming        string       `xml:"time_becoming"`
}

// TafService manages communication with the TAF API of client.
type TafService service

// Get returns a Taf object.
func (s *TafService) Get(opts Options) (*Taf, *Response, error) {
	req, err := s.client.NewRequest("tafs", "retrieve", opts)
	if err != nil {
		return nil, nil, err
	}

	t := new(Taf)
	r, err := s.client.Do(context.Background(), req, t)
	if err != nil {
		return nil, r, err
	}

	return t, r, nil
}
