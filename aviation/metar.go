package aviation

import (
	"context"
)

// Metar represents a metar object.
type Metar struct {
	DataSource  DataSource `xml:"data_source,omitempty"`
	Request     Request    `xml:"request,omitempty"`
	Errors      string     `xml:"errors,omitempty"`
	Warnings    string     `xml:"warnings,omitempty"`
	TimeTakenMs string     `xml:"time_taken_ms,omitempty"`
	Data        Data       `xml:"data,omitempty"`
}

type DataSource struct {
	Text string `xml:",chardata"`
	Name string `xml:"name,attr,omitempty"`
}

type Request struct {
	Text string `xml:",chardata"`
	Type string `xml:"type,attr,omitempty"`
}

type Data struct {
	Text       string  `xml:",chardata"`
	NumResults string  `xml:"num_results,attr,omitempty"`
	METAR      []METAR `xml:"METAR,omitempty"`
}

type METAR struct {
	Text                string              `xml:",chardata"`
	RawText             string              `xml:"raw_text,omitempty"`
	StationID           string              `xml:"station_id,omitempty"`
	ObservationTime     string              `xml:"observation_time,omitempty"`
	Latitude            string              `xml:"latitude,omitempty"`
	Longitude           string              `xml:"longitude,omitempty"`
	TempC               string              `xml:"temp_c,omitempty"`
	DewpointC           string              `xml:"dewpoint_c,omitempty"`
	WindDirDegrees      string              `xml:"wind_dir_degrees,omitempty"`
	WindSpeedKt         string              `xml:"wind_speed_kt,omitempty"`
	VisibilityStatuteMi string              `xml:"visibility_statute_mi,omitempty"`
	AltimInHg           string              `xml:"altim_in_hg,omitempty"`
	QualityControlFlags QualityControlFlags `xml:"quality_control_flags,omitempty"`
	WxString            string              `xml:"wx_string,omitempty"`
	SkyCondition        []SkyCondition      `xml:"sky_condition,omitempty"`
	FlightCategory      string              `xml:"flight_category,omitempty"`
	MetarType           string              `xml:"metar_type,omitempty"`
	ElevationM          string              `xml:"elevation_m,omitempty"`
}

type QualityControlFlags struct {
	Text string `xml:",chardata"`
	Auto string `xml:"auto,omitempty"`
}

type SkyCondition struct {
	Text           string `xml:",chardata"`
	SkyCover       string `xml:"sky_cover,attr,omitempty"`
	CloudBaseFtAgl string `xml:"cloud_base_ft_agl,attr,omitempty"`
}

// MetarService manages communication with the Metar API of client.
type MetarService service

// Get returns a Metar object.
func (s *MetarService) Get(opts Options) (*Metar, *Response, error) {
	req, err := s.client.NewRequest("metars", "retrieve", opts)
	if err != nil {
		return nil, nil, err
	}

	m := new(Metar)
	r, err := s.client.Do(context.Background(), req, m)
	if err != nil {
		return nil, r, err
	}

	return m, r, nil
}
