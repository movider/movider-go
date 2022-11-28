// Copyright 2019 1Moby Co.,Ltd. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Movider REST API client for Go. API support for SMS, Verify, Acknowledge Verify and Cancel Verify.
package verify

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/movider/movider-go/client"
)

const (
	verifyURIPath    = "/verify"
	verifyACKURIPath = "/verify/acknowledge"
	verifyCXLURIPath = "/verify/cancel"
)

// Verify struct is a structure of result and error when a request to Movider.
type Verify struct {
	Result ResultVerify
	Error  client.ResponseError
}

// Acknowledge struct is a structure of result and error when a request to Movider.
type Acknowledge struct {
	Result ResutAcknowledge
	Error  client.ResponseError
}

// Cancel struct is a structure of result and error when a request to Movider.
type Cancel struct {
	Result ResultCancel
	Error  client.ResponseError
}

// ResultVerify struct is a structure of result send Verify when the request has success.
// This struct accepted Content-Type are JSON and XML.
type ResultVerify struct {
	XMLName   xml.Name `json:"-" xml:"xml"`
	RequestId string   `json:"request_id" xml:"request_id"`
	Number    string   `json:"number" xml:"number"`
}

// ResutAcknowledge struct is a structure of result send Acknowledge Verify when the request has success.
// This struct accepted Content-Type are JSON and XML.
type ResutAcknowledge struct {
	XMLName   xml.Name `json:"-" xml:"xml"`
	RequestId string   `json:"request_id" xml:"request_id"`
	Price     float64  `json:"price" xml:"price"`
}

// ResultCancel struct is a structure of result send Cancel Verify when the request has success.
// This struct accepted Content-Type are JSON and XML.
type ResultCancel struct {
	XMLName   xml.Name `json:"-" xml:"xml"`
	RequestId string   `json:"request_id" xml:"request_id"`
}

// Params struct is optional for send Verify.
type Params struct {
	// The length of the verification code. (Must be one of 4 or 6)
	// Default is 6.
	CodeLength int

	// Choose the language of the message template. Possible values are: en-gb, en-us, th-th.
	// Default is en-gb.
	Language string

	// Time in seconds of calling to a receiver for verification code.
	// Default is 180 seconds.
	// Note: Time periods between the first minute - before the last minute of pin_exipre.
	NextEventWait int

	// Time in seconds for verification code.
	// Default is 300 seconds
	// Note: Time periods between 120 - 600 seconds
	PinExpire int

	// Sender name
	From string

	// The tag are grouping message for report feature.
	Tag string
}

// Send a Verify from your Movider account.
// Return 2 arguments are Verify structure and error.
func Send(c *client.Client, to string, params *Params) (*Verify, error) {
	d, err := makeSendRequestData(c, to, params)
	if err != nil {
		return nil, err
	}

	url := client.Endpoint + verifyURIPath
	statusCode, bodyByte, err := c.Request(url, client.ContentTypeJSON, d)
	if err != nil {
		return nil, err
	}

	var rtn Verify
	if statusCode != http.StatusOK {
		err = json.Unmarshal(bodyByte, &rtn.Error)
	} else {
		err = json.Unmarshal(bodyByte, &rtn.Result)
	}

	if err != nil {
		return nil, err
	}

	return &rtn, nil
}

// Send an Acknowledge Verify from your Movider account.
// Return 2 arguments are Acknowledge Verify structure and error.
func SendAcknowledge(c *client.Client, requestId, code string) (*Acknowledge, error) {
	d, err := makeAcknowledgeRequestData(c, requestId, code)
	if err != nil {
		return nil, err
	}

	url := client.Endpoint + verifyACKURIPath
	statusCode, bodyByte, err := c.Request(url, client.ContentTypeJSON, d)
	if err != nil {
		return nil, err
	}

	var rtn Acknowledge
	if statusCode != http.StatusOK {
		err = json.Unmarshal(bodyByte, &rtn.Error)
	} else {
		err = json.Unmarshal(bodyByte, &rtn.Result)
	}

	if err != nil {
		return nil, err
	}

	return &rtn, nil
}

// Send a Cancel Verify from your Movider account.
// Return 2 arguments are Cancel Verify structure and error.
func SendCancel(c *client.Client, requestId string) (*Cancel, error) {
	d, err := makeCancelRequestData(c, requestId)
	if err != nil {
		return nil, err
	}

	url := client.Endpoint + verifyCXLURIPath
	statusCode, bodyByte, err := c.Request(url, client.ContentTypeJSON, d)
	if err != nil {
		return nil, err
	}

	var rtn Cancel
	if statusCode != http.StatusOK {
		err = json.Unmarshal(bodyByte, &rtn.Error)
	} else {
		err = json.Unmarshal(bodyByte, &rtn.Result)
	}

	if err != nil {
		return nil, err
	}

	return &rtn, nil
}

func makeSendRequestData(c *client.Client, to string, params *Params) (url.Values, error) {
	if to == "" {
		return nil, errors.New("to is required.")
	}

	d := url.Values{}
	d.Set("api_key", c.ApiKey)
	d.Set("api_secret", c.ApiSecret)
	d.Set("code_length", strconv.Itoa(params.CodeLength))
	d.Set("language", params.Language)
	d.Set("next_event_wait", strconv.Itoa(params.NextEventWait))
	d.Set("pin_expire", strconv.Itoa(params.PinExpire))
	d.Set("tag", params.Tag)
	d.Set("to", to)
	d.Set("from", params.From)

	return d, nil
}

func makeAcknowledgeRequestData(c *client.Client, requestId, code string) (url.Values, error) {
	if requestId == "" {
		return nil, errors.New("requestId is required.")
	}

	if code == "" {
		return nil, errors.New("code is required.")
	}

	d := url.Values{}
	d.Set("api_key", c.ApiKey)
	d.Set("api_secret", c.ApiSecret)
	d.Set("request_id", requestId)
	d.Set("code", code)

	return d, nil
}

func makeCancelRequestData(c *client.Client, requestId string) (url.Values, error) {
	if requestId == "" {
		return nil, errors.New("requestId is required.")
	}

	d := url.Values{}
	d.Set("api_key", c.ApiKey)
	d.Set("api_secret", c.ApiSecret)
	d.Set("request_id", requestId)

	return d, nil
}
