// Copyright 2019 1Moby Co.,Ltd. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

// Movider API client for Go. API support for SMS, Verify, Acknowledge Verify and Cancel Verify.
package sms

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"net/http"
	"net/url"
	"strings"

	"github.com/movider/movider-go/client"
)

const (
	smsURIPath = "/sms"
)

// SMS struct is a structure of result and error when a request to Movider.
type Sms struct {
	Result ResultSms
	Error  client.ResponseError
}

// ResultSms struct is a structure of result send SMS when the request has success.
// This struct accepted Content-Type are JSON and XML.
type ResultSms struct {
	XMLName            xml.Name      `json:"-" xml:"xml"`
	RemainingBalance   float64       `json:"remaining_balance" xml:"remaining_balance"`
	TotalSms           int           `json:"total_sms" xml:"total_sms"`
	PhoneNumberList    []PhoneNumber `json:"phone_number_list" xml:"phone_number_list"`
	BadPhoneNumberList []BadNumber   `json:"bad_phone_number_list" xml:"bad_phone_number_list"`
}

// PhoneNumber struct is a substructure of ResultSms.
// The list is the correct phone number.
// This struct accepted Content-Type are JSON and XML.
type PhoneNumber struct {
	Number    string  `json:"number" xml:"number"`
	MessageId string  `json:"message_id" xml:"message_id"`
	Price     float64 `json:"price" xml:"price"`
}

// BadNumber struct is a substructure of ResultSms.
// The list is the incorrect phone number.
// This struct accepted Content-Type are JSON and XML.
type BadNumber struct {
	Number string `json:"number" xml:"number"`
	Msg    string `json:"msg" xml:"msg"`
}

// Params struct is optional for send SMS.
type Params struct {
	// The webhook endpoint the delivery report for this sms is sent. This parameter overrides the webhook endpoint you set in dashboard.
	// Example: http://example.com/dr
	CallbackUrl string

	// The method of webhook endpoint the delivery report. Choose between are  GET or POST methods.
	// Default is GET.
	CallbackMethod string

	// The name or number the message should be sent from. You can choose default senders are MOVIDER, MOVIDEROTP, MVDVERIFY and MVDSMS.
	// Default is MOVIDER.
	From string

	// The tag are grouping message for report feature.
	Tag string
}

// Send an SMS from your Movider account.
// Return 2 arguments are SMS structure and error.
func Send(c *client.Client, to []string, text string, params *Params) (*Sms, error) {
	d, err := makeRequestData(c, to, text, params)
	if err != nil {
		return nil, err
	}

	url := client.Endpoint + smsURIPath
	statusCode, bodyByte, err := c.Request(url, client.ContentTypeJSON, d)
	if err != nil {
		return nil, err
	}

	var rtn Sms
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

func makeRequestData(c *client.Client, to []string, text string, params *Params) (url.Values, error) {
	if len(to) == 0 {
		return nil, errors.New("at least 1 receiver is required.")
	}

	if text == "" {
		return nil, errors.New("text is required.")
	}

	d := url.Values{}
	d.Set("api_key", c.ApiKey)
	d.Set("api_secret", c.ApiSecret)
	d.Set("callback_url", params.CallbackUrl)
	d.Set("callback_method", params.CallbackMethod)
	d.Set("from", params.From)
	d.Set("tag", params.Tag)
	d.Set("text", text)
	d.Set("to", strings.Join(to, ","))

	return d, nil
}
