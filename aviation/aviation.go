package aviation

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"
)

const (
	baseURL = "https://www.aviationweather.gov/adds/dataserver_current/httpparam"
)

// Client is used to manages connection with Aviation Weather Center.
type Client struct {
	client *http.Client

	service service

	// BaseURL for API requests.
	BaseURL url.URL

	// Services used to interact with the AWC API.
	Metar   *MetarService
	Taf     *TafService
}

type service struct {
	client *Client
}

// Response is the AWC API response.
type Response struct {
	*http.Response
}

// Options provides all fields expected by Aviation Weather Center
// The list of the available parameters is available at:
// * https://www.aviationweather.gov/dataserver/example?datatype=metar
// * https://www.aviationweather.gov/dataserver/example?datatype=taf
// * https://www.aviationweather.gov/dataserver/example?datatype=sigmet
type Options struct {
	Stations                 *string    `json:"stationString,omitempty"`
	HoursBeforeNow           *float32   `json:"hoursBeforeNow,omitempty,string"`
	MostRecent               *bool      `json:"mostRecent,omitempty,string"`
	StartTime                *string    `json:"startTime,omitempty"`
	EndTime                  *string    `json:"endTime,omitempty"`
	MostRecentForEachStation *string    `json:"mostRecentForEachStation,omitempty,string"`
}

// SetStations populates Options field.
func (opts *Options) SetStations(stations string) {
	opts.Stations = &stations
}

// SetHoursBeforeNow populates Options field.
func (opts *Options) SetHoursBeforeNow(hoursBeforeNow float32) {
	opts.HoursBeforeNow = &hoursBeforeNow
}

// SetMostRecent populates Options field.
func (opts *Options) SetMostRecent(mostRecent bool) {
	opts.MostRecent = &mostRecent
}

// SetStartTime populates Options field.
func (opts *Options) SetStartTime(startTime time.Time) {
	s := startTime.UTC().Format(time.RFC3339)
	opts.StartTime = &s
}

// SetEndTime populates Options field.
func (opts *Options) SetEndTime(endTime time.Time) {
	e := endTime.UTC().Format(time.RFC3339)
	opts.EndTime = &e
}

// SetMostRecentForEachStation populates Options field.
func (opts *Options) SetMostRecentForEachStation(mostRecentForEachStation string) {
	opts.MostRecentForEachStation = &mostRecentForEachStation
}

// Client returns the client used by this client.
func (c *Client) Client() *http.Client {
	clientCopy := *c.client
	return &clientCopy
}

// NewClient return a new AWC API client.
// If httpClient is not provided, a new http.Client will be used.
func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(baseURL)

	c := &Client{client: httpClient, BaseURL: *baseURL}
	c.service.client = c
	c.Metar = (*MetarService)(&c.service)
	c.Taf = (*TafService)(&c.service)
	return c
}

// NewRequest creates a new API request used by the client.
func (c *Client) NewRequest(datasource, requestType string, opts Options) (*http.Request, error) {
	req, err := http.NewRequest("GET", c.BaseURL.String(), nil)
	if err != nil {
		return nil, err
	}

	params, err := optsFromParams(opts)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
	q.Add("datasource", datasource)
	q.Add("requesttype", requestType)
	q.Add("format", "xml")
	for k, v := range params {
		q.Add(k, fmt.Sprintf("%s", v))
	}
	req.URL.RawQuery = q.Encode()

	return req, nil
}

// Do executes the API call.
func (c *Client) Do(ctx context.Context, req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = xml.Unmarshal(body, &v)
	if err != nil {
		return nil, err
	}

	response := Response{
		Response: resp,
	}
	return &response, nil
}

func optsFromParams(opts Options) (params map[string]interface{}, err error) {
	data, err := json.Marshal(opts)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &params)
	return
}
