package commands

import (
	"encoding/json"
	"github.com/scotta01-org/kemp-lb-go-client/pkg/types"
	"strconv"
)

const (
	LIST_VIRTUAL_SERVICES  = "listvs"
	SHOW_VIRTUAL_SERVICE   = "showvs"
	ADD_VIRTUAL_SERVICE    = "addvs"
	DELETE_VIRTUAL_SERVICE = "delvs"
)

const (
	LAYER7 = "1"
	LAYER4 = "0"
)

type ListCommand struct {
	baseCommand
}

type VSByID struct {
	baseCommand
	VirtualService string `json:"vs"`
}

type VSByIP struct {
	VSByID
	Port     string `json:"port"`
	Protocol string `json:"prot"`
	Name     string `json:"NickName,omitempty"`
	Type     string `json:"VSType,omitempty"`
	Layer    string `json:"ForceL7,omitempty"`
}

func (c *ListCommand) JSON() ([]byte, error) {
	return json.Marshal(c)
}

func (c *ListCommand) Execute(url string, apikey string, client HTTPClient) (*types.VirtualServices, error) {
	c.APIKey = apikey
	c.Cmd = LIST_VIRTUAL_SERVICES
	jsonPayload, err := c.JSON()
	if err != nil {
		return nil, err
	}

	resp, err := c.httpRequest("POST", url, jsonPayload, client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	vs := &types.VirtualServices{}
	err = json.NewDecoder(resp.Body).Decode(vs)
	if err != nil {
		return nil, err
	}

	return vs, nil
}

func (c *VSByID) JSON() ([]byte, error) {
	return json.Marshal(c)
}

func (c *VSByID) Get(url string, apikey string, client HTTPClient, id int) (*types.VirtualService, error) {
	c.Cmd = SHOW_VIRTUAL_SERVICE
	c.APIKey = apikey
	c.VirtualService = strconv.Itoa(id)
	jsonPayload, err := c.JSON()
	if err != nil {
		return nil, err
	}

	resp, err := c.httpRequest("POST", url, jsonPayload, client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	vs := &types.VirtualService{}
	err = json.NewDecoder(resp.Body).Decode(vs)
	if err != nil {
		return nil, err
	}

	return vs, nil
}

func (c *VSByID) GetRealServers(url string, apikey string, client HTTPClient, id int) (*[]types.RealServer, error) {
	c.Cmd = SHOW_VIRTUAL_SERVICE
	c.APIKey = apikey
	c.VirtualService = strconv.Itoa(id)
	jsonPayload, err := c.JSON()
	if err != nil {
		return nil, err
	}

	resp, err := c.httpRequest("POST", url, jsonPayload, client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	vs := &types.VirtualService{}
	err = json.NewDecoder(resp.Body).Decode(vs)
	if err != nil {
		return nil, err
	}

	return &vs.Rs, nil
}

func (c *VSByID) Delete(url string, apikey string, client HTTPClient, id int) (*types.VirtualService, error) {
	c.Cmd = DELETE_VIRTUAL_SERVICE
	c.APIKey = apikey
	c.VirtualService = strconv.Itoa(id)
	jsonPayload, err := c.JSON()
	if err != nil {
		return nil, err
	}

	resp, err := c.httpRequest("POST", url, jsonPayload, client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	vs := &types.VirtualService{}
	err = json.NewDecoder(resp.Body).Decode(vs)
	if err != nil {
		return nil, err
	}

	return vs, nil
}

func (c *VSByIP) JSON() ([]byte, error) {
	return json.Marshal(c)
}

func (c *VSByIP) Get(url string, apikey string, client HTTPClient, ip string, port int, protocol string) (*types.VirtualService, error) {
	c.Cmd = SHOW_VIRTUAL_SERVICE
	c.APIKey = apikey
	c.VirtualService = ip
	c.Port = strconv.Itoa(port)
	c.Protocol = protocol
	jsonPayload, err := c.JSON()
	if err != nil {
		return nil, err
	}

	resp, err := c.httpRequest("POST", url, jsonPayload, client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	vs := &types.VirtualService{}
	err = json.NewDecoder(resp.Body).Decode(vs)
	if err != nil {
		return nil, err
	}

	return vs, nil
}

func (c *VSByIP) GetRealServers(url string, apikey string, client HTTPClient, ip string, port int, protocol string) (*[]types.RealServer, error) {
	c.Cmd = SHOW_VIRTUAL_SERVICE
	c.APIKey = apikey
	c.VirtualService = ip
	c.Port = strconv.Itoa(port)
	c.Protocol = protocol
	jsonPayload, err := c.JSON()
	if err != nil {
		return nil, err
	}

	resp, err := c.httpRequest("POST", url, jsonPayload, client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	vs := &types.VirtualService{}
	err = json.NewDecoder(resp.Body).Decode(vs)
	if err != nil {
		return nil, err
	}

	return &vs.Rs, nil
}

func (c *VSByIP) Add(url string, apikey string, client HTTPClient, ip string, port string, protocol string, name string, layer string) (*types.VirtualService, error) {
	c.Cmd = ADD_VIRTUAL_SERVICE
	c.APIKey = apikey
	c.VirtualService = ip
	c.Port = port
	c.Protocol = protocol
	c.Name = name
	c.Type = "gen"
	c.Layer = layer
	jsonPayload, err := c.JSON()
	if err != nil {
		return nil, err
	}

	resp, err := c.httpRequest("POST", url, jsonPayload, client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	vs := &types.VirtualService{}
	err = json.NewDecoder(resp.Body).Decode(vs)
	if err != nil {
		return nil, err
	}

	return vs, nil
}

func (c *VSByIP) Delete(url string, apikey string, client HTTPClient, ip string, port string, protocol string) (*types.VirtualService, error) {
	c.Cmd = DELETE_VIRTUAL_SERVICE
	c.APIKey = apikey
	c.VirtualService = ip
	c.Port = port
	c.Protocol = protocol
	jsonPayload, err := c.JSON()
	if err != nil {
		return nil, err
	}

	resp, err := c.httpRequest("POST", url, jsonPayload, client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	vs := &types.VirtualService{}
	err = json.NewDecoder(resp.Body).Decode(vs)
	if err != nil {
		return nil, err
	}

	return vs, nil
}
