package commands

import (
	"encoding/json"
	"github.com/scotta01-org/kemp-lb-go-client/pkg/types"
	"strconv"
)

const ADD_REAL_SERVER = "addrs"
const DELETE_REAL_SERVER = "delrs"

type RSByIP struct {
	VSByID
	Port           string `json:"port"`
	Protocol       string `json:"prot"`
	Name           string `json:"NickName,omitempty"`
	Type           string `json:"VSType,omitempty"`
	Layer          string `json:"ForceL7,omitempty"`
	RealServerIP   string `json:"rs"`
	RealServerPort string `json:"rsport"`
}

func (c *RSByIP) JSON() ([]byte, error) {
	return json.Marshal(c)
}

func (c *RSByIP) Add(url string, apikey string, client HTTPClient, ip string, port int, protocol string, realserverIP string, realserverPort int) (*types.RealServer, error) {
	c.Cmd = ADD_REAL_SERVER
	c.APIKey = apikey
	c.VirtualService = ip
	c.Port = strconv.Itoa(port)
	c.Protocol = protocol
	c.RealServerIP = realserverIP
	c.RealServerPort = strconv.Itoa(realserverPort)
	jsonPayload, err := c.JSON()
	if err != nil {
		return nil, err
	}

	resp, err := c.httpRequest("POST", url, jsonPayload, client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	vs := &types.RealServer{}
	err = json.NewDecoder(resp.Body).Decode(vs)
	if err != nil {
		return nil, err
	}

	return vs, nil
}

func (c *RSByIP) Delete(url string, apikey string, client HTTPClient, ip string, port int, protocol string, realserverIP string, realserverPort int) (*types.RealServer, error) {
	c.Cmd = DELETE_REAL_SERVER
	c.APIKey = apikey
	c.VirtualService = ip
	c.Port = strconv.Itoa(port)
	c.Protocol = protocol
	c.RealServerIP = realserverIP
	c.RealServerPort = strconv.Itoa(realserverPort)
	jsonPayload, err := c.JSON()
	if err != nil {
		return nil, err
	}

	resp, err := c.httpRequest("POST", url, jsonPayload, client)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	vs := &types.RealServer{}
	err = json.NewDecoder(resp.Body).Decode(vs)
	if err != nil {
		return nil, err
	}

	return vs, nil
}
