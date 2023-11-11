package client

import (
	"encoding/json"
	"github.com/scotta01-org/kemp-lb-go-client/pkg/types"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListVirtualServices(t *testing.T) {
	// Mock data that mimics a response from Kemp Load Balancer
	mockResponse := types.VirtualServices{
		Code: 200,
		VS: []types.VirtualService{
			{
				Status:       "Down",
				Index:        3,
				VSAddress:    "IP address would appear here",
				VSPort:       "53",
				Layer:        7,
				NickName:     "DNS TCP",
				Enable:       true,
				SSLReverse:   false,
				SSLReencrypt: false,
			},
			{
				Status:       "Down",
				Index:        4,
				VSAddress:    "IP address would appear here",
				VSPort:       "53",
				Layer:        7,
				NickName:     "DNS UDP",
				Enable:       true,
				SSLReverse:   false,
				SSLReencrypt: false,
			},
		},
		Status: "ok",
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check if the correct endpoint is being called
		if r.URL.Path == "/accessv2" && r.Method == "POST" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(mockResponse)
		} else {
			w.WriteHeader(http.StatusNotFound)
		}
	}))
	defer server.Close()

	// Initialize KempClient with the mock server URL
	client := NewKempClient(server.URL, "test-api-key", true)

	// Call the method and assert the results
	virtualServices, err := client.ListVirtualServices()
	assert.NoError(t, err)
	assert.Equal(t, mockResponse, *virtualServices, "The returned virtual services should match the mock data")
}
