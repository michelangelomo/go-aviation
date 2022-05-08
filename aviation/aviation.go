package aviation

import (
	"context"
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	baseUrl = "https://www.aviationweather.gov/adds/dataserver_current/httpparam"
)

type Client struct {
	client *http.Client

	service service

	BaseURL url.URL
	Metar *MetarService
}

type service struct {
	client *Client
}

type Response struct {
	*http.Response
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
	return c
}

func (c *Client) NewRequest(datasource, requestType string, opts map[string]string) (*http.Request, error) {
	req, err := http.NewRequest("GET", c.BaseURL.String(), nil)
	if err != nil {
		return nil, err
	}

	q := req.URL.Query()
    q.Add("datasource", datasource)
    q.Add("requesttype", requestType)
	q.Add("format", "xml")
	for k, v := range opts {
		q.Add(k, v)
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
		log.Fatalln(err)
	}
	
	xml.Unmarshal(body, &v)

	response := Response{
		Response: resp,
	}
	return &response, nil
}