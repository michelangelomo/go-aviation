package aviation

import (
	"context"
)

// Metar represents a metar object.
type Metar struct {
	DataSource struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
	} `xml:"data_source"`
	Request struct {
		Text string `xml:",chardata"`
		Type string `xml:"type,attr"`
	} `xml:"request"`
	Errors      string `xml:"errors"`
	Warnings    string `xml:"warnings"`
	TimeTakenMs string `xml:"time_taken_ms"`
	Data        struct {
		Text       string `xml:",chardata"`
		NumResults string `xml:"num_results,attr"`
		METAR      []struct {
			Text                string `xml:",chardata"`
			RawText             string `xml:"raw_text"`
			StationID           string `xml:"station_id"`
			ObservationTime     string `xml:"observation_time"`
			Latitude            string `xml:"latitude"`
			Longitude           string `xml:"longitude"`
			TempC               string `xml:"temp_c"`
			DewpointC           string `xml:"dewpoint_c"`
			WindDirDegrees      string `xml:"wind_dir_degrees"`
			WindSpeedKt         string `xml:"wind_speed_kt"`
			VisibilityStatuteMi string `xml:"visibility_statute_mi"`
			AltimInHg           string `xml:"altim_in_hg"`
			QualityControlFlags struct {
				Text string `xml:",chardata"`
				Auto string `xml:"auto"`
			} `xml:"quality_control_flags"`
			WxString     string `xml:"wx_string"`
			SkyCondition []struct {
				Text           string `xml:",chardata"`
				SkyCover       string `xml:"sky_cover,attr"`
				CloudBaseFtAgl string `xml:"cloud_base_ft_agl,attr"`
			} `xml:"sky_condition"`
			FlightCategory string `xml:"flight_category"`
			MetarType      string `xml:"metar_type"`
			ElevationM     string `xml:"elevation_m"`
		} `xml:"METAR"`
	} `xml:"data"`
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
