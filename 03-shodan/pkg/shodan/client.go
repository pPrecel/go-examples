package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	ApiURL           = "https://api.shodan.io"
	ApiInfoFormat    = ApiURL + "/api-info?key=%s"
	HostSearchFormat = ApiURL + "/shodan/host/search?key=%s&query=%s"
)

type Client struct {
	apiKey string
}

func New(key string) *Client {
	return &Client{
		apiKey: key,
	}
}

func (c *Client) APIInfo() (*APIInfo, error) {
	res, err := http.Get(fmt.Sprintf(ApiInfoFormat, c.apiKey))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret APIInfo
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}

func (c *Client) Hosts(query string) (*HostList, error) {
	res, err := http.Get(fmt.Sprintf(HostSearchFormat, c.apiKey, query))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret HostList
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}
	return &ret, nil
}
