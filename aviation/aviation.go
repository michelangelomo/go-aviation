package aviation

import (
	"net/http"
	"net/url"
)

const (
	baseUrl = "https://www.aviationweather.gov/adds/dataserver_current/httpparam"
)

type Client struct {
	client *http.Client

	service service

	baseURL url.URL
	Metar *MetarService
}

type service struct {
	client *Client
}

func (c *Client) Client() *http.Client {
	clientCopy := *c.client
	return &clientCopy
}

func NewClient(httpClient *http.Client) *Client {
	if httpClient == nil {
		httpClient = &http.Client{}
	}
	baseURL, _ := url.Parse(baseURL)

	c := &Client{client: httpClient, BaseURL: baseURL}
	c.service.client = c
	c.Metar = (*MetarService)(&c.service)
	return c
}