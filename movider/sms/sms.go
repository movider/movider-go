package sms

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"movider"
	"net/http"
	"net/url"
	"strings"
)

const (
	smsURIPath = "/sms"
)

type Sms struct {
	XMLName            xml.Name      `json:"-" xml:"xml"`
	Cmd                string        `json:"cmd" xml:"cmd"`
	RemainingBalance   float64       `json:"remaining_balance" xml:"remaining_balance"`
	TotalSms           int           `json:"total_sms" xml:"total_sms"`
	PhoneNumberList    []PhoneNumber `json:"phone_number_list" xml:"phone_number_list"`
	BadPhoneNumberList []BadNumber   `json:"bad_phone_number_list" xml:"bad_phone_number_list"`
}

type PhoneNumber struct {
	Number    string  `json:"number" xml:"number"`
	MessageId string  `json:"message_id" xml:"message_id"`
	Price     float64 `json:"price" xml:"price"`
}

type BadNumber struct {
	Number string `json:"number" xml:"number"`
	Msg    string `json:"msg" xml:"msg"`
}

type Params struct {
	CallbackUrl    string
	CallbackMethod string
	Tag            string
}

type smsErrorResponse struct {
	Error struct {
		Code        int    `json:"code" xml:"code"`
		Name        string `json:"name" xml:"name"`
		Description string `json:"description" xml:"description"`
	} `json:"error" xml:"error"`
}

func Send(c *movider.Client, to []string, text string, params *Params) (*Sms, error) {
	d, err := makeRequestData(c, to, text, params)
	if err != nil {
		return nil, err
	}

	url := movider.Endpoint + smsURIPath
	statusCode, bodyByte, err := c.Request(url, movider.ContentTypeJSON, d)
	if err != nil {
		return nil, err
	}

	if statusCode != http.StatusOK {
		var v smsErrorResponse
		err = json.Unmarshal(bodyByte, &v)
		if err != nil {
			return nil, err
		}
		msg := fmt.Sprintf("[%d] %s: %s", v.Error.Code, v.Error.Name, v.Error.Description)
		return nil, errors.New(msg)
	}

	var v Sms
	err = json.Unmarshal(bodyByte, &v)
	if err != nil {
		return nil, err
	}

	return &v, nil
}

func makeRequestData(c *movider.Client, to []string, text string, params *Params) (url.Values, error) {
	if len(to) == 0 {
		return nil, errors.New("at least 1 receiver is required.")
	}

	if text == "" {
		return nil, errors.New("text is required.")
	}

	d := url.Values{
		"api_key":         {c.ApiKey},
		"api_secret":      {c.ApiSecret},
		"text":            {text},
		"to":              {strings.Join(to, ",")},
		"callback_url":    {params.CallbackUrl},
		"callback_method": {params.CallbackMethod},
		"tag":             {params.Tag},
	}

	return d, nil
}
