package client

import (
	"github.com/scotta01-org/kemp-lb-go-client/internal/commands"
	"github.com/scotta01-org/kemp-lb-go-client/pkg/types"
)

// GetNetworkInterfaces returns a list of network interfaces
func (c *KempClient) GetNetworkInterfaces() (*types.Stats, error) {
	network := commands.StatsCommand{}
	networkInterfaces, err := network.Get(c.URL, c.APIKey, c.HttpClient)
	if err != nil {
		return nil, err
	}
	return networkInterfaces, nil
}
