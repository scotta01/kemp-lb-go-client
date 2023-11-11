package commands

import (
	"bytes"
	"crypto/tls"
	"net/http"
)

// baseCommand provides common fields for all commands
type baseCommand struct {
	Cmd    string `json:"cmd"`
	APIKey string `json:"apikey"`
}

// httpRequest executes the common HTTP request logic
func (c *baseCommand) httpRequest(method, url string, payload []byte, client HTTPClient) (*http.Response, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	return client.Do(req)
}

// HTTPClient interface abstracts the HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// CustomHTTPClient wraps the standard http.Client
type CustomHTTPClient struct {
	client *http.Client
}

// NewCustomHTTPClient creates a new instance of CustomHTTPClient
func NewCustomHTTPClient(insecureSkipVerify bool) *CustomHTTPClient {
	return &CustomHTTPClient{
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: insecureSkipVerify},
			},
		},
	}
}

// Do send an HTTP request and returns an HTTP response
func (c *CustomHTTPClient) Do(req *http.Request) (*http.Response, error) {
	return c.client.Do(req)
}
