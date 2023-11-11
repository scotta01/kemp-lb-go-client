package commands

import (
	"encoding/json"
	"github.com/scotta01-org/kemp-lb-go-client/pkg/types"
	"strconv"
)

const SHOW_STATS = "stats"
const SHOW_INTERFACE = "showiface"

type StatsCommand struct {
	baseCommand
}

type InterfaceCommand struct {
	baseCommand
	Interface string `json:"interface"`
}

func (c *StatsCommand) JSON() ([]byte, error) {
	return json.Marshal(c)
}

func (c *StatsCommand) Get(url string, apikey string, client HTTPClient) (*types.Stats, error) {
	c.Cmd = SHOW_STATS
	c.APIKey = apikey
	jsonPayload, err := c.JSON()
	if err != nil {
		return nil, err
	}
	resp, err := c.httpRequest("POST", url, jsonPayload, client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	stats := &types.Stats{}
	err = json.NewDecoder(resp.Body).Decode(stats)
	if err != nil {
		return nil, err
	}

	return stats, nil
}

func (c *InterfaceCommand) JSON() ([]byte, error) {
	return json.Marshal(c)
}

func (c *InterfaceCommand) Get(url string, apikey string, client HTTPClient, interfaceID int) (*types.Interfaces, error) {
	c.Cmd = SHOW_INTERFACE
	c.APIKey = apikey
	c.Interface = strconv.Itoa(interfaceID)
	jsonPayload, err := c.JSON()
	if err != nil {
		return nil, err
	}
	resp, err := c.httpRequest("POST", url, jsonPayload, client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	iface := &types.Interfaces{}
	err = json.NewDecoder(resp.Body).Decode(iface)
	if err != nil {
		return nil, err
	}

	return iface, nil
}
