package client

import (
	"crypto/tls"
	"net/http"
	"net/url"
)

type KempClient struct {
	URL        string
	APIKey     string
	HttpClient *http.Client
}

// NewVirtualServiceClient creates a new client with the given settings
func NewKempClient(KempAddress string, apiKey string, skipTLSChecks bool) *KempClient {
	parsedURL, err := url.Parse(KempAddress)
	if err != nil {
		// Handle the error according to your application's requirements
		panic(err)
	}

	// Add https as the default scheme if none is present
	if parsedURL.Scheme == "" {
		parsedURL.Scheme = "https"
	}

	// Reconstruct the URL with the scheme and host, without any path
	finalURL := parsedURL.Scheme + "://" + parsedURL.Host + "/accessv2"

	if parsedURL.Host == "" {
		// If the original URL was missing a scheme, the Host field will be empty.
		// In this case, use the whole original input as the host.
		finalURL = parsedURL.Scheme + "://" + KempAddress + "/accessv2"
	}
	return &KempClient{
		URL:    finalURL,
		APIKey: apiKey,
		HttpClient: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: skipTLSChecks},
			},
		},
	}
}
