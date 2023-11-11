package client

import (
	"net/http"
	"testing"
)

func TestNewKempClient(t *testing.T) {
	kempAddress := "example.com"
	apiKey := "test-api-key"
	skipTLSChecks := true

	client := NewKempClient(kempAddress, apiKey, skipTLSChecks)

	if client.URL != "https://example.com/accessv2" {
		t.Errorf("Expected URL to be 'https://example.com/accessv2', got %s", client.URL)
	}

	if client.APIKey != apiKey {
		t.Errorf("Expected APIKey to be '%s', got %s", apiKey, client.APIKey)
	}

	if client.HttpClient == nil {
		t.Fatal("HttpClient should not be nil")
	}

	transport := client.HttpClient.Transport.(*http.Transport)
	if transport.TLSClientConfig.InsecureSkipVerify != skipTLSChecks {
		t.Errorf("Expected InsecureSkipVerify to be %v, got %v", skipTLSChecks, transport.TLSClientConfig.InsecureSkipVerify)
	}
}
