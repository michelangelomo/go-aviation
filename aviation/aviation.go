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
	baseUrl = "https://www.aviationweather.gov/adds/dataserver_current/httpparam"
)

type Client struct {
	client *http.Client

	service service

	BaseURL url.URL
	Metar   *MetarService
	Taf     *TafService
}

type service struct {
	client *Client
}

type Response struct {
	*http.Response
}

type Options struct {
	Stations                 *string    `json:"stationString,omitempty"`
	HoursBeforeNow           *float32   `json:"hoursBeforeNow,omitempty,string"`
	MostRecent               *bool      `json:"mostRecent,omitempty,string"`
	StartTime                *string    `json:"startTime,omitempty"`
	EndTime                  *string    `json:"endTime,omitempty"`
	MostRecentForEachStation *string    `json:"mostRecentForEachStation,omitempty,string"`
}

func (opts *Options) SetStations(stations string) {
	opts.Stations = &stations
}

func (opts *Options) SetHoursBeforeNow(hoursBeforeNow float32) {
	opts.HoursBeforeNow = &hoursBeforeNow
}

func (opts *Options) SetMostRecent(mostRecent bool) {
	opts.MostRecent = &mostRecent
}

func (opts *Options) SetStartTime(startTime time.Time) {
	s := startTime.UTC().Format(time.RFC3339)
	opts.StartTime = &s
}

func (opts *Options) SetEndTime(endTime time.Time) {
	e := endTime.UTC().Format(time.RFC3339)
	opts.EndTime = &e
}

func (opts *Options) SetMostRecentForEachStation(mostRecentForEachStation string) {
	opts.MostRecentForEachStation = &mostRecentForEachStation
}

func (c *Client) Client() *http.Client {
	clientCopy := *c.client
	return &clientCopy
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(baseUrl)

	c := &Client{client: httpClient, BaseURL: *baseURL}
	c.service.client = c
	c.Metar = (*MetarService)(&c.service)
	c.Taf = (*TafService)(&c.service)
	return c
}

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
