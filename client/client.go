// Copyright 2019 1Moby Co.,Ltd. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Movider REST API client for Go. API support for SMS, Verify, Acknowledge Verify and Cancel Verify.
package client

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	// Endpoint points you to Movider REST API.
	Endpoint = "https://api.movider.co/v1"

	// ExpectTimeout is used to limit http.Client waiting time.
	ExpectTimeout = 15 * time.Second

	ContentTypeJSON           = "application/json"
	ContentTypeXML            = "application/xml"
	ContentTypeFormURLEncoded = "application/x-www-form-urlencoded"
)

// Client struct is authenication REST API.
type Client struct {
	ApiKey    string
	ApiSecret string
}

// ResponseError struct is a structure of error when the request have problems.
// This struct accepted Content-Type are JSON and XML.
type ResponseError struct {
	Error struct {
		Code        int    `json:"code" xml:"code"`
		Name        string `json:"name" xml:"name"`
		Description string `json:"description" xml:"description"`
	} `json:"error" xml:"error"`
}

// New is the first function to starting Movider-Go SDK.
// Return a new Client for use SMS or Verify packages.
func New(apiKey, apiSecret string) *Client {
	return &Client{
		ApiKey:    apiKey,
		ApiSecret: apiSecret,
	}
}

// Request function is a general function for request from Movider-SDK to Movider.
// Return 3 arguments are StatusCode, body and error.
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
