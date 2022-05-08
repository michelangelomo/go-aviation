package aviation

import (
	"context"
)

type Taf struct {
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
		TAF        []struct {
			Text          string `xml:",chardata"`
			RawText       string `xml:"raw_text"`
			StationID     string `xml:"station_id"`
			IssueTime     string `xml:"issue_time"`
			BulletinTime  string `xml:"bulletin_time"`
			ValidTimeFrom string `xml:"valid_time_from"`
			ValidTimeTo   string `xml:"valid_time_to"`
			Latitude      string `xml:"latitude"`
			Longitude     string `xml:"longitude"`
			ElevationM    string `xml:"elevation_m"`
			Forecast      []struct {
				Text                string `xml:",chardata"`
				FcstTimeFrom        string `xml:"fcst_time_from"`
				FcstTimeTo          string `xml:"fcst_time_to"`
				WindDirDegrees      string `xml:"wind_dir_degrees"`
				WindSpeedKt         string `xml:"wind_speed_kt"`
				VisibilityStatuteMi string `xml:"visibility_statute_mi"`
				SkyCondition        []struct {
					Text           string `xml:",chardata"`
					SkyCover       string `xml:"sky_cover,attr"`
					CloudBaseFtAgl string `xml:"cloud_base_ft_agl,attr"`
					CloudType      string `xml:"cloud_type,attr"`
				} `xml:"sky_condition"`
				ChangeIndicator string `xml:"change_indicator"`
				Probability     string `xml:"probability"`
				WxString        string `xml:"wx_string"`
				TimeBecoming    string `xml:"time_becoming"`
			} `xml:"forecast"`
		} `xml:"TAF"`
	} `xml:"data"`
}

type TafOptions struct {
	Stations       string
	HoursBeforeNow string
}

type TafService service

func (s *TafService) Get(opts TafOptions) (*Taf, *Response, error) {
	params := map[string]string{
		"stationString":  opts.Stations,
		"hoursBeforeNow": opts.HoursBeforeNow,
	}
	req, err := s.client.NewRequest("tafs", "retrieve", params)
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
