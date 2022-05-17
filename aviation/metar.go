package aviation

import (
	"context"
)

// Metar represents a metar object.
type Metar struct {
	AWG
	Data MetarData `xml:"data,omitempty"`
}

// MetarData is part of metar object.
type MetarData struct {
	Text       string  `xml:",chardata"`
	NumResults string  `xml:"num_results,attr,omitempty"`
	METAR      []METAR `xml:"METAR,omitempty"`
}

// METAR is part of MetarData object.
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

// QualityControlFlags is part of metar object.
type QualityControlFlags struct {
	Text string `xml:",chardata"`
	Auto string `xml:"auto,omitempty"`
}

// SkyCondition is part of metar object.
type SkyCondition struct {
	Text           string `xml:",chardata"`
	SkyCover       string `xml:"sky_cover,attr,omitempty"`
	CloudBaseFtAgl string `xml:"cloud_base_ft_agl,attr,omitempty"`
	CloudType      string `xml:"cloud_type,attr"`
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
