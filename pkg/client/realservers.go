package client

import (
	"github.com/scotta01-org/kemp-lb-go-client/internal/commands"
	"github.com/scotta01-org/kemp-lb-go-client/pkg/types"
)

// GetRealServersByIP returns a list of real servers for a virtual service by IP, port and protocol
func (c *KempClient) GetRealServersByIP(virtualServiceIP string, virtualServicePort int, virtualServiceProtocol string) *[]types.RealServer {
	virtualService := commands.VSByIP{}
	servers, err := virtualService.GetRealServers(c.URL, c.APIKey, c.HttpClient, virtualServiceIP, virtualServicePort, virtualServiceProtocol)
	if err != nil {
		panic(err)
	}
	return servers

}

// GetRealServersByID returns a list of real servers for a virtual service by ID
func (c *KempClient) GetRealServersByID(virtualServiceID int) *[]types.RealServer {
	virtualService := commands.VSByID{}
	servers, err := virtualService.GetRealServers(c.URL, c.APIKey, c.HttpClient, virtualServiceID)
	if err != nil {
		panic(err)
	}
	return servers
}

// AddRealServer adds a real server to a virtual service by IP, port and protocol
func (c *KempClient) AddRealServer(virtualServiceIP string, virtualServicePort int, virtualServiceProtocol string, realServerIP string, realServerPort int) *types.RealServer {
	addRealServer := commands.RSByIP{}
	addedRealServer, err := addRealServer.Add(c.URL, c.APIKey, c.HttpClient, virtualServiceIP, virtualServicePort, virtualServiceProtocol, realServerIP, realServerPort)
	if err != nil {
		panic(err)
	}
	return addedRealServer
}

// DeleteRealServer deletes a real server from a virtual service by IP, port and protocol
func (c *KempClient) DeleteRealServer(virtualServiceIP string, virtualServicePort int, virtualServiceProtocol string, realServerIP string, realServerPort int) *types.RealServer {
	deleteRealServer := commands.RSByIP{}
	deletedRealServer, err := deleteRealServer.Delete(c.URL, c.APIKey, c.HttpClient, virtualServiceIP, virtualServicePort, virtualServiceProtocol, realServerIP, realServerPort)
	if err != nil {
		panic(err)
	}
	return deletedRealServer
}
