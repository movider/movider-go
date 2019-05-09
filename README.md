
# Movider Client Library for Go
Movider API client for Go. API support for SMS, Verify, Acknowledge Verify and Cancel Verify.

<img align="right" width="159px" src="https://movider.co/icons/icon-144x144.png">

[![Build Status](https://api.travis-ci.org/movider/movider-go.svg)](https://travis-ci.org/movider/movider-go)
[![GoDoc](https://godoc.org/github.com/movider/movider-go?status.svg)](https://godoc.org/github.com/movider/movider-go)
[![MIT licensed](https://img.shields.io/badge/license-MIT-blue.svg)](./LICENSE)

## Installation
The first need Go installed (1.10.4 or higher).
```
$ go version

go version go1.10.4
```
You can use Movider Client Library in your Go project by `go get`:
```
$ go get github.com/movider/movider-go
```

## Examples
Assuming the `go get` installation worked. You can import the Movider's package like this:
```go
import	"github.com/movider/movider-go/client"
```
Then, create a client an instance of *Client*.
```go
c := client.New("your-api-key", "your-api-secret")
```
If you not have *api_key* and *api_secret*, [Sign up](https://dashboard.movider.co/sign-up) Movider's account to use. 

## Send SMS
Send an outbound SMS from your Movider's account. So import the Movider's SMS package like this:
```go
import	"github.com/movider/movider-go/sms"
```
Then, send an SMS.
```go
d, err := sms.Send(c, []string{
	"your-recipient-number",
}, "First an SMS from Movider.", &sms.Params{})
```
`your-recipient-number` are specified numbers in E.164 format such as 66812345678, 14155552671.
## Send Verfication Code
Use Verify request to generate and send a PIN to your user. So import the Movider's Verify package like this:
```go
import	"github.com/movider/movider-go/verify"
```
Then, send a verification code.
```go
d, err := verify.Send(c, "your-recipient-number", &verify.Params{})
```
`your-recipient-number` are specified numbers in E.164 format such as 66812345678, 14155552671.
## Acknowledge Verification Code
Use Verify Acknowledge to confirm that the PIN you received from your user matches the one sent by Movider in your verify request. (Do not forget import Verify package).
```go
import	"github.com/movider/movider-go/verify"

d, err := verify.SendAcknowledge(c, "your-request-id", "your-code")
```
`your-request-id` is returned when you sent verification code complete.
`your-code` is verification code by your user.
## Cancel Verification Code
Control the progress of your verify requests. To cancel an existing verify request. (Do not forget import Verify package).
```go
import	"github.com/movider/movider-go/verify"

d, err := verify.SendCancel(c, "your-request-id")
```
`your-request-id` is returned when you sent verification code complete.
## Documentation
Complete documentation, instructions, and examples are available at [https://movider.co](https://movider.co)

## License
Movider client library for Go is licensed under [The MIT License](./LICENSE).  Copyright (c) 2019 1Moby Co.,Ltd