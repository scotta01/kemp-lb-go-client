package client

import (
	"crypto/tls"
	"net/http"
)

type KempClient struct {
	URL        string
	APIKey     string
	HttpClient *http.Client
}

// NewVirtualServiceClient creates a new client with the given settings
func NewKempClient(KempAddress string, apiKey string, skipTLSChecks bool) *KempClient {
	return &KempClient{
		URL:    "https://" + KempAddress + "/accessv2",
		APIKey: apiKey,
		HttpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: skipTLSChecks},
			},
		},
	}
}
