package commands

import (
	"bytes"
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
