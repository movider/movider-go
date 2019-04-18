// (c) Copyright by CPCORE Co.,Ltd.
//
//
//
//

package movider

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// Endpoint points you to Movider REST API.
	Endpoint = "https://movider-api-gateway.1mobyline.com/v1"

	// ExpectTimeout is used to limit http.Client waiting time.
	ExpectTimeout = 15 * time.Second

	ContentTypeJSON           = "application/json"
	ContentTypeXML            = "application/xml"
	ContentTypeFormURLEncoded = "application/x-www-form-urlencoded"
)

// Client struct is ...
type Client struct {
	ApiKey    string
	ApiSecret string
}

// New function is ...
func New(apiKey, apiSecret string) *Client {
	return &Client{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
	}
}

func (c *Client) Request(url, accept string, d url.Values) (int, []byte, error) {
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(d.Encode()))
	if err != nil {
		return 0, []byte{}, err
	}

	req.Header.Set("Content-Type", ContentTypeFormURLEncoded)
	req.Header.Set("Accept", accept)

	client := &http.Client{
		Timeout: ExpectTimeout,
	}

	resp, err := client.Do(req)
	if err != nil {
		return 0, []byte{}, err
	}
	defer resp.Body.Close()

	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return resp.StatusCode, []byte{}, err
	}

	return resp.StatusCode, bodyByte, nil
}
