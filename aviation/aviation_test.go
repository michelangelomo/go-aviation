package aviation

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

const (
	urlPath = "/httpparam"
)

func setupClient() (client *Client, mux *http.ServeMux) {
	mux = http.NewServeMux()

	apiHandler := http.NewServeMux()
	apiHandler.Handle(urlPath+"/", http.StripPrefix(urlPath, mux))

	server := httptest.NewServer(apiHandler)

	client = NewClient(nil)
	url, _ := url.Parse(server.URL + urlPath + "/")
	client.BaseURL = *url

	return client, mux
}

func testMethod(t *testing.T, r *http.Request, want string) {
	t.Helper()
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}
