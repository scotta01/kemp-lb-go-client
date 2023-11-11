package client

import (
	"github.com/scotta01-org/kemp-lb-go-client/internal/commands"
	"github.com/scotta01-org/kemp-lb-go-client/pkg/types"
	"strconv"
)

// ListVirtualServices returns a list of all virtual services on the Kemp Load Balancer
func (c *KempClient) ListVirtualServices() (*types.VirtualServices, error) {
	listCommand := commands.ListCommand{}
	virtualServices, err := listCommand.Execute(c.URL, c.APIKey, c.HttpClient)
	if err != nil {
		return nil, err
	}
	return virtualServices, nil
}

// GetVirtualServiceByID returns a virtual service by ID
func (c *KempClient) GetVirtualServiceByID(virtualServiceID int) (*types.VirtualService, error) {
	vsCommand := commands.VSByID{}
	virtualService, err := vsCommand.Get(c.URL, c.APIKey, c.HttpClient, virtualServiceID)
	if err != nil {
		return nil, err
	}
	return virtualService, nil
}

// GetVirtualServiceByIP returns a virtual service by IP address, Protocol and Port
func (c *KempClient) GetVirtualServiceByIP(virtualServiceIP string, virtualServicePort int, virtualServiceProtocol string) (*types.VirtualService, error) {
	vsCommand := commands.VSByIP{}
	virtualService, err := vsCommand.Get(c.URL, c.APIKey, c.HttpClient, virtualServiceIP, virtualServicePort, virtualServiceProtocol)
	if err != nil {
		return nil, err
	}
	return virtualService, nil
}

// AddLayer4VirtualService adds a layer 4 virtual service
func (c *KempClient) AddLayer4VirtualService(virtualServiceIP string, virtualServicePort int, virtualServiceProtocol string, virtualServiceName string) (*types.VirtualService, error) {
	vsCommand := commands.VSByIP{}
	addedVirtualService, err := vsCommand.Add(c.URL, c.APIKey, c.HttpClient, virtualServiceIP, strconv.Itoa(virtualServicePort), virtualServiceProtocol, virtualServiceName, commands.LAYER4)
	if err != nil {
		return nil, err
	}
	return addedVirtualService, nil
}

// AddLayer7VirtualService adds a layer 7 virtual service
func (c *KempClient) AddLayer7VirtualService(virtualServiceIP string, virtualServicePort int, virtualServiceProtocol string, virtualServiceName string) (*types.VirtualService, error) {
	vsCommand := commands.VSByIP{}
	addedVirtualService, err := vsCommand.Add(c.URL, c.APIKey, c.HttpClient, virtualServiceIP, strconv.Itoa(virtualServicePort), virtualServiceProtocol, virtualServiceName, commands.LAYER7)
	if err != nil {
		return nil, err
	}
	return addedVirtualService, nil
}

// DeleteVirtualService deletes a virtual service
func (c *KempClient) DeleteVirtualService(virtualServiceIP string, virtualServicePort int, virtualServiceProtocol string) (*types.VirtualService, error) {
	vsCommand := commands.VSByIP{}
	deletedVirtualService, err := vsCommand.Delete(c.URL, c.APIKey, c.HttpClient, virtualServiceIP, strconv.Itoa(virtualServicePort), virtualServiceProtocol)
	if err != nil {
		return nil, err
	}
	return deletedVirtualService, nil
}
